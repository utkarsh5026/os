package schedule

import "sort"

// FCFS implements the First-Come, First-Served scheduling algorithm.
// It takes a slice of Process and modifies their TurnAroundTime and WaitingTime in-place.
func FCFS(processes []Process) {
	// Get the number of processes
	cnt := len(processes)

	// Sort the processes by their arrival time in ascending order
	sort.Slice(processes, func(i int, j int) bool {
		return processes[i].ArrivalTime < processes[j].ArrivalTime
	})

	// Initialize the completion time
	completionTime := 0

	// Iterate over the sorted processes
	for i := 0; i < cnt; i++ {
		p := processes[i]

		// If the process arrives after the current completion time, update the completion time
		if completionTime < p.ArrivalTime {
			completionTime = p.ArrivalTime
		}

		// Add the burst time of the process to the completion time
		completionTime += p.BurstTime

		// Calculate and set the turn around time for the process
		processes[i].TurnAroundTime = completionTime - p.ArrivalTime

		// Calculate and set the waiting time for the process
		processes[i].WaitingTime = p.TurnAroundTime - p.BurstTime
	}
}
