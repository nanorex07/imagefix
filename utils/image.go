package utils

import (
	"image"
	_ "image/jpeg"
	"image/png"
	_ "image/png"
	"log"
	"os"
)

func OpenImage(path string) (image.Image, error) {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	img, _, err := image.Decode(f)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return img, nil
}

func SaveImage(path string, img *image.Image) error {
	file, err := os.Create(path)
	defer file.Close()
	if err != nil {
		return err
	}
	png.Encode(file, *img)
	return nil
}
