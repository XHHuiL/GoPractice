package algorithm

func QuickSort(values []int) {
	quickSort(values, 0, len(values)-1)
}

func quickSort(values []int, left, right int) {
	if left >= right {
		return
	}
	pivot := values[left]
	i := left + 1
	j := right
	for {
		for i <= right {
			if values[i] > pivot {
				break
			}
			i++
		}

		for j >= left {
			if values[j] <= pivot {
				break
			}
			j--
		}
		if i >= j {
			break
		}
		values[i], values[j] = values[j], values[i]
	}
	p := i - 1
	values[left], values[p] = values[p], values[left]
	quickSort(values, left, p-1)
	quickSort(values, p+1, right)
}
