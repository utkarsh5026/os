package schedule

type Process struct {
	Id             int
	ArrivalTime    int
	BurstTime      int
	TurnAroundTime int
	WaitingTime    int
}
