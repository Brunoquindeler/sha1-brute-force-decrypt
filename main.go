package main

import (
	"bufio"
	"crypto/sha1"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/fatih/color"
)

// colors
var (
	blue   = color.New(color.Bold, color.FgBlue)
	yellow = color.New(color.Bold, color.FgYellow)
)

// flags
var (
	commonPasswordsFilePath string
	help                    bool
)

const helpMessage = `Usage:
	Flags:
		-f [ text ]    | especified file with commons passwords
		-h [ boolean ] | help mode

	Examples:
		sha1-brute-force-decrypt -f="common_passwords.txt"
		sha1-brute-force-decrypt -h
`

func main() {
	flag.StringVar(&commonPasswordsFilePath, "f", "", "especified file with commons passwords")
	flag.BoolVar(&help, "h", false, "help mode")
	flag.Parse()

	if help || commonPasswordsFilePath == "" {
		fmt.Print(helpMessage)
		return
	}

	file, err := os.Open(commonPasswordsFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	inMemoryHashesMap := make(map[string]string)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		password := scanner.Text()
		hash := fmt.Sprintf("%x", sha1.Sum([]byte(password)))
		inMemoryHashesMap[hash] = password
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for {
		var hash string
		blue.Print("Hash: ")
		_, err := fmt.Scanln(&hash)
		if err != nil {
			log.Fatal(err)
		}

		v, ok := inMemoryHashesMap[hash]
		if ok {
			yellow.Printf("Hash [ %s ]: %s\n", hash, color.GreenString(v))
			continue
		}

		yellow.Printf("Hash [ %s ]: %s\n", hash, color.RedString("not found"))
	}
}
