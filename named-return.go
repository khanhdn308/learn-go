package main

func Split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return // Naked return, should not be used, this will return x and y
}
