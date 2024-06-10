package schedule

import "sort"

// SRF function implements the Shortest Remaining Time First (SRTF) scheduling algorithm
func SRF(processes []Process, timeQuantum int) {
	n := len(processes)
	remainingTimes := make([]int, n)
	completionTimes := make([]int, n)
	processed := 0

	for i := range processes {
		remainingTimes[i] = processes[i].BurstTime
	}

	sort.Slice(processes, func(i int, j int) bool {
		return processes[i].ArrivalTime < processes[j].ArrivalTime
	})

	time := 0
	for processed < n {
		nextArrivalTime, noProcessFound := findNextArrivalTimeForSrf(processes, time, remainingTimes)

		if noProcessFound {
			break
		}

		if time < nextArrivalTime {
			time = nextArrivalTime
		}

		shortestRemainingIdx := findShortestRemainingIndex(processes, time, remainingTimes)

		timeToAdd, processCompleted := process(shortestRemainingIdx, timeQuantum, remainingTimes)
		time += timeToAdd
		if processCompleted {
			processed++
			completionTimes[shortestRemainingIdx] = time
		}
	}

	finalizeStats(processes, completionTimes)
}

// finalizeStats function calculates the TurnAroundTime and WaitingTime for each process
func finalizeStats(processes []Process, completionTimes []int) {
	for i := range processes {
		p := &processes[i]
		p.TurnAroundTime = completionTimes[i] - p.ArrivalTime
		p.WaitingTime = p.TurnAroundTime - p.BurstTime
	}
}

// process function simulates the execution of a process
func process(shortestRemainingIdx int, timeQuantum int, remainingTimes []int) (int, bool) {
	timeToAdd := 0
	isProcessComplete := false

	timeRemaining := remainingTimes[shortestRemainingIdx]

	if timeRemaining > timeQuantum {
		timeToAdd = timeQuantum
		remainingTimes[shortestRemainingIdx] -= timeQuantum
	} else {
		timeToAdd = timeRemaining
		remainingTimes[shortestRemainingIdx] = 0
		isProcessComplete = true
	}

	return timeToAdd, isProcessComplete
}

// findShortestRemainingIndex function finds the index of the process with the shortest remaining time
func findShortestRemainingIndex(processes []Process, time int, remainingTimes []int) int {
	shortestRemainingIdx := -1
	n := len(processes)
	for i := 0; i < n; i++ {
		p := processes[i]
		hasArrived := p.ArrivalTime <= time
		notYetProcessed := remainingTimes[i] > 0

		if hasArrived && notYetProcessed {
			if shortestRemainingIdx == -1 || remainingTimes[i] < remainingTimes[shortestRemainingIdx] {
				shortestRemainingIdx = i
			}
		}
	}

	return shortestRemainingIdx
}

func findNextArrivalTimeForSrf(processes []Process, currentTime int, remainingTimes []int) (int, bool) {
	nextArrivalTime := -1
	for i, p := range processes {
		if p.ArrivalTime > currentTime && remainingTimes[i] > 0 {
			if nextArrivalTime == -1 || p.ArrivalTime < nextArrivalTime {
				nextArrivalTime = p.ArrivalTime
			}
		}
	}
	if nextArrivalTime == -1 {
		return 0, true
	}
	return nextArrivalTime, false
}
