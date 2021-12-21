package main

import (
	"fmt"
	. "github.com/krezefal/tail-gap-stegosystem/pkg"
)

var srcFileName = "examples/example.html"
var tgFileName = "examples/message_inside.html"

func main() {
	fmt.Println("\nBegin encoding...")

	srcFile, _ := OpenFile(srcFileName)
	defer srcFile.Close()

	tgFile, _ := CreateFile(tgFileName)
	defer tgFile.Close()

	message := GenerateMessage("10100")
	//message := GenerateRandomMessage(5)
	fmt.Println("Message:", message)
	fmt.Println("Begin embed message into file-container...")
	bitCounter, _ := EmbedMessage(srcFile, message, tgFile)
	fmt.Println("Number of message bits inside file-container:", bitCounter)
	fmt.Println("Message embedding completed.")
	tgFile.Close()

	fmt.Println("\nBegin decoding...")

	tgFile, _ = OpenFile(tgFileName)

	fmt.Println("Begin extracting message from file-container...")
	messageOut, _ := ExtractMessage(tgFile)
	fmt.Println("Extracted message:", messageOut)
	fmt.Println("Extracted message length:", len(messageOut))
	fmt.Println("Message extracting completed.")
}
