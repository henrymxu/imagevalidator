package storage

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
)

func GetImageListAsByteBuffer(imgfolder string, imgformat string) bytes.Buffer {
	buffer := bytes.Buffer{}
	files, err := ioutil.ReadDir(imgfolder)
	check(err)
	for _, f := range files {
		if strings.Contains(f.Name(), imgformat) {
			buffer.WriteString(fmt.Sprintf("%s/%s\n", imgfolder, f.Name()))
		}
	}
	return buffer
}

func ReadJsonResultsFile(path string) []map[string]interface{} {
	data, err := ioutil.ReadFile(path)
	check(err)
	var formatted []map[string]interface{}
	json.Unmarshal(data, &formatted)
	return formatted
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
