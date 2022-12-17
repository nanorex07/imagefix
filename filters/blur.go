package filters

import (
	"image/color"

	"gonum.org/v1/gonum/mat"
)

func BoxBlur3x3(pixels *[][]color.Color) {

	boxKernal := mat.NewDense(3, 3, []float64{
		1.0 / 9, 1.0 / 9, 1.0 / 9,
		1.0 / 9, 1.0 / 9, 1.0 / 9,
		1.0 / 9, 1.0 / 9, 1.0 / 9,
	})
	BoxBlur(pixels, boxKernal)
}

func GaussianBlur(pixels *[][]color.Color) {

	gaussianKernel := mat.NewDense(5, 5, []float64{
		1.0 / 256, 4.0 / 256, 6.0 / 256, 4.0 / 256, 1.0 / 256,
		4.0 / 256, 16.0 / 256, 24.0 / 256, 16.0 / 256, 4.0 / 256,
		6.0 / 256, 24.0 / 256, 36.0 / 256, 24.0 / 256, 6.0 / 256,
		4.0 / 256, 16.0 / 256, 24.0 / 256, 16.0 / 256, 4.0 / 256,
		1.0 / 256, 4.0 / 256, 6.0 / 256, 4.0 / 256, 1.0 / 256,
	})
	BoxBlur(pixels, gaussianKernel)
}

func BoxBlur(pixels *[][]color.Color, boxKernal *mat.Dense) {
	pini := *pixels

	rows, cols := boxKernal.Dims()
	offset := float64(rows / 2)
	klen := float64(cols)

	newImage := make([][]color.Color, len(pini))
	for i := 0; i < len(pini); i++ {
		newImage[i] = make([]color.Color, len(pini[i]))
	}
	copy(newImage, pini)

	for x := offset; x < float64(len(pini))-offset; x++ {
		for y := offset; y < float64(len(pini[0]))-offset; y++ {
			npix := color.RGBA{}
			for a := 0.0; a < klen; a++ {
				for b := 0.0; b < klen; b++ {
					xn := x + a - offset
					yn := y + a - offset
					r, g, bb, aa := pini[int(xn)][int(yn)].RGBA()
					npix.R += uint8(float64(uint8(r)) * (boxKernal.At(int(a), int(b))))
					npix.G += uint8(float64(uint8(g)) * (boxKernal.At(int(a), int(b))))
					npix.B += uint8(float64(uint8(bb)) * (boxKernal.At(int(a), int(b))))
					npix.A += uint8(float64(uint8(aa)) * (boxKernal.At(int(a), int(b))))

				}
			}
			newImage[int(x)][int(y)] = npix
		}
	}

	*pixels = newImage

}
