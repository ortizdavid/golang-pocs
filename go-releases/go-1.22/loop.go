package main

func main() {

	myMapArray := []map[string]int {
		{"a":1},
		{"b":2},
		{"c":3},
		{"d":4},
		{"e":4},
	}
	/*for i := range 10 {
		println(i)
	}*/
	for k, v := range myMapArray {
		println(k, v["a"])
	}
}