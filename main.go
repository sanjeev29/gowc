package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"unicode"
)

type fileStats struct {
	bytes      int
	lines      int
	words      int
	characters int
}

func getFileStats(file *os.File) fileStats {
	var numberOfBytes, numberOfLines, numberOfWords, numberOfChars int

	reader := bufio.NewReader(file)

	isWord := false
	for {
		r, size, err := reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				if isWord {
					numberOfWords += 1
				}
				break
			}

			log.Fatalln(err)
		}

		if unicode.IsSpace(r) {
			if isWord {
				numberOfWords += 1
			}

			if r == '\n' {
				numberOfLines += 1
			}

			isWord = false
		} else {
			isWord = true
		}

		numberOfBytes += size
		numberOfChars += 1
	}

	return fileStats{
		bytes:      numberOfBytes,
		lines:      numberOfLines,
		words:      numberOfWords,
		characters: numberOfChars,
	}
}

func main() {
	var byteCount, lineCount, wordCount, charCount bool

	flag.BoolVar(&byteCount, "c", false, "number of bytes")
	flag.BoolVar(&lineCount, "l", false, "number of lines")
	flag.BoolVar(&wordCount, "w", false, "number of words")
	flag.BoolVar(&charCount, "m", false, "number of characters")

	flag.Parse()

	if !byteCount && !lineCount && !wordCount && !charCount {
		byteCount = true
		lineCount = true
		wordCount = true
	}

	filename := flag.CommandLine.Arg(0)

	fileStat, err := os.Stdin.Stat()
	if err != nil {
		log.Fatal(err)
	}

	var file *os.File

	if (fileStat.Mode() & os.ModeCharDevice) == 0 {
		// String is being piped to stdin
		file = os.Stdin
	} else {
		// Open file
		file, err = os.Open(filename)
		if err != nil {
			log.Fatal(err)
		}

		// Close file
		defer file.Close()
	}

	stats := getFileStats(file)

	printArgs := []string{}

	if lineCount {
		printArgs = append(printArgs, fmt.Sprintf("%*d", 8, stats.lines))
	}

	if wordCount {
		printArgs = append(printArgs, fmt.Sprintf("%*d", 8, stats.words))
	}

	if byteCount {
		printArgs = append(printArgs, fmt.Sprintf("%*d", 8, stats.bytes))
	}

	if charCount {
		printArgs = append(printArgs, fmt.Sprintf("%*d", 8, stats.characters))
	}

	printArgs = append(printArgs, filename)
	fmt.Println(strings.Join(printArgs, " "))
}
