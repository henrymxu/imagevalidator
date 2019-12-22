package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func GetImageList(imageFolder string, imageFormat string) []string {
	files, err := ioutil.ReadDir(imageFolder)
	check(err)
	list := make([]string, len(files))
	actualSize := 0
	for _, f := range files {
		if strings.Contains(f.Name(), imageFormat) {
			list[actualSize] = fmt.Sprintf("%s/%s", imageFolder, f.Name())
			actualSize++
		}
	}
	return list[:actualSize]
}

func FilterImageList(prevResults []map[string]interface{}, imgList []string) []string {
	validatedImages := make(map[string]bool)
	for _, result := range prevResults {
		validatedImages[result["Path"].(string)] = true
	}
	i := 0
	for _, fileName := range imgList {
		if _, ok := validatedImages[fileName]; !ok {
			imgList[i] = fileName
			i++
		}
	}
	return imgList[:i]
}

func ConvertListToByteBuffer(list []string) bytes.Buffer {
	buffer := bytes.Buffer{}
	for _, value := range list {
		buffer.WriteString(fmt.Sprintf("%s\n", value))
	}
	return buffer
}

func ReadJsonResultsFile(path string) []map[string]interface{} {
	data, err := ioutil.ReadFile(path)
	check(err)
	var formatted []map[string]interface{}
	err = json.Unmarshal(data, &formatted)
	check(err)
	return formatted
}

func MakeFolder(path string) {
	err := os.MkdirAll(path, os.ModePerm)
	check(err)
}

func MoveFile(originalPath string, newPath string) {
	_ = os.Rename(originalPath, newPath)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
