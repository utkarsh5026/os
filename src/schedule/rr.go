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
	for i := range processes {
		remainingBurstTimes[i] = processes[i].BurstTime
	}

	time := 0        // The current time
	queue := []int{} // Queue to hold process indices

	// Keep processing until all processes are done
	for {
		// Load all processes that have arrived by the current time into the queue
		for i := 0; i < n; i++ {
			hasArrived := processes[i].ArrivalTime <= time
			notCompletelyProcessed := remainingBurstTimes[i] > 0

			if hasArrived && notCompletelyProcessed && !contains(queue, i) {
				queue = append(queue, i)
			}
		}

		// If the queue is empty, move time to the next process arrival time
		if len(queue) == 0 {
			nextArrivalTime := findNextArrivalTime(processes, remainingBurstTimes)

			if nextArrivalTime == -1 {
				break // All processes are done
			}

			time = nextArrivalTime
			continue
		}

		// Get the first process in the queue
		processSelectedIdx := queue[0]
		queue = queue[1:]

		// Process the selected process
		time = processInTimeQuantum(time, timeQuantum, remainingBurstTimes, completionTimes, processSelectedIdx)

		// Update the queue with newly arrived processes
		queue = processQueueUpdate(time, timeQuantum, processes, remainingBurstTimes, queue)

		// If the current process is not finished, put it back in the queue
		if remainingBurstTimes[processSelectedIdx] > 0 {
			queue = append(queue, processSelectedIdx)
		}
	}

	// Update the turnaround time and waiting time of each process
	updateProcessStats(processes, completionTimes)
}

// findNextArrivalTime returns the next arrival time of processes that have not completed yet.
func findNextArrivalTime(processes []RoundRobinProcess, remainingBurstTimes []int) int {
	nextArrivalTime := -1

	for i, p := range processes {
		if remainingBurstTimes[i] > 0 {
			noArrivalTimeFound := nextArrivalTime == -1

			if noArrivalTimeFound || p.ArrivalTime < nextArrivalTime {
				nextArrivalTime = p.ArrivalTime
			}
		}
	}
	return nextArrivalTime
}

// updateProcessStats updates the turnaround time and waiting time of each process.
// The turnaround time is the time from the arrival of the process to its completion.
// The waiting time is the turnaround time minus the burst time of the process.
func updateProcessStats(processes []RoundRobinProcess, completionTimes []int) {
	for i := range processes {
		p := &processes[i].Process
		p.TurnAroundTime = completionTimes[i] - p.ArrivalTime
		p.WaitingTime = p.TurnAroundTime - p.BurstTime
	}
}

// processInTimeQuantum processes the time quantum for a selected process.
// It takes the current time, the time quantum, the remaining burst times slice,
// the completion times slice, and the index of the selected process as input.
// If the remaining burst time of the selected process is greater than the time quantum,
// it subtracts the time quantum from the remaining burst time and adds it to the current time.
// If the remaining burst time of the selected process is less than or equal to the time quantum,
// it adds the remaining burst time to the current time, sets the completion time of the process,
// and sets the remaining burst time to 0.
func processInTimeQuantum(time int, timeQuantum int, remainingBurstTimes []int, completionTimes []int, processSelectedIdx int) int {
	if remainingBurstTimes[processSelectedIdx] > timeQuantum {
		time += timeQuantum
		remainingBurstTimes[processSelectedIdx] -= timeQuantum
	} else {
		time += remainingBurstTimes[processSelectedIdx]
		completionTimes[processSelectedIdx] = time
		remainingBurstTimes[processSelectedIdx] = 0
	}
	return time
}

// processQueueUpdate updates the process queue based on the arrival time, remaining burst time, and current queue state.
// It takes the current time, the time quantum, the processes slice, the remaining burst times slice, and the queue as input.
// It uses a map to optimize the check for whether a process is already in the queue.
// For each process, it checks if the process has arrived, has not yet been processed, and is not already in the queue.
// If all these conditions are met, it adds the process to the queue.
func processQueueUpdate(time int, timeQuantum int, processes []RoundRobinProcess, remainingBurstTimes []int, queue []int) []int {
	queueMap := make(map[int]bool)
	for _, i := range queue {
		queueMap[i] = true
	}

	for i := 0; i < len(processes); i++ {
		p := processes[i]
		hasArrived := p.ArrivalTime > time-timeQuantum && p.ArrivalTime <= time
		notYetProcessed := remainingBurstTimes[i] > 0
		notAlreadyInQueue := !queueMap[i]

		if hasArrived && notYetProcessed && notAlreadyInQueue {
			queue = append(queue, i)
			queueMap[i] = true
		}
	}

	return queue
}

// contains checks if a slice contains a specific element.
func contains(slice []int, element int) bool {
	for _, v := range slice {
		if v == element {
			return true
		}
	}
	return false
}
