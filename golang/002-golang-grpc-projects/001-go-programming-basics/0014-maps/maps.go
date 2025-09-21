package main

import (
	"fmt"
)

func main() {

	// var mapVariable map[keyType]valueType
	// mapVariable := make(map[keyType]valueType)

	// Using a Map Literal
	// mapVariable = map[keyType]valueType{
	//	key1: value1,
	//	key2: value2,
	// }

	// Working with Maps
	myMap := make(map[string]int)
	fmt.Println(myMap)

	myMap["key1"] = 10
	myMap["code"] = 99
	fmt.Println(myMap)

	// In case key is not exists then returns a default value of the type
	fmt.Println(myMap["key2"])

	// Modify the value of map
	fmt.Println("Value before modify the map:", myMap["code"])
	myMap["code"] = 88
	fmt.Println("Value after modify the map:", myMap["code"])

	// Delete map
	delete(myMap, "key1")
	fmt.Println("Deleted the key1:", myMap)

	// Clear the map
	myMap["key1"] = 12
	myMap["key2"] = 16
	myMap["key3"] = 21
	clear(myMap)
	fmt.Println("Cleared the map:", myMap)

	// Map value check
	myMap["key1"] = 12
	myMap["key2"] = 16
	myMap["key3"] = 21
	if value, ok := myMap["key1"]; !ok {
		fmt.Println("value not found with key1.")
	} else {
		fmt.Println("key1 value is:", value)
	}

	// Create map without using the make function
	myMap1 := map[string]int{
		"key1": 22,
		"key2": 34,
		"key3": 56,
	}
	fmt.Println(myMap1)

	// Check equal for map
	myMap2 := map[string]int{
		"key1": 22,
		"key2": 34,
		"key3": 56,
	}

	// Range over map
	for key, value := range myMap2 {
		fmt.Printf("Map key: %v and value: %d", key, value)
	}

	// Zero value map
	var zeroValueMap map[string]string
	if zeroValueMap == nil {
		fmt.Println("The map is initialized to nil value")
	} else {
		fmt.Println("The map is not initialized to nil value")
	}

	// Zero value for map string
	val := zeroValueMap["key"]
	fmt.Println(val)

	// Add data to zero value map and we need to use make function
	zeroValueMap = make(map[string]string)
	zeroValueMap["key"] = "Value Added"
	fmt.Println("New key and value added to map:", zeroValueMap)

	// Length of map
	fmt.Println("Length of map for zeroValueMap:", len(zeroValueMap))

	// Nested map
	nestedMap := make(map[string]map[string]string)
	nestedMap["zeroValueMap"] = zeroValueMap
	fmt.Println("Nested map using zeroValueMap:", nestedMap)
}

/*
----------------------------------------Output---------------------------
map[]
map[code:99 key1:10]
0
Value before modify the map: 99
Value after modify the map: 88
Deleted the key1: map[code:88]
Cleared the map: map[]
key1 value is: 12
map[key1:22 key2:34 key3:56]
Map key: key1 and value: 22Map key: key2 and value: 34Map key: key3 and value: 56The map is initialized to nil value

New key and value added to map: map[key:Value Added]
Length of map for zeroValueMap: 1
Nested map using zeroValueMap: map[zeroValueMap:map[key:Value Added]]
----------------------------------------Output---------------------------
*/