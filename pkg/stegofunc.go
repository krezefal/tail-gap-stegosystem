package pkg

import (
	"bufio"
	"io"
	"os"
	"strings"
)

// EmbedMessage pkg embeds given message into file-container so this message becomes "hidden" inside text file
// it returns the number of bits which fit inside the container
func EmbedMessage(f *os.File, message []byte, file *os.File) (int, error) {
	reader := bufio.NewReader(f)
	var bitCounter int

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return -1, err
			}

		}

		line = PrepareLine(line)

		result := embedInLine(line, message, bitCounter)
		_, err = file.WriteString(result)
		if err != nil {
			return -1, err
		}

		if bitCounter < len(message) {
			bitCounter++
		}
	}

	return bitCounter, nil
}

// ExtractMessage pkg extracts "hidden" message from file-container and returns it
func ExtractMessage(file *os.File) ([]byte, error) {
	reader := bufio.NewReader(file)

	var message []byte

	for {
		line, err := reader.ReadString('\n')

		if err != nil {
			if err == io.EOF {
				break
			} else {
				return nil, err
			}

		}

		message = append(message, extractFromLine(line))
	}

	return message, nil
}

// embedInLine pkg embeds current message symbol in the end of line
// bitCounter here acts as message index
func embedInLine(line string, message []byte, bitCounter int) string {

	if bitCounter >= len(message) {
		line = line + "\n"
		return line
	}

	if message[bitCounter] == 1 {
		line = line + " \n"
		return line
	} else {
		line = line + "\n"
		return line
	}
}

// extractFromLine pkg extracts current hidden message symbol from the end of line and returns it
func extractFromLine(containerString string) (messageBit byte) {
	if strings.HasSuffix(containerString, " \n") {
		return 0x01
	} else {
		return 0x00
	}
}
