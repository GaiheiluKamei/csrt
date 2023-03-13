package main

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func correctSRT(input, output, delimiter string) error {
	if strings.Contains(delimiter, "r") {
		delimiter = "\r\n"
	} else {
		delimiter = "\n"
	}

	fmt.Printf("%#v", delimiter)
	f, err := os.ReadFile(input)
	if err != nil {
		return err
	}

	result := [][]byte{}
	chunks := bytes.Split(f, []byte(strings.Repeat(delimiter, 2)))

	for _, chunk := range chunks {
		subChunk := bytes.Split(chunk, []byte(delimiter))
		if len(subChunk) == 5 {
			result = append(result, bytes.Join(subChunk[1:4], []byte(delimiter)))
		}

		if len(subChunk) == 4 {
			result = append(result, bytes.Join(subChunk[:3], []byte(delimiter)))
		}
	}

	if output == "" {
		output = fmt.Sprintf("srt_%s", filepath.Base(input))
	}

	return os.WriteFile(output, bytes.Join(result, []byte(strings.Repeat(delimiter, 2))), 0644)
}
