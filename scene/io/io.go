package io

import (
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"
)

func LoadImage(imagePath string) (image.Image, string, error) {

	file, err := os.Open(imagePath)
	if err != nil {
		log.Printf("Error during LoadImage: %s", err)
		return nil, "", err
	}
	defer file.Close()
	loadedImage, format, err := image.Decode(file)

	return loadedImage, format, err

}
