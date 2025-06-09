package main

import (
	"fmt"
	"os"
)

func page11() error {

	// if err := lcdClear(); err != nil {
	// 	fmt.Fprintf(os.Stderr, "Failed to clear LCD: %v\n", err)
	// 	os.Exit(1)
	// }
	if err := lcdWriteFontText3("ESIL", 0, 0); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to clear LCD: %v\n", err)
		os.Exit(1)
	}
	if err := lcdWriteFontText3("POWER", 2, 2); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to clear LCD: %v\n", err)
		os.Exit(1)
	}
	if err := lcdWriteFontText3("SYSTEM", 4, 4); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to clear LCD: %v\n", err)
		os.Exit(1)
	}
	if err := lcdWriteFontText3("PRESS ENTER", 6, 0); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to write to LCD: %v\n", err)
		os.Exit(1)
	}

	// for i := 0; i < 10; i++ {
	// 	// Очистить надпись (заменить на пробелы той же длины)
	// 	if err := lcdWriteFontText3("PRESS         ", 6, 0); err != nil {
	// 		fmt.Fprintf(os.Stderr, "Failed to clear LCD: %v\n", err)
	// 		os.Exit(1)
	// 	}
	// 	time.Sleep(1 * time.Second)
	// 	// Показать надпись
	// 	if err := lcdWriteFontText3("PRESS ENTER", 6, 0); err != nil {
	// 		fmt.Fprintf(os.Stderr, "Failed to write to LCD: %v\n", err)
	// 		os.Exit(1)
	// 	}
	// 	time.Sleep(1 * time.Second)

	// }
	return nil
}
func page12() error {

	if err := lcdClear(); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to clear LCD: %v\n", err)
		os.Exit(1)
	}
	if err := lcdWriteFontText3("ENTER PRESS", 0, 0); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to clear LCD: %v\n", err)
		os.Exit(1)
	}

	return nil
}

func page1() error {
	// if err := lcdClear(); err != nil {
	// 	return err
	// }
	if err := lcdWriteFontText3("PAGE 1", 0, 3); err != nil {
		return err
	}
	if err := lcdWriteFontText3("IPP:", 2, 0); err != nil {
		return err
	}
	return lcdWriteFontText3(globalConfig.IP, 3, 0)
}

func page2() error {
	// if err := lcdClear(); err != nil {
	// 	return err
	// }
	if err := lcdWriteFontText3("PAGE 2", 0, 3); err != nil {
		return err
	}
	return lcdWriteFontText3("SET1: "+fmt.Sprintf("%v", globalConfig.Set1), 3, 0)
}

func page3() error {
	if err := lcdClear(); err != nil {
		return err
	}
	if err := lcdWriteFontText3("PAGE 3", 0, 3); err != nil {
		return err
	}
	return lcdWriteFontText3("SET2: "+fmt.Sprintf("%v", globalConfig.Set2), 3, 0)
}

func page4() error {
	if err := lcdClear(); err != nil {
		return err
	}
	if err := lcdWriteFontText3("PAGE 4", 0, 3); err != nil {
		return err
	}
	return lcdWriteFontText3("SET3: "+fmt.Sprintf("%v", globalConfig.Set3), 3, 0)
}
