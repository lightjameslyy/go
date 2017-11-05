package light

import "fmt"

// TestTypes : array, slice, map
func TestTypes() {

	/*********************************
	 * array
	**********************************/
	var arr1 [5]int
	fmt.Println(arr1)

	arr2 := [5]int{10, 20, 30, 40, 50}
	fmt.Println(arr2)

	arr3 := [...]int{10, 20, 30, 40, 50}
	fmt.Println(arr3)

	arr4 := [5]int{1: 10, 2: 20}
	fmt.Println(arr4)

	var arr5 [4][2]int
	fmt.Println(arr5)

	arr6 := [4][2]int{{10, 11}, {20, 21}, {30, 31}, {40, 41}}
	fmt.Println(arr6)

	/*********************************
	 * slice
	**********************************/
	// size = 5, capacity = 5
	slice1 := make([]string, 5)
	fmt.Println(slice1)

	// size = 3, capacity = 5
	slice2 := make([]int, 3, 5)
	fmt.Println(slice2)

	// similar to array, without specify length in []
	slice3 := []string{"red", "blue", "green", "yellow", "pink"}
	fmt.Println(slice3)

	// nil slice
	var slice4 []int
	fmt.Println(len(slice4))

	// empty slice
	slice5 := make([]int, 0)
	slice6 := []int{}
	fmt.Println(len(slice5))
	fmt.Println(len(slice6))

	// use the third index to specify capacity
	// slice[i:j:k] : size = j-i, capacity = k-i
	slice7 := slice3[2:3:3]

	// reallocate memory for slice7
	slice7 = append(slice7, "brown")
	fmt.Println(slice7)

	for index, value := range slice7 {
		fmt.Printf("index: %d, value: %s\n", index, value)
	}

	/*********************************
	 * map
	**********************************/
	// initialize a map using make
	dict1 := make(map[string]int)
	dict1["a"] = 'a'
	fmt.Println(dict1)

	// initialize a map using map literal
	dict2 := map[string]string{"red": "#da1337", "orange": "#e95a22"}
	fmt.Println(dict2)

	// nil map
	var colors map[string]string
	fmt.Println(colors)

	fmt.Println("hello Go!")
}
