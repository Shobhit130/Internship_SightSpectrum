package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`              //- means : we don't want this field to be reflected whoever is consuming our API
	Tags     []string `json:"tags,omitempty"` // don't show field with nil value
}

func main() {
	fmt.Println("Hello")
	encodeJSON()
	decodeJSON()
}

func encodeJSON() {
	sampleCourses := []course{
		{"ReactJS", 299, "xyz.com", "aksj12", []string{"web-dev", "js"}},
		{"MERN", 199, "xyz.com", "ksjd981", []string{"full-stack", "js"}},
		{"Angular", 499, "xyz.com", "isjs901", nil},
	}

	//package this data as JSON data (encoding the data into JSON)
	finalJSON, err := json.Marshal(sampleCourses)

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finalJSON)

	//more readable format
	finalJSON2, err := json.MarshalIndent(sampleCourses, "", "\t")

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finalJSON2)
}

func decodeJSON() {
	//data in JSON format which is to be decoded
	//decode the data from JSON to struct
	jsonData := []byte(`
		{
			"coursename": "ReactJS",
			"Price": 299,
			"website": "xyz.com",
			"tags": ["web-dev","js"]
		}
	`)

	//validating the JSON data
	var lcoCourse course

	checkValid := json.Valid(jsonData)

	if checkValid {
		fmt.Println("Valid JSON data")
		//decoding lcoCourse struct from JSON format
		err := json.Unmarshal(jsonData, &lcoCourse)

		if err != nil {
			fmt.Println(err)
		}

		//printing the details of the decoded data
		fmt.Println("Struct is: ", lcoCourse)
		fmt.Printf("%s's price is %d and it is on %s.\n", lcoCourse.Name,
			lcoCourse.Price, lcoCourse.Platform)
		fmt.Printf("%#v\n", lcoCourse)
	} else {
		fmt.Println("Invalid JSON data")
	}

	//in some cases we just want to add data to key value

	var myData map[string]interface{}

	json.Unmarshal(jsonData,&myData)
	fmt.Printf("%v\n", myData)

}
