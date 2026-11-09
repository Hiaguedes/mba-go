package main

import (
	"bufio"
	"os"
)

func main() {
	file, err := os.Create("07-files/file.txt")

	if err != nil {
		panic(err)
	}

	lengthWritten, err := file.WriteString( "Hello, World!") // WriteString is used to write a string to the file. It takes the file and the string as arguments and returns the number of bytes written and an error if any.

	if err != nil {
		panic(err)
	}

	println("Bytes written:", lengthWritten) // Print the number of bytes written to the file.

	defer file.Close() // defer is used to ensure that the file is closed when the function exits, even if an error occurs. It is a common practice to use defer for resource cleanup in Go.

	// leitura

	bytesRead, err := os.ReadFile("07-files/file.txt") // Open is used to open a file for reading. It takes the file path as an argument and returns a file object and an error if any.

	if err != nil {
		panic(err)
	}

	println("Content:", string(bytesRead)) // Print the content read from the file. The buffer is sliced to the number of bytes read to avoid printing any extra bytes that may be in the buffer.

	// stream reading

	streamFile, err := os.Open("07-files/file.txt") // Open is used to open a file for reading. It takes the file path as an argument and returns a file object and an error if any.

	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(streamFile) // NewReader is used to create a new buffered reader. It takes the file object as an argument and returns a buffered reader.

	buffer := make([]byte, 1) // Create a buffer to hold the bytes read from the file. The size of the buffer is set to 5 bytes.

	for {
		b, err := reader.Read(buffer) // Read is used to read bytes from the buffered reader. It takes the buffer as an argument and returns the number of bytes read and an error if any.

		if err != nil {
			break
		}

		println("Byte read:", string(buffer[:b])) // Print the byte read from the file. The buffer is sliced to the number of bytes read to avoid printing any extra bytes that may be in the buffer.
	}
	
	defer streamFile.Close() // Ensure the file is closed when the function exits.

}