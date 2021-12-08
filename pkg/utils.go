package pkg

import (
	"os"
)

// OpenFile pkg opens file
func OpenFile(fileName string) (*os.File, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	return file, nil
}

// CreateFile pkg creates file
func CreateFile(fileName string) (*os.File, error) {
	file, err := os.Create(fileName)
	if err != nil {
		return nil, err
	}
	return file, nil
}

// GenerateMessage pkg generates either random or specified binary sequence with specified length
// TODO: random generation mode + specified sequence
func GenerateMessage(length int) []byte {
	message := make([]byte, length)

	for i := 0; i < cap(message); i++ {
		if i % 2 == 1 {
			message[i] = 0x00
		} else {
			message[i] = 0x01
		}
	}

	return message
}

// PrepareLine pkg prepares container-line by removing all spaces at the end of string and newline symbol
// this sequence will interfere correct message extracting during the extracting session
func PrepareLine(line string) string {

	if line == "\n" {
		return ""
	}

	var cutIndex int

	for cutIndex = len(line) - 2; cutIndex >= 0; cutIndex-- {
		if line[cutIndex] != ' ' {
			break
		}
	}

	return line[:cutIndex + 1]
}
