package pkg

import (
	"math/rand"
	"os"
	"time"
)

// OpenFile opens the file.
func OpenFile(fileName string) (*os.File, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	return file, nil
}

// CreateFile creates the file.
func CreateFile(fileName string) (*os.File, error) {
	file, err := os.Create(fileName)
	if err != nil {
		return nil, err
	}
	return file, nil
}

// GenerateMessage generates specified binary sequence.
func GenerateMessage(sequence string) []byte {
	message := make([]byte, len(sequence))

	for i, bit := range sequence {
		if bit == '1' {
			message[i] = 0x01
		} else {
			message[i] = 0x00
		}
	}

	return message
}

// GenerateRandomMessage generates random binary sequence with specified length.
func GenerateRandomMessage(length int) []byte {
	rMessage := make([]byte, length)

	seed := rand.NewSource(time.Now().UnixNano())
	r := rand.New(seed)

	for i := 0; i < cap(rMessage); i++ {
		if r.Int()%2 == 1 {
			rMessage[i] = 0x00
		} else {
			rMessage[i] = 0x01
		}
	}

	return rMessage
}

// PrepareLine prepares container-line by removing all spaces at the end of string with newline and carriage return
// symbols. Otherwise, this sequence will interfere correct message extracting during the extracting session. Returns
// prepared line and carriage return presence flag.
func PrepareLine(line string) (string, bool) {

	var carriageReturnFlag bool

	if line == "\n" {
		return "", carriageReturnFlag
	}

	idx := len(line) - 2

	if line[idx] == '\r' {
		carriageReturnFlag = true
		idx--
	}
	for ; idx >= 0; idx-- {
		if line[idx] != ' ' {
			break
		}
	}

	return line[:idx+1], carriageReturnFlag
}
