package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

const (
	LCD_RD  = 12  // PA12
	LCD_WR  = 13  // PA13
	LCD_A0  = 14  // PA14
	LCD_RES = 15  // PA15
	LCD_CS  = 202 // PG10
	LCD_BL  = 38  // PB6
	LCD_D0  = 11  // PA11
	LCD_D1  = 8   // PA8
	LCD_D2  = 7   // PA7
	LCD_D3  = 6   // PA6
	LCD_D4  = 5   // PA5
	LCD_D5  = 4   // PA4
	LCD_D6  = 3   // PA3
	LCD_D7  = 2   // PA2
)

type GPIO struct {
	pin int
	dir string
	fd  *os.File
}

// NewGPIO создаёт новый экземпляр GPIO
func NewGPIO(pin int) *GPIO {
	return &GPIO{
		pin: pin,
		dir: fmt.Sprintf("/sys/class/gpio/gpio%d/", pin),
	}
}

// Init инициализирует пин GPIO
func (g *GPIO) Init() error {
	// Проверка, экспортирован ли пин
	if _, err := os.Stat(g.dir); err == nil {
		// Разэкспортировать пин
		ufd, err := os.OpenFile("/sys/class/gpio/unexport", os.O_WRONLY, 0644)
		if err != nil {
			return fmt.Errorf("failed to open /sys/class/gpio/unexport: %v", err)
		}
		defer ufd.Close()
		_, err = ufd.WriteString(fmt.Sprintf("%d", g.pin))
		if err != nil {
			return fmt.Errorf("failed to unexport GPIO pin %d: %v", g.pin, err)
		}
		time.Sleep(5 * time.Millisecond)
	}

	// Экспорт пина
	efd, err := os.OpenFile("/sys/class/gpio/export", os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed to open /sys/class/gpio/export: %v", err)
	}
	defer efd.Close()

	_, err = efd.WriteString(fmt.Sprintf("%d", g.pin))
	if err != nil {
		return fmt.Errorf("failed to export GPIO pin %d: %v", g.pin, err)
	}

	time.Sleep(5 * time.Millisecond)

	// Установка направления на выход
	dfd, err := os.OpenFile(filepath.Join(g.dir, "direction"), os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed to open %sdirection: %v", g.dir, err)
	}
	defer dfd.Close()

	_, err = dfd.WriteString("out")
	if err != nil {
		return fmt.Errorf("failed to set direction for pin %d: %v", g.pin, err)
	}

	time.Sleep(5 * time.Millisecond)

	// Открытие файла value
	g.fd, err = os.OpenFile(filepath.Join(g.dir, "value"), os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed to open %svalue: %v", g.dir, err)
	}

	return nil
}

// Write устанавливает значение пина GPIO (0 или 1)
func (g *GPIO) Write(val int) error {
	if g.fd == nil {
		return fmt.Errorf("file descriptor invalid for pin %d", g.pin)
	}

	var value string
	if val != 0 {
		value = "1"
	} else {
		value = "0"
	}

	_, err := g.fd.WriteString(value)
	if err != nil {
		return fmt.Errorf("failed to write %s to %svalue: %v", value, g.dir, err)
	}

	return nil
}

// GPIO instances
var (
	gpioRD  = NewGPIO(LCD_RD)
	gpioWR  = NewGPIO(LCD_WR)
	gpioA0  = NewGPIO(LCD_A0)
	gpioRES = NewGPIO(LCD_RES)
	gpioCS  = NewGPIO(LCD_CS)
	gpioBL  = NewGPIO(LCD_BL)

	gpioD0 = NewGPIO(LCD_D0)
	gpioD1 = NewGPIO(LCD_D1)
	gpioD2 = NewGPIO(LCD_D2)
	gpioD3 = NewGPIO(LCD_D3)
	gpioD4 = NewGPIO(LCD_D4)
	gpioD5 = NewGPIO(LCD_D5)
	gpioD6 = NewGPIO(LCD_D6)
	gpioD7 = NewGPIO(LCD_D7)
)

// lcdInitPins инициализирует все пины GPIO
func lcdInitPins() error {
	gpios := []*GPIO{gpioRD, gpioWR, gpioA0, gpioRES, gpioCS, gpioBL, gpioD0, gpioD1, gpioD2, gpioD3, gpioD4, gpioD5, gpioD6, gpioD7}
	for _, gpio := range gpios {
		if err := gpio.Init(); err != nil {
			return err
		}
	}
	return nil
}

// lcdWriteByte записывает байт на пины данных
func lcdWriteByte(byte uint8) error {
	if err := gpioD0.Write(int(byte & 1)); err != nil {
		return err
	}
	if err := gpioD1.Write(int(byte & 2)); err != nil {
		return err
	}
	if err := gpioD2.Write(int(byte & 4)); err != nil {
		return err
	}
	if err := gpioD3.Write(int(byte & 8)); err != nil {
		return err
	}
	if err := gpioD4.Write(int(byte & 16)); err != nil {
		return err
	}
	if err := gpioD5.Write(int(byte & 32)); err != nil {
		return err
	}
	if err := gpioD6.Write(int(byte & 64)); err != nil {
		return err
	}
	if err := gpioD7.Write(int(byte & 128)); err != nil {
		return err
	}
	return nil
}

// lcdWriteCmd записывает команду на LCD
func lcdWriteCmd(cmd uint8) error {
	if err := lcdWriteByte(0); err != nil {
		return err
	}
	time.Sleep(20 * time.Microsecond)

	if err := gpioA0.Write(0); err != nil {
		return err
	}
	if err := gpioRD.Write(1); err != nil {
		return err
	}

	if err := gpioWR.Write(0); err != nil {
		return err
	}
	if err := lcdWriteByte(cmd); err != nil {
		return err
	}
	time.Sleep(20 * time.Microsecond)
	if err := gpioWR.Write(1); err != nil {
		return err
	}

	return nil
}

// lcdInit инициализирует LCD
func lcdInit() error {
	if err := gpioCS.Write(0); err != nil {
		return err
	}
	if err := lcdWriteCmd(0xE2); err != nil { // reset
		return err
	}
	if err := gpioRES.Write(1); err != nil {
		return err
	}

	time.Sleep(50 * time.Millisecond)
	if err := gpioRES.Write(0); err != nil {
		return err
	}
	time.Sleep(50 * time.Millisecond)
	if err := gpioRES.Write(1); err != nil {
		return err
	}
	time.Sleep(50 * time.Millisecond)

	// Команды с исправленной ориентацией (0xA1, 0xC 00xA0, 0xC8){0xA0, 0xC8, 0xA2, 0x2F, 0x26, 0x81, 0x15, 0x40, 0xAF, 0xA6, 0xA4}
	// commands := []uint8{0xA0, 0xC8, 0xA2, 0x2F, 0x26, 0x81, 0x15, 0x40, 0xAF, 0xA6, 0xA4}
	// д 0xA1 и 0xA0 (зеркальный SEG) или 0xC8 и 0xC0 (COM направление), или просто заменить 0xA1 на 0xA0 (отражение по горизонтали) и 0xC8 на 0xC0 (отражение по вертикали) — зависит от желаемой ориентаци
	commands := []uint8{0xA0, 0xC8, 0xA2, 0x2F, 0x26, 0x81, 0x15, 0x40, 0xAF, 0xA6, 0xA4}
	for _, cmd := range commands {
		if err := lcdWriteCmd(cmd); err != nil {
			return err
		}
		if cmd == 0xAF {
			time.Sleep(50 * time.Millisecond)
		}
	}

	if err := gpioBL.Write(1); err != nil {
		return err
	}

	return nil
}

// lcdWriteData записывает данные на LCD
func lcdWriteData(dat uint8) error {
	if err := lcdWriteByte(0); err != nil {
		return err
	}
	time.Sleep(20 * time.Microsecond)

	if err := gpioA0.Write(1); err != nil {
		return err
	}
	if err := gpioRD.Write(1); err != nil {
		return err
	}

	if err := gpioWR.Write(0); err != nil {
		return err
	}
	if err := lcdWriteByte(dat); err != nil {
		return err
	}
	time.Sleep(20 * time.Microsecond)

	if err := gpioWR.Write(1); err != nil {
		return err
	}

	return nil
}

// lcdClear очищает дисплей
func lcdClear() error {
	for i := uint8(0); i < 9; i++ {
		if err := lcdWriteCmd(0xB0 + i); err != nil {
			return err
		}
		if err := lcdWriteCmd(0x00); err != nil {
			return err
		}
		if err := lcdWriteCmd(0x10); err != nil {
			return err
		}
		for j := 0; j < 128; j++ {
			if err := lcdWriteData(0); err != nil {
				return err
			}
		}
	}
	return nil
}

func lcdSetPosition(page int, column int) error {
	if page < 0 || page > 7 {
		return fmt.Errorf("invalid page: %d", page)
	}
	column = column * 8
	if column < 0 || column > 127 {
		return fmt.Errorf("invalid column: %d", column)
	}

	if err := lcdWriteCmd(0xB0 + uint8(page)); err != nil {
		return fmt.Errorf("failed to set page: %v", err)
	}

	lowNibble := column & 0x0F
	highNibble := (column >> 4) & 0x0F

	if err := lcdWriteCmd(0x00 | uint8(lowNibble)); err != nil {
		return fmt.Errorf("failed to set column low nibble: %v", err)
	}
	if err := lcdWriteCmd(0x10 | uint8(highNibble)); err != nil {
		return fmt.Errorf("failed to set column high nibble: %v", err)
	}

	return nil
}

func cleanupGPIO() error {
	gpios := []*GPIO{gpioRD, gpioWR, gpioA0, gpioRES, gpioCS, gpioBL, gpioD0, gpioD1, gpioD2, gpioD3, gpioD4, gpioD5, gpioD6, gpioD7}
	for _, gpio := range gpios {
		if gpio.fd != nil {
			gpio.fd.Close()
		}
		ufd, err := os.OpenFile("/sys/class/gpio/unexport", os.O_WRONLY, 0644)
		if err != nil {
			return fmt.Errorf("failed to open /sys/class/gpio/unexport: %v", err)
		}
		defer ufd.Close()
		_, err = ufd.WriteString(fmt.Sprintf("%d", gpio.pin))
		if err != nil && !os.IsNotExist(err) {
			return fmt.Errorf("failed to unexport GPIO pin %d: %v", gpio.pin, err)
		}
	}
	return nil
}
