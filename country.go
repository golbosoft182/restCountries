package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Country struct {
	Name string `json:"name"`
}

func main() {
	resp, err := http.Get("https://restcountries.com/v2/all")
	if err != nil {
		fmt.Println("Error retrieving countries:", err)
		return
	}
	defer resp.Body.Close()

	var countries []Country
	if err := json.NewDecoder(resp.Body).Decode(&countries); err != nil {
		fmt.Println("Error decoding countries:", err)
		return
	}

	fmt.Println("List of all countries:")
	for _, country := range countries {
		fmt.Println(country.Name)
	}
}
