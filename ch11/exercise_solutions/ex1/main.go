package main

import (
	"errors"
	"fmt"
	"io"
	"os"
)

var FILE_POST_FIX = "_rights.txt"

func existsFile(file string) bool {
	_, err := os.Stat(file)
	return err == nil
}

func catFile(file string) error {
	fmt.Println(file)
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	buf := make([]byte, 2048)
	for {
		cnt, err := f.Read(buf)
		os.Stdout.Write(buf[:cnt])
		if err != nil {
			if err != io.EOF {
				return err
			}
			break
		}
	}

	return nil
}

func printUDHR(lang string) error {
	file := fmt.Sprintf("%s%s", lang, FILE_POST_FIX)
	if existsFile(file) {
		if err := catFile(fmt.Sprintf("%s%s", lang, FILE_POST_FIX)); err != nil {
			return err
		}
	} else {
		return errors.New(
			fmt.Sprintf("languege %s is not support", lang))
	}

	return nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "usage: %s <languege>", os.Args[0])
		os.Exit(1)
	}

	err := printUDHR(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
	}
}
