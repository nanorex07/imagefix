package main

import (
	"log"
	"os"

	"github.com/nanorex07/imagefix/filters"
	"github.com/nanorex07/imagefix/utils"
)

func main() {
	fileName := os.Args[1]
	img, err := utils.OpenImage(fileName)
	if err != nil {
		log.Fatal(err)
	}
	img_tensor := utils.ConvertImageTensor(img)

	//perform here
	filters.GrayScale(&img_tensor)
	filters.GaussianBlur(&img_tensor)

	img = utils.ConvertTensorImage(img_tensor)
	err = utils.SaveImage("output.png", &img)

	if err != nil {
		log.Fatal(err)
	}
}
