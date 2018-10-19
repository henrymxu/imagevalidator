package imagevalidator

import (
	"fmt"
	"github.com/henrymxu/imagevalidator/utils/cmd"
	"github.com/henrymxu/imagevalidator/utils/storage"
	"os/exec"
	"strings"
)

type imagevalidator struct {
	darknetpath string
}

func New(darknetpath string) imagevalidator {
	return imagevalidator{darknetpath}
}

func (imgvld imagevalidator) ValidateImages(imgfolder string, imgformat string) {
	buffer := storage.GetImageListAsByteBuffer(imgfolder, imgformat)
	darknet := exec.Command("./darknet",
		strings.Fields(fmt.Sprintf("validate cfg/yolov3.cfg cfg/yolov3.weights %s/results.json", imgfolder))...)
	darknet.Dir = imgvld.darknetpath
	out, _ := darknet.StdoutPipe()
	darknet.Stdin = &buffer

	darknet.Start()
	cmd.DisplayCommandProgress(out)
	darknet.Wait()
}
