package main

import (
	"fmt"
	"os"
	"time"
)

var globalConfig Config

func lcdWriteFontText3(text string, page, col int) error {
	err := lcdSetPosition(page, col)
	if err != nil {
		return err
	}

	for _, r := range text {
		charData, ok := Font8[r]
		if !ok {

			charData = [8]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
		}

		for _, b := range charData {
			if err := lcdWriteData(b); err != nil {
				return err
			}
		}

		// Добавить 1 байт пустоты между символами (опционально)
		// if err := lcdWriteData(0x00); err != nil {
		// 	return err
		// }
	}

	return nil
}
func GetCurrentKey() string {
	keyMutex.Lock()
	defer keyMutex.Unlock()
	return currentKey
}

func main() {
	go mainkey()
	defer func() {
		if err := cleanupGPIO(); err != nil {
			fmt.Fprintf(os.Stderr, "Failed to cleanup GPIO: %v\n", err)
		}
	}()
	configFile := "config.json"
	config, err := LoadOrCreateConfig(configFile)
	globalConfig = config
	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
		return
	}

	// Используйте config.IP, config.Set1, config.Set2, config.Set3 в вашем коде
	// fmt.Printf("IP: %s, Set1: %v, Set2: %v, Set3: %v\n", config.IP, config.Set1, config.Set2, config.Set3)
	if err := lcdInitPins(); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to initialize pins: %v\n", err)
		os.Exit(1)
	}
	time.Sleep(30 * time.Millisecond)

	if err := lcdInit(); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to initialize LCD: %v\n", err)
		os.Exit(1)
	}
	time.Sleep(30 * time.Millisecond)

	if err := lcdClear(); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to clear LCD: %v\n", err)
		os.Exit(1)
	}
	if err := page11(); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to page1: %v\n", err)
		os.Exit(1)
	}
	// config.IP = "192.168.120.105" // Новое значение IP
	err = SaveConfig(config, configFile)
	if err != nil {
		fmt.Printf("Ошибка при сохранении конфигурации: %v\n", err)
		return
	}

	pages := []func() error{page1, page2, page3, page4}

	currentPage := 0
	selected := false // false = режим просмотра, true = режим выбора

	// Отображаем дефолтную страницу page11()
	if err := page11(); err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка при запуске page11: %v\n", err)
		os.Exit(1)
	}

	for {
		key := GetCurrentKey()

		if selected {
			switch key {
			case "UP":
				// Перемещаемся вперёд по списку страниц
				currentPage++
				if currentPage >= len(pages) {
					currentPage = 0
				}
				_ = lcdClear()
				_ = lcdWriteFontText3(fmt.Sprintf("SELECT: PAGE %d", currentPage+1), 0, 0)

			case "DOWN":
				// Перемещаемся назад по списку страниц
				currentPage--
				if currentPage < 0 {
					currentPage = len(pages) - 1
				}
				_ = lcdClear()
				_ = lcdWriteFontText3(fmt.Sprintf("SELECT: PAGE %d", currentPage+1), 0, 0)

			case "ENT":
				// Подтверждаем выбор — запускаем соответствующую страницу
				selected = false
				_ = pages[currentPage]()

			case "ESC":
				// Отмена выбора — возвращаемся на главную
				selected = false
				_ = page11()
			}
		} else {
			// Вход в режим выбора по нажатию ENT
			if key == "ENT" {
				selected = true
				_ = lcdClear()
				_ = lcdWriteFontText3(fmt.Sprintf("SELECT: PAGE %d", currentPage+1), 0, 0)
			}
		}

		time.Sleep(100 * time.Millisecond)
	}

}

// if err := mainkey(); err != nil {
// 	fmt.Fprintf(os.Stderr, "Failed to page1: %v\n", err)
// 	os.Exit(1)
// }
// for {
// 	key := GetCurrentKey()
// 	if key != "" {
// 		// fmt.Println("Текущее состояние клавиши:", key)
// 	}

// 	// Можно выполнять действия в зависимости от нажатия:
// 	if key == "ENT" {
// 		_ = page12()

// 	}
// 	time.Sleep(500 * time.Millisecond)
// }

// key := GetCurrentKey()
// switch key {
// case "UP":
// 	_ = lcdClear()
// 	_ = lcdWriteFontText3("UP PRESS", 0, 0)
// case "ESC":
// 	_ = lcdClear()
// 	_ = lcdWriteFontText3("ESC PRESS", 0, 0)
// case "ENT":
// 	_ = lcdClear()
// 	_ = lcdWriteFontText3("ENT PRESS", 0, 0)
// case "DOWN":
// 	_ = lcdClear()
// 	_ = lcdWriteFontText3("DOWN PRESS", 0, 0)
// case "NO_KEY":
// 	_ = page1()
// }
