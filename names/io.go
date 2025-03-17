package names

import (
	"bufio"
	"log"
	"os"
)

func ReadLines(name string, lines chan string) {
	defer close(lines)

	file, fileOpenError := os.Open(name)
	if fileOpenError != nil {
		log.Panicf("Error opening file: %v", fileOpenError)
	}
	defer func(file *os.File) {
		fileCloseError := file.Close()
		if fileCloseError != nil {
			log.Panicf("Error closing file: %v", fileCloseError)
		}
	}(file)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines <- scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		log.Panicf("Error scanning file: %v", err)
	}
}
