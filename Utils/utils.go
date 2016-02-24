package utils

import (
	"crypto/rand"
	"fmt"
	"io"
	"log"
	"os"
)

func getID() string {
	s := make([]byte, 16)
	rand.Read(s)
	return fmt.Sprint(s)

}

func Copy(dst, src string) {
	sFile, err := os.Open(src)
	if err != nil {
		log.Fatal(err)
	}
	defer sFile.Close()

	eFile, err := os.Create(dst)
	if err != nil {
		log.Fatal(err)
	}
	defer eFile.Close()

	_, err = io.Copy(eFile, sFile) // first var shows number of bytes
	if err != nil {
		log.Fatal(err)
	}

	err = eFile.Sync()
	if err != nil {
		log.Fatal(err)
	}

}
