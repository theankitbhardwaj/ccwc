package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode/utf8"
)

func main() {
	byteCountFlagPtr := flag.Bool("c", false, "print the byte counts")
	newLineCountPtr := flag.Bool("l", false, "print the newline counts")
	characterCountPtr := flag.Bool("m", false, "print the character counts")
	wordCountPtr := flag.Bool("w", false, "print the word counts")
	flag.Parse()

	filePath := os.Args[len(os.Args)-1]
	f, err := os.Open(filePath)

	if err != nil {
		fmt.Printf("E#1WQ6IT: Error opening file: %v\n", err)
		os.Exit(1)
	}

	defer f.Close()

	if *newLineCountPtr {
		totalLines, err := lineCounter(f)
		if err != nil {
			panic(fmt.Sprintf("P#1WQAA0: %v\n", err))
		}
		fmt.Printf("%v ", totalLines)
	}

	if *wordCountPtr {
		fb, err := os.ReadFile(f.Name())
		totalWords := wordCounter(fb)
		if err != nil {
			panic(fmt.Sprintf("P#1WQAA0: %v\n", err))
		}
		fmt.Printf("%v ", totalWords)
	}

	if *characterCountPtr {
		fb, err := os.ReadFile(f.Name())
		totalChars := charaterCounter(fb)
		if err != nil {
			panic(fmt.Sprintf("P#1WQAG2: %v\n", err))
		}
		fmt.Printf("%v ", totalChars)
	}

	if *byteCountFlagPtr {
		totalBytes := byteCounter(*f)
		fmt.Printf("%v ", totalBytes)
	}

	if !*byteCountFlagPtr && !*newLineCountPtr && !*wordCountPtr && !*characterCountPtr {
		printAll(f)
	}

	fmt.Printf("%v\n", f.Name())

}

func printAll(f *os.File) {
	// totalLines
	totalLines, err := lineCounter(f)
	if err != nil {
		panic(fmt.Sprintf("P#1WQAA0: %v\n", err))
	}

	// totalWords
	fb, err := os.ReadFile(f.Name())
	totalWords := wordCounter(fb)
	if err != nil {
		panic(fmt.Sprintf("P#1WQAA0: %v\n", err))
	}

	// totalBytes
	totalBytes := byteCounter(*f)

	fmt.Printf("%v %v %v ", totalLines, totalWords, totalBytes)

}

func byteCounter(f os.File) int64 {
	fileStat, err := f.Stat()

	if err != nil {
		panic(fmt.Sprintf("P#1WQ8J8: %v", err))
	}

	totalBytes := fileStat.Size()

	return totalBytes
}

func lineCounter(r io.Reader) (int, error) {
	buf := make([]byte, 32*1024)
	count := 0
	lineSep := []byte{'\n'}

	for {
		ch, err := r.Read(buf)
		count += bytes.Count(buf[:ch], lineSep)

		switch {
		case err == io.EOF:
			return count, nil
		case err != nil:
			return count, err
		}
	}
}

func charaterCounter(f []byte) int64 {
	return int64(utf8.RuneCount(f))
}

func wordCounter(f []byte) int64 {
	words := strings.Fields(string(f))

	return int64(len(words))
}
