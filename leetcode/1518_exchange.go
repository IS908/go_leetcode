package leetcode

func numWaterBottles(numBottles int, numExchange int) int {
	var (
		total      = numBottles
		rest       = numBottles % numExchange
		genBottles = numBottles / numExchange
	)
	total += genBottles
	numBottles = rest + genBottles
	for numBottles >= numExchange {
		rest = numBottles % numExchange
		genBottles = numBottles / numExchange
		total += genBottles
		numBottles = rest + genBottles
	}
	return total
}
