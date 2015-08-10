package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func main() {

	values := map[string]string{}
	file, err := ioutil.ReadFile("./file.json")

	if err != nil {
		fmt.Println("error1:", err)
	}
	err2 := json.Unmarshal(file, &values)

	if err2 != nil {
		fmt.Println("error2:", err2)
	}
	fmt.Printf("Results: %v\n", values)
}
