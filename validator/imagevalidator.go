package imagevalidator

import (
	"fmt"
	"github.com/henrymxu/imagevalidator/internal"
	"github.com/henrymxu/imagevalidator/internal/utils"
	"os/exec"
)

const ResultsFile = "%s/results.json"

type ImageValidator struct {
	darknetPath string
}

func New(darknetPath string) ImageValidator {
	return ImageValidator{darknetPath}
}

func (imgvld ImageValidator) ValidateImages(imgfolder string, imgformat string, overwrite bool) {
	resultsPath := fmt.Sprintf(ResultsFile, imgfolder)
	imgList := utils.GetImageList(imgfolder, imgformat)
	imgList = utils.FilterImageList(utils.ReadJsonResultsFile(resultsPath), imgList)
	buffer := utils.ConvertListToByteBuffer(imgList)
	flags := []string{"detector", "validateSet", "cfg/coco.data", "cfg/yolov3.cfg", "cfg/yolov3.weights", resultsPath}
	darknet := exec.Command("./darknet",
		flags...)
	darknet.Dir = imgvld.darknetPath
	out, _ := darknet.StdoutPipe()
	darknet.Stdin = &buffer

	_ = darknet.Start()
	utils.DisplayCommandProgress(out)
	_ = darknet.Wait()

	internal.ParseResults(resultsPath, imgfolder)
}
