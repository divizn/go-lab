package main

import (
	"bufio"
	"crypto/sha256"
	"fmt"
	"io"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: metafile <command> [args]")
		fmt.Println("Commands: count-lines, count-words, size, checksum")
		os.Exit(1)
	}

	cmd := os.Args[1]

	switch cmd {
	case "count-lines":
		if len(os.Args) < 3 {
			fmt.Println("Usage: metafile count-lines <file>")
			return
		}
		countLines(os.Args[2])
	case "count-words":
		if len(os.Args) < 3 {
			fmt.Println("Usage: metafile count-words <file>")
			return
		}
		countWords(os.Args[2])
	case "size":
		if len(os.Args) < 3 {
			fmt.Println("Usage: metafile size <file>")
			return
		}
		fileSize(os.Args[2])
	case "checksum":
		if len(os.Args) < 3 {
			fmt.Println("Usage: metafile checksum <file>")
			return
		}
		checksum(os.Args[2])
	default:
		fmt.Println("Unknown command:", cmd)
	}
}

func countLines(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := 0
	for scanner.Scan() {
		lines++
	}
	fmt.Printf("%s: %d lines\n", filename, lines)
}

func countWords(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	words := 0
	for scanner.Scan() {
		words++
	}
	fmt.Printf("%s: %d words\n", filename, words)
}

func fileSize(filename string) {
	info, err := os.Stat(filename)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("%s: %d bytes\n", filename, info.Size())
}

func checksum(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	h := sha256.New()
	if _, err := io.Copy(h, file); err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("%s: %x\n", filename, h.Sum(nil))
}
