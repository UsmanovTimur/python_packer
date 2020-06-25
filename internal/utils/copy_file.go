package utils

import (
	"io"
	"log"
	"os"
)

func CopyFile(pathFrom, pathTO string) {
	from, err := os.Open(pathFrom)
	if err != nil {
		log.Fatal(err)
	}
	defer from.Close()

	to, err := os.OpenFile(pathTO, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer to.Close()

	_, err = io.Copy(to, from)
	if err != nil {
		log.Fatal(err)
	}
}
