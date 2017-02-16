# HTTP Post Request with Base64 Barcode Image in Golang
Send base64 image data to Dynamsoft Barcode web service by using HTTP POST request in Go.

## How to Use
1. Get the package:

	```
	go get github.com/dynamsoft-dbr/golang/web-service
	```

2. Import the package to Go project:

	```
	import "github.com/dynamsoft-dbr/golang/web-service"
	```

3. Create **main.go**:

	```Golang
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

	```

4. Build and run the console app:

	```
	go install
	<GOPATH>/bin/main <barcode image file>
	```
