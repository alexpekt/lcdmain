package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
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
	editingIP := false
	ipParts := strings.Split(globalConfig.IP, ".")
	currentIPPart := 0

	for {
		key := GetCurrentKey()

		if editingIP {
			// Режим редактирования IP
			switch key {
			case "UP":
				// Увеличиваем текущую часть IP
				part, _ := strconv.Atoi(ipParts[currentIPPart])
				part++
				if part > 255 {
					part = 0
				}
				ipParts[currentIPPart] = strconv.Itoa(part)
				displayIPEditing(ipParts, currentIPPart)

			case "DOWN":
				// Уменьшаем текущую часть IP
				part, _ := strconv.Atoi(ipParts[currentIPPart])
				part--
				if part < 0 {
					part = 255
				}
				ipParts[currentIPPart] = strconv.Itoa(part)
				displayIPEditing(ipParts, currentIPPart)

			case "ENT":
				// Переходим к следующей части IP или сохраняем
				currentIPPart++
				if currentIPPart >= len(ipParts) {
					// Сохраняем новый IP
					globalConfig.IP = strings.Join(ipParts, ".")
					err = SaveConfig(globalConfig, configFile)
					if err != nil {
						fmt.Printf("Ошибка при сохранении конфигурации: %v\n", err)
					}
					editingIP = false
					_ = lcdClear()
					_ = page11()
					continue
				}
				displayIPEditing(ipParts, currentIPPart)

			case "ESC":
				// Отмена редактирования
				editingIP = false
				_ = lcdClear()
				_ = page11()
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
				// Если на странице page1 (предполагаем, что это страница с IP), начинаем редактирование
				if currentPage == 0 {
					editingIP = true
					ipParts = strings.Split(globalConfig.IP, ".")
					currentIPPart = 0
					displayIPEditing(ipParts, currentIPPart)
				} else {
					selected = false
				}

			case "ESC":
				selected = false
				_ = lcdClear()
				_ = page11()
			}
		} else {
			// Основной режим
			if key == "ENT" {
				selected = true
				_ = lcdClear()
				_ = pages[currentPage]()
			}
		}

		time.Sleep(100 * time.Millisecond)
	}
}

func displayIPEditing(ipParts []string, currentPart int) error {
	if err := lcdClear(); err != nil {
		return err
	}

	// Отображаем IP с выделением редактируемой части
	line1 := "Edit IP:"
	line2 := ""
	for i, part := range ipParts {
		if i == currentPart {
			line2 += "|" + part + "|"
		} else {
			line2 += part
		}
		if i < len(ipParts)-1 {
			line2 += "."
		}
	}

	if err := lcdWriteFontText3(line1, 0, 0); err != nil {
		return err
	}
	return lcdWriteFontText3(line2, 3, 0)
}
