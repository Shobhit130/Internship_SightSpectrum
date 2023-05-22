package main

import (
	"fmt"
	"net/url"
)

const myurl = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=sjfh938jnd"

func main() {
	fmt.Println("Hello")

	fmt.Println(myurl)

	//parsing the url
	result, _ := url.Parse(myurl)

	fmt.Println(result.Scheme) //https
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qparams := result.Query()
	fmt.Printf("The type of query params is %T\n",qparams)

	fmt.Println(qparams["coursename"])

	for _,val := range qparams{
		fmt.Println("Param is : ",val)
	}

	//constructing a URL
	partOfUrl := &url.URL{
		Scheme: "https",
		Host: "lco.dev",
		Path: "/learn",
		RawPath: "user=shobhit",
	}
	finalUrl := partOfUrl.String()
	fmt.Println(finalUrl)
}
