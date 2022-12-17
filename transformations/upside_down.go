package transformations

import "image/color"

func UpsideDown(pixels [][]color.Color) {
	for i := 0; i < len(pixels); i++ {
		tr := pixels[i]
		for j := 0; j < len(tr)/2; j++ {
			k := len(tr) - j - 1
			tr[j], tr[k] = tr[k], tr[j]
		}
	}
}
