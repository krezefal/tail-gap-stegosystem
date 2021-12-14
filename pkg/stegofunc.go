package pkg

import (
	"bufio"
	"io"
	"os"
	"strings"
)

// EmbedMessage func embeds given message into file-container so this message becomes "hidden" inside text file
// It returns the number of bits which fit inside the container
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

		line, carriageReturnFlag := PrepareLine(line)

		result := embedInLine(line, message, bitCounter, carriageReturnFlag)
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

// ExtractMessage func extracts "hidden" message from file-container and returns it
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

// embedInLine func embeds current message symbol in the end of line
// bitCounter here acts as message index
func embedInLine(line string, message []byte, bitCounter int, carriageReturnFlag bool) string {

	if bitCounter >= len(message) {
		if carriageReturnFlag {
			return line + "\r\n"
		}
		return line + "\n"
	}

	if message[bitCounter] == 1 {
		if carriageReturnFlag {
			return line + " \r\n"
		}
		return line + " \n"
	} else {
		if carriageReturnFlag {
			return line + "\r\n"
		}
		return line + "\n"
	}
}

// extractFromLine func extracts current hidden message symbol from the end of line and returns it
func extractFromLine(containerString string) (messageBit byte) {
	if strings.HasSuffix(containerString, " \n") || strings.HasSuffix(containerString, " \r\n") {
		return 0x01
	} else {
		return 0x00
	}
}
