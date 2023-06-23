package main

import (
	"fmt"
	"tmp_session_client"
)

func main() {

	const id = "10.34.108.15"
	const auxKey = "123x456"
	const exampleData = "Example data."

	// Create session
	client := new(tmp_session_client.Client)
	userCode, error := client.CreateSession(id, auxKey, exampleData)
	if error != nil {
		fmt.Println("Error: ", error)
	} else {
		fmt.Println("UserCode: ", userCode)
	}

	// Get session first time
	data, error := client.GetSessionData(id, auxKey, userCode)
	if error != nil {
		fmt.Println("Error: ", error)
	} else {
		fmt.Println("Data: ", data)
	}

	// Retrive session
	retrive, error := client.RetriveSession(id, auxKey, userCode)
	if error != nil {
		fmt.Println("Error: ", error)
	} else {
		fmt.Println("Data: ", retrive)
	}

	// Get data second time
	_, error2 := client.GetSessionData(id, auxKey, userCode)
	if error2 != nil {
		fmt.Println("Error: ", error2)
	}
}
