package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func fileLen(fileName string) (int, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return 0, nil
	}
	defer f.Close()

	data := make([]byte, 2048)
	sizeTotal := 0
	for {
		count, err := f.Read(data)
		if err != nil {
			if err != io.EOF {
				return 0, nil
			}
			break
		}
		sizeTotal += count
	}
	return sizeTotal, nil
}

func fileLenFromOSStat(fileName string) (int64, error) {
	fi, err := os.Stat(fileName)
	if err != nil {
		return 0, err
	}
	return fi.Size(), nil
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Please set file name")
	}
	fileName := os.Args[1]
	sizeTotal, err := fileLen(fileName)
	if err != nil {
		log.Fatal("cloud not get file size", err)
	}
	fmt.Println(sizeTotal)

	// get file size from os.Stat
	sizeStat, err := fileLenFromOSStat(fileName)
	if err != nil {
		log.Fatal("cloud not get file size from os.Stat", err)
	}
	fmt.Println(sizeStat)
}
