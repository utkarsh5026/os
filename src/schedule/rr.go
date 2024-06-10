package schedule

type RoundRobinProcess struct {
	OriginalBurstTime int
	Process
}

// RoundRobin implements the Round Robin scheduling algorithm.
// It takes a slice of RoundRobinProcess and a time quantum as input.
// Each process is given a slice of the CPU's time in a cyclic order.
// The time that each process has the CPU is equal to the time quantum.
// If a process does not finish its execution within the time quantum, it is preempted and put at the back of the queue.
func RoundRobin(processes []RoundRobinProcess, timeQuantum int) {
	n := len(processes)
	completionTimes := make([]int, n)     // Stores the completion times of each process
	remainingBurstTimes := make([]int, n) // Stores the remaining burst times of each process

	// Initialize the remaining burst times array with the burst times of each process
	initializeBurstTimes(remainingBurstTimes, processes)

	time := 0             // The current time
	allProcessed := false // Flag to check if all processes have been processed

	// Keep looping until all processes have been processed
	for !allProcessed {
		allProcessed = true

		// Loop through each process
		for i := 0; i < n; i++ {
			// If the process has not been completely executed
			if remainingBurstTimes[i] > 0 {
				allProcessed = false

				// If the remaining burst time is greater than the time quantum
				if remainingBurstTimes[i] > timeQuantum {
					time += timeQuantum
					remainingBurstTimes[i] -= timeQuantum
				} else {
					// If the remaining burst time is less than or equal to the time quantum
					time += remainingBurstTimes[i]
					completionTimes[i] = time
					remainingBurstTimes[i] = 0
				}
			}
		}

		// Update the turnaround time and waiting time of each process
		updateProcessesStats(processes, completionTimes)
	}
}

// initializeBurstTimes initializes the remaining burst times array with the burst times of each process.
func initializeBurstTimes(burstTimes []int, processes []RoundRobinProcess) {
	for i, process := range processes {
		burstTimes[i] = process.BurstTime
	}
}

// updateProcessesStats updates the turnaround time and waiting time of each process.
// The turnaround time is the time from the arrival of the process to its completion.
// The waiting time is the turnaround time minus the burst time of the process.
func updateProcessesStats(processes []RoundRobinProcess, completionTimes []int) {
	for i := range processes {
		p := processes[i]
		p.TurnAroundTime = completionTimes[i] - p.ArrivalTime
		p.WaitingTime = p.TurnAroundTime - p.BurstTime
	}
}
