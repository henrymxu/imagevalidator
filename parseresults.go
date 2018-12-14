package imagevalidator

import (
	"fmt"
	"github.com/henrymxu/imagevalidator/utils/storage"
	"math"
	"sort"
)

const SIZE_DIFF_THRESHOLD = 80

func ParseResults(path string) {
	results := storage.ReadJsonResultsFile(path)
	for index, v := range results {
		fmt.Printf("key[%d] value[%s]\n", index, v)
	}
}

//TODO image focus, number of classifications
func determineImageFocus(image map[string]interface{}) bool {
	detections := image["Detections"].([]map[string]interface{})
	for _, detection := range detections {
		box := detection["Box"].([]int)
		area := math.Abs(float64(box[0] - box[2])) * math.Abs(float64(box[1] - box[3]))
		detection["Area"] = area
	}
	sortedKeys := sortDetections(detections, "Area")
	if len(detections) > 1 {
		return isSingleFocus(detections[sortedKeys[0]]["Area"].(int), detections[sortedKeys[1]]["Area"].(int))
	}

	return true
}

func sortDetections(detections []map[string]interface{}, key string) []int {
	type kv struct {
		Key   int
		Value float64
	}

	var ss []kv
	for k, v := range detections {
		ss = append(ss, kv{k, v[key].(float64)})
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value
	})

	var keys []int
	for _, v := range ss {
		keys = append(keys, v.Key)
	}

	return keys
}

func isSingleFocus(detection1 int, detection2 int) bool {
	diff := float64(detection1 - detection2)
	return (diff / float64(detection1)) * 100 > SIZE_DIFF_THRESHOLD
}