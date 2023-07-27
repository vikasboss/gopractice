package main

import (
	"encoding/json"
	"fmt"
)

type Man struct {
	Name   string `json:"name`
	Weight string `json:"weight`
}

func main() {
	json_strings :=
		`{
		"name":"Rocky",
		"weight" "90"
	}`

	rocky := new(Man)
	json.Unmarshal([]byte(json_strings), rocky)
	fmt.Println(rocky)
	spot := new(Man)
	spot.Name = "Spot"
	spot.Weight = "20"
	jsonStr, _ := json.Marshal(spot)
	fmt.Printf("%s\n", jsonStr)

}
