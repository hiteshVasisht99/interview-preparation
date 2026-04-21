package basicdsa

//gas := []int{1,2,3,4,5}
//cost := []int {3,4,5,1,2}

func CanCompleteCircuit(gas, cost []int) int {
	//Circuit complete is only possible in case totalGas >= totalCost
	totalGas, totalCost := 0, 0
	for i := 0; i < len(gas); i++ {
		totalGas += gas[i]
		totalCost += cost[i]
	}

	if totalGas < totalCost {
		return -1
	}

	currentGas := 0
	startIndex := 0

	for i := 0; i < len(gas); i++ {
		currentGas += gas[i] - cost[i]

		if currentGas < 0 {
			currentGas = 0
			startIndex = i + 1
		}

	}
	return startIndex
}
