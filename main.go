package main

import (
	"fmt"
	"os"
	"time"
)

var globalConfig Config

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

	pages := []func() error{page1, page2, page3, page4}
	currentPage := 0
	selected := false // false = режим просмотра, true = режим выбора
	editing := false  // true = режим редактирования IP
	ipCursorPos := 0  // позиция курсора при редактировании IP

	// Отображаем дефолтную страницу page11()
	if err := page11(); err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка при запуске page11: %v\n", err)
		os.Exit(1)
	}

	for {
		key := GetCurrentKey()

		if editing {
			// Режим редактирования IP
			switch key {
			case "UP":
				// Увеличиваем цифру на текущей позиции
				globalConfig.IP = incrementIPDigit(globalConfig.IP, ipCursorPos)
				_ = lcdClear()
				_ = displayIPEditScreen(globalConfig.IP, ipCursorPos)

			case "DOWN":
				// Уменьшаем цифру на текущей позиции
				globalConfig.IP = decrementIPDigit(globalConfig.IP, ipCursorPos)
				_ = lcdClear()
				_ = displayIPEditScreen(globalConfig.IP, ipCursorPos)

			case "RIGHT":
				// Перемещаем курсор вправо, пропуская точки
				ipCursorPos++
				for ipCursorPos < len(globalConfig.IP) && globalConfig.IP[ipCursorPos] == '.' {
					ipCursorPos++
				}
				if ipCursorPos >= len(globalConfig.IP) {
					ipCursorPos = 0
				}
				_ = lcdClear()
				_ = displayIPEditScreen(globalConfig.IP, ipCursorPos)

			case "LEFT":
				// Перемещаем курсор влево, пропуская точки
				ipCursorPos--
				for ipCursorPos >= 0 && globalConfig.IP[ipCursorPos] == '.' {
					ipCursorPos--
				}
				if ipCursorPos < 0 {
					ipCursorPos = len(globalConfig.IP) - 1
					// Убедимся, что не остановились на точке
					for ipCursorPos >= 0 && globalConfig.IP[ipCursorPos] == '.' {
						ipCursorPos--
					}
				}
				_ = lcdClear()
				_ = displayIPEditScreen(globalConfig.IP, ipCursorPos)

			case "ENT":
				// Сохраняем изменения и выходим из режима редактирования
				err = SaveConfig(globalConfig, configFile)
				if err != nil {
					fmt.Printf("Ошибка при сохранении конфигурации: %v\n", err)
				}
				editing = false
				_ = lcdClear()
				_ = page1() // Возвращаемся на страницу с IP

			case "ESC":
				// Отмена редактирования без сохранения
				editing = false
				_ = lcdClear()
				_ = page1()
			}
		} else if selected {

			// Режим выбора страницы
			switch key {
			case "UP":
				currentPage++
				if currentPage >= len(pages) {
					currentPage = 0
				}
				_ = lcdClear()
				_ = pages[currentPage]()

			case "DOWN":
				currentPage--
				if currentPage < 0 {
					currentPage = len(pages) - 1
				}
				_ = lcdClear()
				_ = pages[currentPage]()

			case "ENT":
				if currentPage == 0 {
					// Если выбрана страница с IP, входим в режим редактирования
					editing = true
					ipCursorPos = 0
					_ = lcdClear()
					_ = displayIPEditScreen(globalConfig.IP, ipCursorPos)
				} else {
					// Для других страниц просто выходим из режима выбора
					selected = false
				}

			case "ESC":
				// отмена выбора — возврат на дефолтную page11
				selected = false
				_ = lcdClear()
				_ = page11()
			}
		} else {
			// Обычный режим просмотра
			if key == "ENT" {
				// вход в режим выбора и сразу показываем текущую страницу
				selected = true
				currentPage = 0 // начинаем с первой страницы
				_ = lcdClear()
				_ = pages[currentPage]()
			}
		}

		time.Sleep(100 * time.Millisecond)
	}
}

// Функция для отображения экрана редактирования IP
func displayIPEditScreen(ip string, cursorPos int) error {
	// Отображаем IP с подчеркиванием под текущей позицией курсора
	// Первая строка - заголовок
	if err := lcdWriteFontText3("EDIT IP:", 0, 0); err != nil {
		return err
	}

	// Вторая строка - IP адрес
	if err := lcdWriteFontText3(ip, 1, 0); err != nil {
		return err
	}

	// Третья строка - курсор (подчеркивание под текущей цифрой)
	cursorLine := ""
	for i := 0; i < len(ip); i++ {
		if i == cursorPos && ip[i] != '.' {
			cursorLine += "^"
		} else {
			cursorLine += " "
		}
	}
	if err := lcdWriteFontText3(cursorLine, 2, 0); err != nil {
		return err
	}

	// Четвертая строка - подсказка
	if err := lcdWriteFontText3("ENT:SAVE ESC:CANCEL", 3, 0); err != nil {
		return err
	}

	return nil
}

// Функция для увеличения цифры в IP на указанной позиции
func incrementIPDigit(ip string, pos int) string {
	bytes := []byte(ip)

	// Пропускаем точки
	for pos < len(bytes) && bytes[pos] == '.' {
		pos++
		if pos >= len(bytes) {
			pos = 0
		}
	}

	if pos >= len(bytes) {
		return ip
	}

	if bytes[pos] == '9' {
		bytes[pos] = '0'
	} else if bytes[pos] >= '0' && bytes[pos] < '9' {
		bytes[pos]++
	}

	return string(bytes)
}

// Функция для уменьшения цифры в IP на указанной позиции
func decrementIPDigit(ip string, pos int) string {
	bytes := []byte(ip)

	// Пропускаем точки
	for pos < len(bytes) && bytes[pos] == '.' {
		pos++
		if pos >= len(bytes) {
			pos = 0
		}
	}

	if pos >= len(bytes) {
		return ip
	}

	if bytes[pos] == '0' {
		bytes[pos] = '9'
	} else if bytes[pos] > '0' && bytes[pos] <= '9' {
		bytes[pos]--
	}

	return string(bytes)
}
