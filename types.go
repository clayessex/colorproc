package main

import "fmt"

type Rgb struct {
	R uint8
	G uint8
	B uint8
}

type Hsl struct {
	H uint16
	S uint8
	L uint8
}

type Hex string

func (rgb Rgb) String() string {
	return fmt.Sprintf("RGB(%v, %v, %v)", rgb.R, rgb.G, rgb.B)
}

func (hsl Hsl) String() string {
	return fmt.Sprintf("HSL(%v°, %v%%, %v%%)", hsl.H, hsl.S, hsl.L)
}

func (hex Hex) String() string {
	return string(hex)
}
