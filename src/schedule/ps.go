package schedule

import "sort"

type PritorityProcess struct {
	Priority int
	Age      int
	Process
}

// sortProcesses sorts the processes by their arrival time in ascending order.
func sortProcesses(processes []PritorityProcess) {
	sort.Slice(processes, func(i int, j int) bool {
		// Compare arrival times of processes to determine their order
		return processes[i].ArrivalTime < processes[j].ArrivalTime
	})
}

// findMaxPriorityProcess finds the process with the maximum priority that has arrived and not been processed yet.
func findMaxPriorityProcess(processes []PritorityProcess, completionTime int) int {
	n := len(processes)
	maxPriorityIdx := -1

	for j := 0; j < n; j++ {
		p := processes[j]

		// Check if the process has arrived by the current completion time
		hasArrived := p.ArrivalTime <= completionTime
		// Check if the process has not yet been processed
		notYetProcessed := p.BurstTime > 0

		if hasArrived && notYetProcessed {
			maxPriorityNotYetFound := maxPriorityIdx == -1

			// Update maxPriorityIdx if a higher priority process is found
			if maxPriorityNotYetFound || (p.Priority+p.Age) > (processes[maxPriorityIdx].Priority+processes[maxPriorityIdx].Age) {
				maxPriorityIdx = j
			}
		}
	}

	return maxPriorityIdx
}

// processHighPriorityProcess processes the high priority process and updates its properties.
func processHighPriorityProcess(p *PritorityProcess, completionTime int) int {
	// Update completion time by adding the burst time of the current process
	completionTime += p.BurstTime

	p.TurnAroundTime = completionTime - p.ArrivalTime
	p.WaitingTime = p.TurnAroundTime - p.BurstTime
	p.BurstTime = 0
	p.Age = 0

	return completionTime
}

// PriorityScheduling implements the Priority Scheduling algorithm.
func PriorityScheduling(processes []PritorityProcess, ageIncrement int) {
	n := len(processes)

	// Sort processes based on arrival time
	sortProcesses(processes)

	completionTime := 0

	for i := 0; i < n; {
		// Find the process with the highest priority that has arrived
		maxPriorityIdx := findMaxPriorityProcess(processes, completionTime)

		// If no process has arrived yet, increment completion time
		if maxPriorityIdx == -1 {
			completionTime++
			continue
		}

		// Get the process with the highest priority
		highPriorityProcess := &processes[maxPriorityIdx]

		// Process the high priority process and update completion time
		completionTime = processHighPriorityProcess(highPriorityProcess, completionTime)

		// Update the process list
		processes[i] = *highPriorityProcess
		i++

		// Increment the age of all processes that have arrived and not yet completed
		incrementAge(processes, ageIncrement, completionTime)
	}
}

// incrementAge increments the age of all processes that have arrived and are not yet completed.
func incrementAge(processes []PritorityProcess, ageIncrement int, completionTime int) {
	n := len(processes)
	for i := 0; i < n; i++ {
		p := processes[i]

		// Increment age if the process has arrived and not yet completed
		if p.ArrivalTime <= completionTime && p.BurstTime > 0 {
			processes[i].Age += ageIncrement
		}
	}
}
