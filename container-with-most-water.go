package main

type Container struct {
	left  int
	right int
	area  int
}

func maxArea(height []int) int {
	var maxContainer Container

	for i := 0; i < len(height); i++ {
		if height[i] < maxContainer.left {
			continue
		}
    var container Container
		for j := len(height) - 1; j > i; j-- {
      if height[j] > container.right {
        container = 
        if container.area > maxContainer.area
      }
		}

	}

	return maxContainer.area
}
