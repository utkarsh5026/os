package schedule

import "sort"

// SJF implements the Shortest Job First scheduling algorithm.
// It takes a slice of Process and modifies their TurnAroundTime and WaitingTime in-place.
func SJF(processes []Process) {
	// Get the number of processes
	n := len(processes)

	// Sort the processes by their arrival time in ascending order
	sort.Slice(processes, func(i int, j int) bool {
		return processes[i].ArrivalTime < processes[j].ArrivalTime
	})

	// Initialize the completion time
	completionTime := 0

	// Iterate over the sorted processes
	for i := 0; i < n; {
		// Initialize the index of the process with the minimum burst time
		minBurstIdx := -1

		// Find the process with the minimum burst time that has arrived and not been processed yet
		for j := 0; j < n; j++ {
			arrivesAtTime := processes[j].ArrivalTime <= completionTime
			notAlreadyProcessed := processes[j].BurstTime > 0

			if arrivesAtTime && notAlreadyProcessed {
				minNotFound := minBurstIdx == -1

				if minNotFound || processes[j].BurstTime < processes[minBurstIdx].BurstTime {
					minBurstIdx = j
				}
			}
		}

		// If no process can be processed at this time, increment the completion time and continue
		if minBurstIdx == -1 {
			completionTime++
			continue
		}

		// Process the shortest job
		shortest := processes[minBurstIdx]

		completionTime += shortest.BurstTime

		shortest.TurnAroundTime = completionTime - shortest.ArrivalTime

		shortest.WaitingTime = shortest.TurnAroundTime - shortest.BurstTime

		// Set the burst time of the process to 0 to indicate it has been processed
		shortest.BurstTime = 0

		// Update the process in the slice
		processes[i] = shortest

		// Move to the next process
		i++
	}
}
