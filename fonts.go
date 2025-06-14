package main

var Font8 = map[rune][8]byte{
	// Английские заглавные буквы: A–Z
	'A': {0x00, 0x00, 0xFC, 0x22, 0x21, 0x22, 0xFC, 0x00},
	'B': {0x00, 0x00, 0xFF, 0x91, 0x91, 0x91, 0x6E, 0x00},
	'C': {0x00, 0x00, 0x7E, 0x81, 0x81, 0x81, 0x42, 0x00},
	'D': {0x00, 0x00, 0xFF, 0x81, 0x81, 0x42, 0x3C, 0x00},
	'E': {0x00, 0x00, 0xFF, 0x91, 0x91, 0x91, 0x81, 0x00},
	'F': {0x00, 0x00, 0xFF, 0x11, 0x11, 0x11, 0x01, 0x00},
	'G': {0x00, 0x00, 0x7E, 0x81, 0x81, 0x91, 0x72, 0x00},
	'H': {0x00, 0x00, 0xFF, 0x10, 0x10, 0x10, 0xFF, 0x00},
	'I': {0x00, 0x00, 0x81, 0x81, 0xFF, 0x81, 0x81, 0x00},
	'J': {0x00, 0x60, 0x80, 0x80, 0x80, 0x7F, 0x00, 0x00},
	'K': {0x00, 0x00, 0xFF, 0x18, 0x24, 0x42, 0x81, 0x00},
	'L': {0x00, 0x00, 0xFF, 0x80, 0x80, 0x80, 0x80, 0x00},
	'M': {0x00, 0x00, 0xFF, 0x02, 0x0C, 0x02, 0xFF, 0x00},
	'N': {0x00, 0x00, 0xFF, 0x02, 0x04, 0x08, 0xFF, 0x00},
	'O': {0x00, 0x00, 0x7E, 0x81, 0x81, 0x81, 0x7E, 0x00},
	'P': {0x00, 0x00, 0xFF, 0x11, 0x11, 0x11, 0x0E, 0x00},
	'Q': {0x00, 0x00, 0x7E, 0x81, 0x85, 0x82, 0x7D, 0x00},
	'R': {0x00, 0x00, 0xFF, 0x11, 0x19, 0x15, 0xE2, 0x00},
	'S': {0x00, 0x00, 0x46, 0x89, 0x91, 0x91, 0x62, 0x00},
	'T': {0x00, 0x00, 0x01, 0x01, 0xFF, 0x01, 0x01, 0x00},
	'U': {0x00, 0x00, 0x7F, 0x80, 0x80, 0x80, 0x7F, 0x00},
	'V': {0x00, 0x00, 0x1F, 0x60, 0x80, 0x60, 0x1F, 0x00},
	'W': {0x00, 0x00, 0x7F, 0x80, 0x70, 0x80, 0x7F, 0x00},
	'X': {0x00, 0x00, 0xC3, 0x24, 0x18, 0x24, 0xC3, 0x00},
	'Y': {0x00, 0x00, 0x07, 0x08, 0xF0, 0x08, 0x07, 0x00},
	'Z': {0x00, 0x00, 0xC1, 0xA1, 0x91, 0x89, 0x85, 0x00},
	'a': {0x00, 0x40, 0xA8, 0xA8, 0xA8, 0xF0, 0x00, 0x00},
	'b': {0x00, 0xFE, 0x90, 0x90, 0x90, 0x60, 0x00, 0x00},
	'c': {0x00, 0x70, 0x88, 0x88, 0x88, 0x00, 0x00, 0x00},
	'd': {0x00, 0x60, 0x90, 0x90, 0x90, 0xFE, 0x00, 0x00},
	'e': {0x00, 0x70, 0xA8, 0xA8, 0xA8, 0x30, 0x00, 0x00},
	'f': {0x00, 0xFC, 0x12, 0x12, 0x00, 0x00, 0x00, 0x00},
	'g': {0x00, 0x18, 0xA4, 0xA4, 0x78, 0x00, 0x00, 0x00},
	// Основные спецсимволы: ! @ # $ % & * ( ) . , ?
	'!': {0x00, 0x00, 0x00, 0x00, 0xFD, 0x00, 0x00, 0x00},
	'@': {0x00, 0x7C, 0x82, 0xB9, 0xA5, 0xB9, 0x42, 0x3C},
	'#': {0x00, 0x48, 0xFC, 0x48, 0x48, 0xFC, 0x48, 0x00},
	'$': {0x00, 0x24, 0x54, 0xD6, 0x54, 0x48, 0x00, 0x00},
	'%': {0x00, 0xC2, 0xC4, 0x08, 0x10, 0x26, 0x46, 0x00},
	'&': {0x00, 0x6C, 0x92, 0xAA, 0x44, 0x0A, 0x00, 0x00},
	'*': {0x00, 0x00, 0x2A, 0x1C, 0x1C, 0x2A, 0x00, 0x00},
	'(': {0x00, 0x00, 0x00, 0x38, 0xC6, 0x01, 0x00, 0x00},
	')': {0x00, 0x00, 0x00, 0x01, 0xC6, 0x38, 0x00, 0x00},
	'.': {0x00, 0x00, 0x00, 0x00, 0xC0, 0x00, 0x00, 0x00},
	',': {0x00, 0x00, 0x00, 0x00, 0xD0, 0x00, 0x00, 0x00},
	'?': {0x00, 0x00, 0x02, 0x01, 0xB1, 0x09, 0x06, 0x00},
	':': {0x00, 0x00, 0x00, 0x66, 0x66, 0x00, 0x00, 0x00},
	// Цифры: 0–9
	'0': {0x00, 0x00, 0x7E, 0x81, 0x81, 0x7E, 0x00, 0x00},
	'1': {0x00, 0x00, 0x00, 0x01, 0xFF, 0x00, 0x00, 0x00},
	'2': {0x00, 0x00, 0xC2, 0xA1, 0x91, 0x8E, 0x00, 0x00},
	'3': {0x00, 0x00, 0x42, 0x81, 0x89, 0x76, 0x00, 0x00},
	'4': {0x00, 0x00, 0x30, 0x28, 0x24, 0xFE, 0x20, 0x00},
	'5': {0x00, 0x00, 0x4F, 0x89, 0x89, 0x71, 0x00, 0x00},
	'6': {0x00, 0x00, 0x7E, 0x91, 0x91, 0x62, 0x00, 0x00},
	'7': {0x00, 0x00, 0x01, 0xF1, 0x09, 0x05, 0x03, 0x00},
	'8': {0x00, 0x00, 0x76, 0x89, 0x89, 0x76, 0x00, 0x00},
	'9': {0x00, 0x00, 0x4E, 0x91, 0x91, 0x7E, 0x00, 0x00},
}

// 'a': { 0x00, 0x00, 0xF0, 0xA8, 0xA8, 0xA8, 0x40, 0x00 },
// 'b': { 0xE0, 0xA8, 0xA8, 0xA8, 0xA8, 0xA8, 0x00, 0x00 },
// 'c': { 0x00, 0x00, 0x70, 0x88, 0x88, 0x88, 0x00, 0x00 },
// 'd': { 0x40, 0xA8, 0xA8, 0xA8, 0xA8, 0xA8, 0x00, 0x00 },
// 'e': { 0x00, 0x00, 0x70, 0x88, 0xF8, 0x88, 0x70, 0x00 },
// 'f': { 0xE0, 0x20, 0x20, 0x30, 0x20, 0x20, 0x00, 0x00 },
// 'g': { 0x00, 0x40, 0xB0, 0xA8, 0xA8, 0xA8, 0x70, 0x00 },
// 'h': { 0xE0, 0x08, 0x08, 0x08, 0x08, 0x08, 0x00, 0x00 },
// 'i': { 0x00, 0x00, 0x20, 0x00, 0x20, 0x20, 0xE0, 0x00 },
// 'j': { 0x00, 0x00, 0x20, 0x00, 0x20, 0x20, 0x70, 0x00 }, //
// 'l': { 0xE0, 0x20, 0x20, 0x20, 0x20, 0x20, 0x00, 0x00 },
// 'm': { 0xE0, 0x04, 0x08, 0x10, 0x08, 0x04, 0x00, 0x00 },
// 'n': { 0xE0, 0x04, 0x04, 0x04, 0x04, 0x04, 0x00, 0x00 },
// 'o': { 0x00, 0x00, 0x70, 0x88, 0x88, 0x88, 0x70, 0x00 },
// 'p': { 0xE0, 0xA8, 0xA8, 0xA8, 0xA0, 0x00, 0x00, 00x00 },
// 'q': { 0x00, 0x00, 0x70, 0x88, 0x88, 0x88, 0x40, 0x00 },
// 'r': { 0xE0, 0x08, 0x08, 0x10, 0x20, 0x00, 0x00, 0x00 },
// 's': { 0x00, 0x00, 0x40, 0xA8, 0x70, 0x10, 0x00, 0x00 },
// 't': { 0x00, 0x20, 0xF8, 0x20, 0x20, 0x20, 0x00, 0x00 },
// 'u': { 0xE0, 0x04, 0x04, 0x04, 0x04, 0xE0, 0x00, 0x00 },
// 'v': { 0xE0, 0x04, 0x04, 0x08, 0x04, 0x00, 0x00, 0x00 },
// 'w': { 0xE0, 0x04, 0x08, 0x10, 0x08, 0xE0, 0x00, 0x00 },
// 'x': { 0x00, 0x00, 0xA8, 0x50, 0x50, 0xA8, 0x00, 0x00 },
// 'y': { 0x00, 0x04, 0x08, 0x10, 0x08, 0x04, 0xE0, 0x00 },
// 'z': { 0x00, 0x00, 0x50, 0x50, 0x50, 0x50, 0x00, 0x00 }
