package main

import (
	"fmt"
	"os"
	"time"
)

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

func main() {
	defer func() {
		if err := cleanupGPIO(); err != nil {
			fmt.Fprintf(os.Stderr, "Failed to cleanup GPIO: %v\n", err)
		}
	}()
	configFile := "config.json"
	config, err := LoadOrCreateConfig(configFile)
	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
		return
	}
	// Используйте config.IP, config.Set1, config.Set2, config.Set3 в вашем коде
	fmt.Printf("IP: %s, Set1: %v, Set2: %v, Set3: %v\n", config.IP, config.Set1, config.Set2, config.Set3)
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
	if err := page1(); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to page1: %v\n", err)
		os.Exit(1)
	}
	config.IP = "192.168.120.104" // Новое значение IP
	err = SaveConfig(config, configFile)
	if err != nil {
		fmt.Printf("Ошибка при сохранении конфигурации: %v\n", err)
		return
	}
	if err := mainkey(); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to page1: %v\n", err)
		os.Exit(1)
	}

}
