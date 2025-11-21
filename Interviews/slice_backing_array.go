
func main() {

	arr := [5]int{1, 2, 3, 4, 5}
	slice := arr[:3]
	// 0, 1

	slice = append(slice, 500, 600)
	// 0, 1, 500, 600

	slice[2] = 1000
	// 0, 1, 1000, 600

	fmt.Printf("arr: %v\n", arr)
	fmt.Printf("slice: %v\n", slice)
}

// ACTUAL OUTPUT
// arr: [1 2 1000 500 600]
// slice: [1 2 1000 500 600]
