package main

import (
	"encoding/json"
	"log"
	"os"
	"reflect"
)

func sum(i interface{}, skipRed bool) int64 {

	if mm, ok := i.(map[string]interface{}); ok {
		mapSum := int64(0)
		for _, value := range mm {
			if vstr, ok := value.(string); ok && vstr == "red" && skipRed {
				return 0
			}
			mapSum += sum(value, skipRed)
		}
		return mapSum
	}

	if arr, ok := i.([]interface{}); ok {
		arrSum := int64(0)
		for _, v := range arr {
			arrSum += sum(v, skipRed)
		}
		return arrSum
	}

	if num, ok := i.(int64); ok {
		return num
	}

	if num, ok := i.(float64); ok {
		return int64(num)
	}

	if _, ok := i.(string); ok {
		return 0
	}

	log.Printf("unsupported type: %v, kind: %v", i, reflect.TypeOf(i).Kind())

	return 0
}

func main() {
	f, err := os.Open("2015/Day12/input")
	if err != nil {
		panic(err)
	}

	body := []interface{}{}

	if err := json.NewDecoder(f).Decode(&body); err != nil {
		panic(err)
	}

	sump1 := int64(0)
	sump2 := int64(0)
	for _, e := range body {
		sump1 += sum(e, false)
		sump2 += sum(e, true)
	}
	// sum := int64(0)
	// for s.Scan() {
	// 	b := bytes.NewReader([]byte(s.Text()))
	// 	number, err := binary.ReadVarint(b)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	sum += number
	// }

	log.Printf("sum of all numbers (part1): %d", sump1)
	log.Printf("sum of all numbers, excluding red (part2): %d", sump2)
}
