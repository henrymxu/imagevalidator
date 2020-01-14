package main

import (
	"github.com/henrymxu/imagevalidator/internal/utils"
	"github.com/henrymxu/imagevalidator/validator"
)

func main() {
	params := utils.GetInitialParams()
	validator := imagevalidator.New(utils.ParsePath(params.DarknetPath))
	validator.ValidateImages(utils.ParsePath(params.Folder), params.ImageFormat, params.Overwrite)
}
