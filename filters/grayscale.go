package filters

import (
	"image/color"
	"sync"
)

func GrayScale(pixels *[][]color.Color) {
	pinitial := *pixels
	rows := len(pinitial)
	cols := len(pinitial[0])

	newImage := make([][]color.Color, rows)
	for i := 0; i < rows; i++ {
		newImage[i] = make([]color.Color, cols)
	}
	wg := sync.WaitGroup{}
	for x := 0; x < rows; x++ {
		for y := 0; y < cols; y++ {
			wg.Add(1)
			go func(x, y int) {

				pixel := pinitial[x][y]
				original_color, _ := color.RGBAModel.Convert(pixel).(color.RGBA)
				gray := uint8(float64(original_color.R)*0.299 + float64(original_color.G)*0.587 + float64(original_color.B)*0.114)
				newImage[x][y] = color.RGBA{
					gray,
					gray,
					gray,
					original_color.A,
				}
				wg.Done()
			}(x, y)
		}
	}
	wg.Wait()
	*pixels = newImage
}
