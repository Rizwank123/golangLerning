package main

import (
	"fmt"
	"sort"
)

func main() {
	myMap := make(map[string]int)
	myMap["one"] = 1
	myMap["two"] = 2
	myMap["three"] = 3
	fmt.Println("Map created using make: ", myMap)

	fmt.Println("-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-")
	usingLitral()
}
func usingLitral() {
	myMap := map[string]int{
		"four": 4,
		"five": 5,
		"six":  6,
	}
	fmt.Println("map created using litrel: ", myMap)
	fmt.Println("-=-=-=-=-=-=-=-=--0==-=-=-==-")
	usingVar()
}
func usingVar() {
	var myMap map[string]int
	myMap = make(map[string]int)
	myMap["seven"] = 7
	myMap["eight"] = 8
	myMap["nine"] = 9
	fmt.Println("Map created by var: ", myMap)
	fmt.Println("-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-")
	itrationMap(myMap)

	addValue(myMap)
}

func itrationMap(myMap map[string]int) {
	for key, val := range myMap {
		fmt.Printf("Key: %s, Value: %d\n", key, val)
	}

}
func addValue(myMap map[string]int) {
	myMap["ten"] = 10
	myMap["eleven"] = 11
	myMap["towel"] = 12
	itrationMap(myMap)
	keys := sortMapByValue(myMap)
	fmt.Println("Sorted map by value")
	fmt.Println("-=-=-=-=-=-")
	for _, key := range keys {
		fmt.Printf("key: %s, value: %d\n", key, myMap[key])
	}
}
func sortMapByValue(m map[string]int) []string {
	var keys []string
	for key := range m {
		keys = append(keys, key)
	}
	sort.Slice(keys, func(i, j int) bool {
		return m[keys[i]] < m[keys[j]]
	})
	return keys

}
