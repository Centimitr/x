package ints

func maxOrMin(nums []int, max bool) int {
	if len(nums) == 0 {
		panic("input length zero")
	}
	most := nums[0]
	for i := 1; i < len(nums); i++ {
		if max && nums[i] > most {
			most = nums[i]
		} else if !max && nums[i] < most {
			most = nums[i]
		}
	}
	return most
}

func Max(nums []int) int {
	return maxOrMin(nums, true)
}

func Min(nums []int) int {
	return maxOrMin(nums, true)
}

func Sum(nums []int) int {
	var sum int
	for _, n := range nums {
		sum += n
	}
	return sum
}
