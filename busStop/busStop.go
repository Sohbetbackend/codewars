package main

func Number(busStops [][2]int) int {
	sum := 0
	for i := 0; i < len(busStops); i++ {
		sum += busStops[i][0] - busStops[i][1]
	}
	return sum
}

func main() {
	Number([][2]int{{10, 0}, {3, 5}, {5, 8}})
}
