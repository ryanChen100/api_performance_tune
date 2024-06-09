package apilib

import (
	"api/performance_tune/model"
	"api/performance_tune/pkg/config"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

var (
	gzipUrlPrefix string = "http://" + config.GetTestSetting().RouterIp + ":" + config.GetTestSetting().GzipPort + "/"
)

func GzipPostItem(routerName, id string) {
	var url = gzipUrlPrefix + routerName + "/items"
	b, marshalErr := json.Marshal(&model.Item{
		ID:   id,
		Name: "ryan" + id,
	})

	if marshalErr != nil {
		fmt.Println(marshalErr)
	}

	req, newRequestErr := http.NewRequest("POST", url, bytes.NewBuffer(b))
	if newRequestErr != nil {
		fmt.Println(newRequestErr)
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	_, doErr := client.Do(req)
	if doErr != nil {
		fmt.Println(doErr)
	}
	// resBody, err := ioutil.ReadAll(res.Body)
	// if err != nil {
	// 	fmt.Printf("client: could not read response body: %s\n", err)
	// 	os.Exit(1)
	// }
	// fmt.Printf("PostItem: response body: %s\n", resBody)
}

func GzipGetItem(routerName, id string) {
	var url = gzipUrlPrefix + routerName + "/items/" + id

	req, newRequestErr := http.NewRequest("GET", url, nil)
	if newRequestErr != nil {
		fmt.Println(newRequestErr)
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	_, doErr := client.Do(req)
	if doErr != nil {
		fmt.Println(doErr)
	}

	// resBody, err := ioutil.ReadAll(res.Body)
	// if err != nil {
	// 	fmt.Printf("client: could not read response body: %s\n", err)
	// 	os.Exit(1)
	// }
	// fmt.Printf("GetItem: response body: %s\n", resBody)
}

func GzipPutItem(routerName, id string) {
	var url = gzipUrlPrefix + routerName + "/items/" + id
	b, marshalErr := json.Marshal(&model.Item{
		ID:   id,
		Name: "ryanryan" + id,
	})

	if marshalErr != nil {
		fmt.Println(marshalErr)
	}

	req, newRequestErr := http.NewRequest("PUT", url, bytes.NewBuffer(b))
	if newRequestErr != nil {
		fmt.Println(newRequestErr)
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	_, doErr := client.Do(req)
	if doErr != nil {
		fmt.Println(doErr)
	}

	// resBody, err := ioutil.ReadAll(res.Body)
	// if err != nil {
	// 	fmt.Printf("client: could not read response body: %s\n", err)
	// 	os.Exit(1)
	// }
	// fmt.Printf("PutItem: response body: %s\n", resBody)
}

func GzipDeleteItem(routerName, id string) {
	var url = gzipUrlPrefix + routerName + "/items/" + id

	req, newRequestErr := http.NewRequest("DELETE", url, nil)
	if newRequestErr != nil {
		fmt.Println(newRequestErr)
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	_, doErr := client.Do(req)
	if doErr != nil {
		fmt.Println(doErr)
	}

	// resBody, err := ioutil.ReadAll(res.Body)
	// if err != nil {
	// 	fmt.Printf("client: could not read response body: %s\n", err)
	// 	os.Exit(1)
	// }
	// fmt.Printf("DeleteItem: response body: %s\n", resBody)
}
