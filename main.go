package main

import (
	"fmt"
	"github.com/helloworld/dll"
)

const dataFile = "data/data.json"

func init() {
	fmt.Println("init")

}
func main() {
	////fmt.Println(os.Getwd());
	//feeds, err := search.RetrieveFeed()
	//
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(feeds)

	//array.TestArray()

	//array.TestMap()

	dll.OpenDll()
}
