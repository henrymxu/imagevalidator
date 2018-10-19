package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
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
	buffer := getImageListAsByteBuffer(imgfolder, imgformat)
	darknet := exec.Command("./darknet",
		strings.Fields(fmt.Sprintf("validate cfg/yolov3.cfg cfg/yolov3.weights %s/results.json", imgfolder))...)
	darknet.Dir = imgvld.darknetpath
	out, _ := darknet.StdoutPipe()
	darknet.Stdin = &buffer

	darknet.Start()

	scanner := bufio.NewScanner(out)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		m := scanner.Text()
		fmt.Println(m)
	}

	darknet.Wait()
}

func getImageListAsByteBuffer(imgfolder string, imgformat string) bytes.Buffer {
	buffer := bytes.Buffer{}
	files, err := ioutil.ReadDir(imgfolder)
	handleError(err)
	for _, f := range files {
		if strings.Contains(f.Name(), imgformat) {
			buffer.WriteString(fmt.Sprintf("%s/%s\n", imgfolder, f.Name()))
		}
	}
	return buffer
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}