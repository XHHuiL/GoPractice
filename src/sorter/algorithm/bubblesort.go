package algorithm

func BubbleSort(values []int) {
	flag := true
	for i := 1; i < len(values); i++ {
		flag = false
		for j := len(values) - 1; j >= i; j-- {
			if values[j] < values[j-1] {
				values[j], values[j-1] = values[j-1], values[j]
				flag = true
			}
		}
		if !flag {
			break
		}
	}
}
