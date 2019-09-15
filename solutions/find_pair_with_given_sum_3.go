package solutions

import "fmt"

func GetPairUsingHash(arr []int, sum int)  {
	hash := make(map[int]int)

	for i := 0; i < len(arr); i++ {
		if _, exists := hash[sum-arr[i]]; exists {
			fmt.Printf("pair found at index: %d and %d", hash[sum-arr[i]], i)
			return
		}
		hash[arr[i]] = i
	}
	fmt.Println("Could not found result")
}
