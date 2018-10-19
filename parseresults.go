package imagevalidator

import (
	"fmt"
	"github.com/henrymxu/imagevalidator/utils/storage"
)

func ParseResults(path string) {
	results := storage.ReadJsonResultsFile(path)
	for index, v := range results {
		fmt.Printf("key[%d] value[%s]\n", index, v)
	}
}

//TODO image focus, number of classifications
func determineImageFocus(image map[string]interface{}) {
	detections := image["Detections"].([]map[string]interface{})
	for _, detection := range detections {
		box := detection["Box"].([]int)
		fmt.Println(box)
	}
}