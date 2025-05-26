package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// Config структура для настроек
type Config struct {
	IP   string `json:"ip"`
	Set1 string `json:"set1"`
	Set2 string `json:"set2"`
	Set3 string `json:"set3"`
}

// LoadOrCreateConfig проверяет наличие config.json и создает шаблонный, если он отсутствует
func LoadOrCreateConfig(configFile string) (Config, error) {
	// Проверяем существование файла
	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		// Файл не существует, создаем шаблонный
		defaultConfig := Config{
			IP:   "192.168.120.9",
			Set1: "YES",
			Set2: "NO",
			Set3: "NO",
		}

		// Преобразуем структуру в JSON
		configData, err := json.MarshalIndent(defaultConfig, "", "  ")
		if err != nil {
			return Config{}, fmt.Errorf("ошибка при создании JSON: %v", err)
		}

		// Записываем JSON в файл
		err = os.WriteFile(configFile, configData, 0644)
		if err != nil {
			return Config{}, fmt.Errorf("ошибка при записи файла: %v", err)
		}
		fmt.Println("Шаблонный файл config.json успешно создан")
		return defaultConfig, nil
	} else if err != nil {
		// Ошибка при проверке файла
		return Config{}, fmt.Errorf("ошибка при проверке файла: %v", err)
	}

	// Файл существует, читаем его
	data, err := os.ReadFile(configFile)
	if err != nil {
		return Config{}, fmt.Errorf("ошибка при чтении файла: %v", err)
	}

	var config Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		return Config{}, fmt.Errorf("ошибка при разборе JSON: %v", err)
	}
	fmt.Printf("Настройки загружены: %+v\n", config)
	return config, nil
}

// SaveConfig записывает конфигурацию в файл
func SaveConfig(config Config, configFile string) error {
	configData, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return fmt.Errorf("ошибка при создании JSON: %v", err)
	}

	err = os.WriteFile(configFile, configData, 0644)
	if err != nil {
		return fmt.Errorf("ошибка при записи файла: %v", err)
	}
	fmt.Println("Конфигурация успешно сохранена в", configFile)
	return nil
}
