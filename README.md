# Tail Gap Stegosystem (Text Steganography)

## Description

The tail gap method assumes appending one space at the end of each line of the 
file-container in case of encoding a `1` bit of the stegano message. 
In case of encoding a `0` bit, the space at the end of the line is not appended.

While processing, the text file is read line by line. All whitespace characters 
(spaces, carriage return and newline symbols) are removed from the end of the line 
and then, depending on the value of the current bit of the stegano message 
represented in binary form, a decision of appending space or not is made.
Lines converted in this way is written to the result file.

The bandwidth of such a cryptosystem is rather low, cause the throughput of this 
method is highly dependent on the file-container lines amount. Accepting
the file-container format described above, it is possible to calculate the 
throughput of this steganosystem. It's just about ~0.15%

## Implementation

This program is a Go module that provides message embedding and extracting procedure
to/from file-container.

To build the source, execute

`go build -o bin/tail-gap-stegosystem cmd/main.go`

Or just run it w/o building the binary:

`go run cmd/main.go`

Program also has unit test for most difficult to understand function :D

Folder with examples contains a couple of files on which program can be tested 
"out of box".
