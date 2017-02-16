package web_service

import (
	"io/ioutil"
	"fmt"
	"encoding/json"
	"log"
	"net/http"
	"bytes"
	"io"
	"encoding/base64"
)

func ReadBarcode(channel chan string) {
	base64data := <-channel
	url := "http://demo1.dynamsoft.com/dbr/webservice/BarcodeReaderService.svc/Read"

	// construct JSON data
	data := make(map[string]interface{})
	data["image"] = base64data
	data["barcodeFormat"] = 234882047
	data["maxNumPerPage"] = 1
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}

	// invoke Dynamsoft Barcode service
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode == 200 {
		result, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(result))

		// decode JSON
		const resultKey = "displayValue"
		dec := json.NewDecoder(bytes.NewReader(result))
		for {
			t, err := dec.Token()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatal(err)
			}
			tmp := fmt.Sprintf("%v", t)
			if tmp == resultKey {
				t, _ := dec.Token()
				tmp := fmt.Sprintf("%v", t)
				fmt.Println("Barcode result: ", tmp)
				break
			}
		}

	} else {
		fmt.Println("Fail to read barcode. Status: ", resp.Status)
	}
}

func File2Base64(filename string, channel chan string) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	}
	channel <- base64.StdEncoding.EncodeToString(data)
}

