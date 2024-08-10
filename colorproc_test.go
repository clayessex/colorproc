package main

import "testing"

func TestRgbToHsl(t *testing.T) {
	data := []struct {
		input []int
		want  []int
	}{
		{[]int{255, 255, 255}, []int{0, 0, 100}},
		{[]int{128, 128, 128}, []int{0, 0, 50}},
		{[]int{255, 0, 255}, []int{300, 100, 50}},
		{[]int{247, 243, 217}, []int{52, 65, 91}},
		{[]int{127, 28, 158}, []int{286, 70, 36}},
		{[]int{11, 34, 38}, []int{189, 55, 10}},
		{[]int{33, 255, 144}, []int{150, 100, 56}},
		{[]int{206, 191, 136}, []int{47, 42, 67}},
		{[]int{128, 128, 255}, []int{240, 100, 75}},
		{[]int{65, 27, 234}, []int{251, 83, 51}},
	}

	for i, v := range data {
		h, s, l := RgbToHsl(v.input[0], v.input[1], v.input[2])
		if h != v.want[0] {
			t.Fatalf("hue calc failed on test #%v: got: %v, want %v", i, h, v.want[0])
		}
		if s != v.want[1] {
			t.Fatalf("sat calc failed on test #%v: got: %v, want %v", i, s, v.want[1])
		}
		if l != v.want[2] {
			t.Fatalf("lum calc failed on test #%v: got: %v, want %v", i, l, v.want[2])
		}
	}
}
