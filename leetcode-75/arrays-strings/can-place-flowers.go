package main

func canPlaceFlowers(flowerbed []int, n int) bool {
	if n == 0 {
		return true
	}
	if n > len(flowerbed) {
		return false
	}

	possible := 0
	for i := 0; i < len(flowerbed)-1; i++ {
		if flowerbed[i] == 0 && flowerbed[i+1] == 0 {
			possible++
			i++
		}
		if flowerbed[i] == 1 {
			i++
		}
	}

	return possible >= n
}
