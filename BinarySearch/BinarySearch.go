package BinarySearch

func IndexOf(n int,codeslice []int) (index int) {
	var lo int
	var hi  = len(codeslice)
	for lo < hi {
		mid := lo + (hi - lo) / 2
	
		if n > codeslice[mid] {
			lo = mid + 1
		} else if n < codeslice[mid] {
			hi = mid - 1
		} else {
			return mid
		}

	}	
	return -1
}