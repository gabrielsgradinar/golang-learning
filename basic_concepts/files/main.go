package main

import (
	"io/ioutil"
	"log"
	"os"
)

func main() {
	file, err := os.OpenFile(
		"test.txt",
		os.O_WRONLY|os.O_TRUNC|os.O_CREATE,
		0644,
	)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	byteSlice := []byte("I learn Golang!")
	bytesWritten, err := file.Write(byteSlice)

	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Bytes Written? %d\n", bytesWritten)

	bs := []byte("Go Programming is top!")
	err = ioutil.WriteFile("opa.txt", bs, 0644)
	if err != nil {
		log.Fatal(err)
	}

}
