package main

import (
	"os"
	"fmt"
	"github.com/dynamsoft-dbr/golang/web-service"
)

func main()  {
	var filename string
	if len(os.Args) == 1 {
		fmt.Println("Please specify a file.")
		return
	}
	filename = os.Args[1]

	_, err := os.Stat(filename)

	if err != nil {
		fmt.Println(err)
		fmt.Println("Please specify a vailid file name.")
		return
	}

	channel := make(chan string)
	// read file to base64
	go web_service.File2Base64(filename, channel)
	// read barcode with Dynamsoft web service
	web_service.ReadBarcode(channel)
	fmt.Println("Done.")
}
