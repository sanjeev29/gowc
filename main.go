package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

func readBytes(file *os.File) int {
	numberOfBytes := 0

	reader := bufio.NewReader(file)

	for {
		_, size, err := reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				break
			}

			log.Fatalln(err)
		}

		numberOfBytes += size
	}

	return numberOfBytes
}

func main() {
	var bytes bool

	flag.BoolVar(&bytes, "c", false, "number of bytes")

	flag.Parse()

	filename := flag.CommandLine.Arg(0)

	_, err := os.Stat(filename)
	if err != nil {
		log.Fatal(err)
	}

	// Open file
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	// Close file
	defer file.Close()

	var numberOfBytes int

	if bytes {
		numberOfBytes = readBytes(file)
	}

	// Print number of bytes
	fmt.Printf("  %d %s\n", numberOfBytes, filename)
}
