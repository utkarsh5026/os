package schedule

// Process represents a single process in a scheduling system.
type Process struct {
	ArrivalTime    int
	BurstTime      int
	CompletionTime int
	TurnAroundTime int
	WaitingTime    int
}

// FindNextArrivalTime iterates over a slice of processes and returns the smallest arrival time.
// If no process is found (i.e., the slice is empty), it returns -1 and true.
// If at least one process is found, it returns the smallest arrival time and false.
func FindNextArrivalTime(processes []Process) (int, bool) {
	nextArrivalTime := -1

	for i := range processes {
		p := processes[i]
		noProcessFound := nextArrivalTime == -1

		if noProcessFound || p.ArrivalTime < nextArrivalTime {
			nextArrivalTime = p.ArrivalTime
		}
	}

	// Return the next arrival time and a boolean indicating whether no process was found
	return nextArrivalTime, nextArrivalTime == -1
}
