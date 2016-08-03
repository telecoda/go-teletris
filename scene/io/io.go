package io

import (
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"log"

	"golang.org/x/mobile/asset"
)

func LoadImage(assetName string) (image.Image, string, error) {

	a, err := asset.Open(assetName)
	if err != nil {
		log.Printf("Error during LoadImage: %s", err)
		return nil, "", err
	}
	defer a.Close()
	loadedImage, format, err := image.Decode(a)

	return loadedImage, format, err

}
