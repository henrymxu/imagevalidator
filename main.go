package main

import (
	"github.com/henrymxu/imagevalidator/internal/utils"
	"github.com/henrymxu/imagevalidator/validator"
)

func main() {
	params := utils.GetInitialParams()
	validator := imagevalidator.New(params.DarknetPath)
	validator.ValidateImages(params.Folder, params.ImageFormat, params.Overwrite)
}
