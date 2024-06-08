package utils

import (
	"log"
	"os"
)

func ReadFile(file string) string {
	var content, err = os.ReadFile(file)
	CheckErr(err)
	return string(content)
}

func CheckErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func CheckArgs() {
	if len(os.Args) != 2 {
		log.Fatalln("Usage: go run main.go <file>")
	}
}
