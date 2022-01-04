# Tail Gap Stegosystem (Text Steganography)

## Description

The **tail gap method** assumes appending one space at the end of each line of the 
file-container in case of encoding a `1` bit of the stegano message (watermark). In
case of encoding a `0` bit, the space at the end of the line is not appended.

While processing, the text file is read line by line. All whitespace characters 
(spaces, carriage return and newline symbols) are removed from the end of the line 
and then, depending on the value of the current watermark bit a decision of appending
space or not is made. Lines converted in this way is written to the result file.

## Implementation

This program is a Go module that provides message embedding and extracting procedure
to/from file-container.

To build the source, execute

`go build -o bin/tail-gap-stegosystem cmd/main.go`

Or just run it w/o building the binary:

`go run cmd/main.go`

Program also has unit test for most difficult to understand function :)

Folder with examples contains a couple of files on which program can be tested "out
of box".

### Using

To change the source or target files, overwrite `8-9` lines in `cmd/main.go`

### Principle of operation

**Embedding**

Watermark binary sequence is formed, after what, the container file is read line by
line. During the read process, all trailing spaces, carriage returns, and linebreaks
are removed from the current line. 

Depending on the current watermark bit, either a space with a newline symbol or just
a newline symbol is appended at the end of the prepared line. If there was a carriage
return symbol in the source line, it is also appended to the resulting line. 

If the watermark binary sequence has ended, and there are still lines in the file, 
only the trailing spaces are removed from them so that there are no erroneous `1` 
bits during the subsequent extracting procedure. If the binary sequence is larger
than the number of lines, it won't be fully embedded (only the first *N* bits, where
*N* is the number of lines).

As a result, the modified file-container is formed.

**Extracting**

Extraction of the stegano message is as follows: the program reads the file-container
line by line, and forms a binary sequence: `1`, in the case of a space at the end of
the line is detected and `0` in the absence. 

The watermark contained in the file-container may be less than the number of lines
in the file. In this case, the extracted bit sequence will contain all "zeros" in 
place of the extra lines of the file-container.

The extracted watermark message is displayed on the screen.

## Theory

The bandwidth of such a stegosystem is rather low, cause the throughput of this 
method is highly dependent on the file-container lines amount. Accepting the 
file-container format described above, it is possible to calculate the throughput of
this stegosystem. It's just about ~0.15%