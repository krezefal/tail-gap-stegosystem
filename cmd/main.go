package main

import (
	"fmt"
	. "github.com/krezefal/tail-gap-stegosystem/pkg"
)

func main() {
	fmt.Println("Begin encoding...")

	srcFile, _ := OpenFile("examples/example.txt")
	defer srcFile.Close()

	tgFile, _ := CreateFile("examples/message_inside.txt")
	defer tgFile.Close()

	message := GenerateMessage(100)
	fmt.Println("Message:", message)
	fmt.Println("Begin embed message into file-container...")
	bitCounter, _ := EmbedMessage(srcFile, message, tgFile)
	fmt.Println("Message bit number inside file-container:", bitCounter)
	fmt.Println("Message embedding completed.")
	tgFile.Close()

	fmt.Println("Begin decoding...")

	tgFile, _ = OpenFile("examples/message_inside.txt")

	fmt.Println("Begin extracting message from file-container...")
	messageOut, _ := ExtractMessage(tgFile)
	fmt.Println("Extracted message:", messageOut)
	fmt.Println("Message extracting completed.")
}

