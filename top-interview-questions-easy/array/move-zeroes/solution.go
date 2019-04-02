package main

func moveZeros(nums []int) {
	j := -1
	for i := 0; i < len(nums); i++ {
		if j != -1 {
			if nums[i] != 0 {
				nums[j] = nums[i]
				if i != j {
					nums[i] = 0
				}
				for ; nums[j] != 0 && i != j; j++ {
				}
			}
		} else {
			if nums[i] == 0 {
				j = i
			}
		}
	}
}

func moveZeros2(nums []int) {
	i := 0
	for j := 0; j < len(nums); j++ {
		if nums[j] != 0 {
			nums[i] = nums[j]
			i++
		}
		j++
	}
	for i < len(nums) {
		nums[i] = 0
		i++
	}
}
