package utils

import (
	"bufio"
	"flag"
	"fmt"
	"io"
)

func DisplayCommandProgress(closer io.ReadCloser) {
	scanner := bufio.NewScanner(closer)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		m := scanner.Text()
		fmt.Println(m)
	}
}

type Params struct {
	DarknetPath string
	Folder      string
	ImageFormat string
	Overwrite   bool
}

func GetInitialParams() Params {
	darknetPath := flag.String("darknet", "", "darknet path")
	folder := flag.String("folder", "", "name of the folder")
	imageFmt := flag.String("format", ".jpg", "image file type")
	overwrite := flag.Bool("overwrite", false, "overwrite previous validations")

	flag.Parse()
	return Params{
		*darknetPath,
		*folder,
		*imageFmt,
		*overwrite,
	}
}
