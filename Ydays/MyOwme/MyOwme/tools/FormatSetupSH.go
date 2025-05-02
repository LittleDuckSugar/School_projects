package main

import (
	"io/ioutil"
)

func main() {

	var res string
	g, _ := ioutil.ReadFile("../Setup.sh")
	for _, i := range string(g) {
		if i != '\r' {
			res += string(i)
		}
	}
	ioutil.WriteFile("../Setup.sh", []byte(res), 777)

	// g, _ := ioutil.ReadFile("../Setup.sh")
	// for _, i := range string(g) {
	// 	if i == '\r' {
	// 		fmt.Print("omg")
	// 	}
	// }
}
