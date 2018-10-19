package cmd

import (
	"bufio"
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