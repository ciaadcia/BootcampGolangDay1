package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println("NO 1")
	var words1 = [4]string{"code", "java", "cool", "java"}
	fmt.Println(upperCaseExcept(words1))

	var words2 = [4]string{"black", "pink", "venom", "venom"}
	fmt.Println(upperCaseExcept(words2))

	fmt.Println("NO 2")
	var nums1 = [10]int{2, 3, 4, 5, 6, 7, 8, 9, 1, 10}
	fmt.Println(findMinMax(nums1))

	fmt.Println("NO 3")
	var nums2 = [10]int{5, 3, 4, 2, 6, 7, 8, 9, 1, 10}
	fmt.Println(findMinRange(nums2, 0, 10))
	fmt.Println(findMinRange(nums2, 0, 7))

	var nums3 = [10]int{1, 22, 3, 4, 5, 10, 7, 8, 9, 49}
	fmt.Println(findMaxRange(nums3, 0, 10))
	fmt.Println(findMaxRange(nums3, 2, 7))

	fmt.Println("NO 4")
	var nums4 = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(oddEvenOrder(nums4))

	fmt.Println("NO 5")
	var nums5 = [5]int{12, 15, 1, 5, 20}
	fmt.Println(rotateRight(nums5, 1))
	fmt.Println(rotateRight(nums5, 2))
	fmt.Println(rotateRight(nums5, 3))
}

func upperCaseExcept(words [4]string) [4]string {
	var result [4]string
	last := words[3]

	for i := 0; i < 4; i++ {
		if words[i] != last {
			result[i] = strings.ToUpper(words[i])
		} else {
			result[i] = words[i]
		}
	}
	return result
}

func findMinMax(nums [10]int) [2]int {
	var result [2]int
	min := nums[0]
	max := nums[0]

	for i := 0; i < 10; i++ {
		if nums[i] < min {
			min = nums[i]
		}
		if nums[i] > max {
			max = nums[i]
		}
	}
	result[0] = min
	result[1] = max
	return result
}

func findMinRange(nums [10]int, start, end int) [2]int {
	var result [2]int
	min := nums[start]
	index := start

	for i := start; i < end; i++ {
		if nums[i] < min {
			min = nums[i]
			index = i
		}
	}
	result[0] = min
	result[1] = index
	return result
}

func findMaxRange(nums [10]int, start, end int) [2]int {
	var result [2]int
	max := nums[start]
	index := start

	for i := start; i < end; i++ {
		if nums[i] > max {
			max = nums[i]
			index = i
		}
	}
	result[0] = max
	result[1] = index
	return result
}

func oddEvenOrder(nums [10]int) [10]int {
	var result [10]int
	pos := 0

	for i := 0; i < 10; i++ {
		if nums[i]%2 == 0 {
			result[pos] = nums[i]
			pos++
		}
	}
	for i := 0; i < 10; i++ {
		if nums[i]%2 != 0 {
			result[pos] = nums[i]
			pos++
		}
	}
	return result
}

func rotateRight(nums [5]int, n int) [5]int {
	var result [5]int
	n = n % 5

	for i := 0; i < 5; i++ {
		newIndex := i + n
		if newIndex >= 5 {
			newIndex -= 5
		}
		result[newIndex] = nums[i]
	}
	return result
}
