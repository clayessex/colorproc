package colors

import (
	"math"
	"testing"
)

// 32bit Abs - modified from abs.go in the stdlib
func Abs32(x float32) float32 {
	return math.Float32frombits(math.Float32bits(x) &^ (1 << 31))
}

func eq(a, b, eps float64) bool {
	return math.Abs(a-b) < eps
}

func eq32(a, b, eps float32) bool {
	return Abs32(a-b) < eps // not correct
}

var Data = []struct {
	hex Hex
	rgb Rgb
	hsl Hsl
}{
	{Hex("#ffffff"), Rgb{R: 255, G: 255, B: 255}, Hsl{H: 0.000000000000000, S: 0.000000000000000, L: 1.000000000000000}},
	{Hex("#000000"), Rgb{R: 0, G: 0, B: 0}, Hsl{H: 0.000000000000000, S: 0.000000000000000, L: 0.000000000000000}},
	{Hex("#808080"), Rgb{R: 128, G: 128, B: 128}, Hsl{H: 0.000000000000000, S: 0.000000000000000, L: 0.501960813999176}},
	{Hex("#ff00ff"), Rgb{R: 255, G: 0, B: 255}, Hsl{H: 300.000000000000000, S: 1.000000000000000, L: 0.500000000000000}},
	{Hex("#f7f3d9"), Rgb{R: 247, G: 243, B: 217}, Hsl{H: 52.000000000000000, S: 0.652173936367035, L: 0.909803926944733}},
	{Hex("#7f1c9e"), Rgb{R: 127, G: 28, B: 158}, Hsl{H: 285.692321777343750, S: 0.698924720287323, L: 0.364705890417099}},
	{Hex("#0b2226"), Rgb{R: 11, G: 34, B: 38}, Hsl{H: 188.888885498046875, S: 0.551020383834839, L: 0.096078433096409}},
	{Hex("#21ff90"), Rgb{R: 33, G: 255, B: 144}, Hsl{H: 150.000000000000000, S: 1.000000000000000, L: 0.564705908298492}},
	{Hex("#cebf88"), Rgb{R: 206, G: 191, B: 136}, Hsl{H: 47.142856597900391, S: 0.416666656732559, L: 0.670588254928589}},
	{Hex("#8080ff"), Rgb{R: 128, G: 128, B: 255}, Hsl{H: 240.000000000000000, S: 1.000000000000000, L: 0.750980377197266}},
	{Hex("#411bea"), Rgb{R: 65, G: 27, B: 234}, Hsl{H: 251.014495849609375, S: 0.831325292587280, L: 0.511764705181122}},
	{Hex("#223344"), Rgb{R: 34, G: 51, B: 68}, Hsl{H: 210.000000000000000, S: 0.333333343267441, L: 0.200000002980232}},
}

func TestConvertHex(t *testing.T) {
	for _, v := range Data {
		rgb, err := v.hex.ToRgb()
		if err != nil || rgb != v.rgb {
			t.Fatalf("rgb conversion failed from: %v,\n  want: %v,\n   got: %v", v.hex, v.rgb, rgb)
		}
		hsl, err := v.hex.ToHsl()
		if err != nil || hsl != v.hsl {
			t.Fatalf("hsl conversion failed from: %v,\n  want: %v,\n   got: %v", v.hex, v.hsl.FString(), hsl.FString())
		}
	}
}

func TestConvertRgb(t *testing.T) {
	for _, v := range Data {
		if vhex := v.rgb.ToHex(); vhex != v.hex {
			t.Fatalf("hex conversion failed from: %v,\n  want: %v,\n   got: %v", v.rgb, vhex, v.hex)
		}
		if vhsl := v.rgb.ToHsl(); vhsl != v.hsl {
			t.Fatalf("hsl conversion failed from: %v,\n  want: %v,\n   got: %v", v.rgb, vhsl, v.hsl)
		}
	}
}

func TestConvertHsl(t *testing.T) {
	for _, v := range Data {
		if vhex := v.hsl.ToHex(); vhex != v.hex {
			t.Fatalf("hex conversion failed from: %v,\n  want: %v,\n   got: %v", v.rgb, vhex, v.hex)
		}
		if vrgb := v.hsl.ToRgb(); vrgb != v.rgb {
			t.Fatalf("rgb conversion failed from: %v,\n  want: %v,\n   got: %v", v.hsl, vrgb, v.hsl)
		}
	}
}
