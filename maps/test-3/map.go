package main
import "fmt"



func main()  {
	countries := make(map[string]string) 
	countries["US"] = "United States"
	countries["IND"] = "India"
	for country, capital := range countries {
		fmt.Println(country, "-->", capital)
	}
}
