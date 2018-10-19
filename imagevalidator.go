package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
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
	createImageListFile(imgfolder, imgformat)
	imgcmd := exec.Command("cat", fmt.Sprintf("%s/images.txt", imgfolder))
	darknet := exec.Command("./darknet",
		strings.Fields(fmt.Sprintf("validate cfg/yolov3.cfg cfg/yolov3.weights %s/results.json", imgfolder))...)
	darknet.Dir = imgvld.darknetpath
	out, _ := darknet.StdoutPipe()

	pipe, _ := imgcmd.StdoutPipe()
	defer pipe.Close()
	darknet.Stdin = pipe

	darknet.Start()
	imgcmd.Run()

	scanner := bufio.NewScanner(out)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		m := scanner.Text()
		fmt.Println(m)
	}

	darknet.Wait()
}

func createImageListFile(imgfolder string, imgformat string) {
	output, err := os.Create(fmt.Sprintf("%s/images.txt", imgfolder))
	handleError(err)
	defer output.Close()
	files, err := ioutil.ReadDir(imgfolder)
	handleError(err)
	for _, f := range files {
		if strings.Contains(f.Name(), imgformat) {
			output.WriteString(fmt.Sprintf("%s/%s\n", imgfolder, f.Name()))
		}
	}
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}