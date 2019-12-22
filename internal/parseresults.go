package internal

import (
	"fmt"
	"github.com/henrymxu/imagevalidator/internal/utils"
	"math"
	"path/filepath"
	"sort"
)

const SizeDiffThreshold = 80

func ParseResults(resultsPath string, path string) {
	results := utils.ReadJsonResultsFile(resultsPath)
	index := 0
	for _, v := range results {
		if !imageHasSingleFocus(v) {
			results[index] = v
			index++
		}
	}
	results = results[:index]

	fmt.Printf("%d images are invalid\n", len(results))
	invalid := fmt.Sprintf("%s/invalid", path)
	invalidLabelled := fmt.Sprintf("%s/invalidLabels", invalid)
	utils.MakeFolder(invalidLabelled)
	for _, v := range results {
		file := v["Path"].(string)
		base := filepath.Base(file)
		imagePath := fmt.Sprintf("%s/%s", path, base)
		labelledImagePath := fmt.Sprintf("%s/predictions/%s", filepath.Dir(file), base)
		utils.MoveFile(imagePath, fmt.Sprintf("%s/%s", invalid, base))
		utils.MoveFile(labelledImagePath, fmt.Sprintf("%s/%s", invalidLabelled, base))
	}
}

//TODO image focus, number of classifications
func imageHasSingleFocus(image map[string]interface{}) bool {
	detections := image["Detections"].([]interface{})
	for _, detection := range detections {
		box := detection.(map[string]interface{})["Box"].([]interface{})
		area := math.Abs(box[0].(float64)-box[2].(float64)) * math.Abs(box[1].(float64)-box[3].(float64))
		detection.(map[string]interface{})["Area"] = area
	}
	sortedKeys := sortDetections(detections, "Area")
	if len(detections) > 1 {
		largestDetectionArea := detections[sortedKeys[0]].(map[string]interface{})["Area"].(float64)
		secondLargestDetectionArea := detections[sortedKeys[1]].(map[string]interface{})["Area"].(float64)
		return isSingleFocus(largestDetectionArea, secondLargestDetectionArea)
	}
	return true
}

func sortDetections(detections []interface{}, key string) []int {
	type kv struct {
		Key   int
		Value float64
	}

	var ss []kv
	for k, v := range detections {
		ss = append(ss, kv{k, v.(map[string]interface{})[key].(float64)})
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

func isSingleFocus(detection1 float64, detection2 float64) bool {
	diff := detection1 - detection2
	return (diff/detection1)*100 > SizeDiffThreshold
}
