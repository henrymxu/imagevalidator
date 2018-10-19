package main

import "github.com/henrymxu/imagevalidator"

func main() {
	imgvldr := imagevalidator.New("/Users/henryxu/Documents/darknet")
	imgvldr.ValidateImages("/Users/henryxu/Documents/ImageSet/corvette", ".jpg")
	imagevalidator.ParseResults("/Users/henryxu/Documents/ImageSet/corvette/results.json")
}
