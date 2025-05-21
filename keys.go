package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

const (
	adcPath = "/sys/bus/iio/devices/iio:device0/in_voltage4_raw"
	vref    = 3.3    // Опорное напряжение АЦП
	maxADC  = 4095.0 // 12-битный АЦП
)

func ReadADC() (float64, error) {
	data, err := ioutil.ReadFile(adcPath)
	if err != nil {
		return 0, err
	}
	rawStr := strings.TrimSpace(string(data))
	rawVal, err := strconv.Atoi(rawStr)
	if err != nil {
		return 0, err
	}
	voltage := (float64(rawVal) / maxADC) * vref
	return voltage, nil
}

func DetectKey(voltage float64) string {
	switch {
	case voltage < 0.1:
		return "ESC"
	case voltage >= 0.65 && voltage <= 0.85:
		return "ENT"
	case voltage >= 1.3 && voltage <= 1.5:
		return "DOWN"
	case voltage >= 2.1 && voltage <= 2.3:
		return "UP"
	case voltage >= 2.5:
		return "NO_KEY"
	default:
		return "UNKNOWN"
	}
}

func mainkey() error {
	var lastKey string

	for {
		voltage, err := ReadADC()
		if err != nil {
			fmt.Println("Ошибка чтения ADC:", err)
			time.Sleep(500 * time.Millisecond)
			continue
		}

		key := DetectKey(voltage)
		if key != lastKey {
			fmt.Printf("Vadc = %.3f V, Key = %s\n", voltage, key)
			if key == "UP" {
				if err := lcdClear(); err != nil {

				}
				if err := lcdWriteFontText3("UP PRESS", 0, 0); err != nil {

				}
			}
			if key == "ESC" {
				if err := lcdClear(); err != nil {

				}
				if err := lcdWriteFontText3("ESC PRESS", 0, 0); err != nil {

				}
			}
			if key == "ENT" {
				if err := lcdClear(); err != nil {

				}
				if err := lcdWriteFontText3("ENT PRESS", 0, 0); err != nil {

				}
			}
			if key == "DOWN" {
				if err := lcdClear(); err != nil {

				}
				if err := lcdWriteFontText3("DOWN PRESS", 0, 0); err != nil {

				}
			}
			if key == "NO_KEY" {
				if err := page1(); err != nil {

				}
			}
			lastKey = key
		}

		time.Sleep(100 * time.Millisecond) // 10 Гц
	}

}
