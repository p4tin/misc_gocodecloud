package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const WORKERS = 3
const SECSPERMINUTE = 60

type StatusInfo struct {
	WorkerId          int
	WorkCompleted     int
	LastWorkTimestamp time.Time
}

var wg sync.WaitGroup
var workerChans [WORKERS]chan bool

type SentinelService struct {
	AvgJobsPerSec     int
	WorkForLastPeriod int
	Status            chan StatusInfo
	WorkerStatuses    map[int][]StatusInfo
}

func (ss SentinelService) On() {
	for {
		select {
		case stat := <-ss.Status:
			fmt.Printf("Received a status: %+v\n", stat)
			ss.WorkerStatuses[stat.WorkerId] = append(ss.WorkerStatuses[stat.WorkerId], stat)
			if len(ss.WorkerStatuses[stat.WorkerId]) >= SECSPERMINUTE {
				ss.WorkerStatuses[stat.WorkerId] = append(ss.WorkerStatuses[stat.WorkerId][1:])
			}
			ss.WorkForLastPeriod = stat.WorkCompleted
			ss.AvgJobsPerSec = 0
			for _, work := range ss.WorkerStatuses[stat.WorkerId] {
				ss.AvgJobsPerSec = ss.AvgJobsPerSec + work.WorkCompleted
			}
			ss.AvgJobsPerSec = ss.AvgJobsPerSec / 60
			fmt.Printf("Status for Worker %d: %+v\n", stat.WorkerId, ss)
		default:
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func Worker(id int, done chan bool, status chan StatusInfo) {
	defer wg.Done()
	numJobs := 0
	timer := time.NewTicker(1 * time.Second)
	defer timer.Stop()

	for {
		select {
		case <-done:
			fmt.Println("Go Routime", id, "exiting on demand by the monitor")
			return
		case <-timer.C:
			status <- StatusInfo{WorkerId: id, WorkCompleted: numJobs, LastWorkTimestamp: time.Now()}
			numJobs = 0
		default:
			//Real Work happens here
			r := rand.Intn(100)
			time.Sleep(time.Duration(r) * time.Millisecond)
			numJobs++
		}
	}
}

func main() {
	wg.Add(WORKERS)
	workerChans = [WORKERS]chan bool{
		make(chan bool),
		make(chan bool),
		make(chan bool),
	}
	ss := SentinelService{
		Status:         make(chan StatusInfo),
		WorkerStatuses: make(map[int][]StatusInfo),
	}
	go ss.On()

	for x := 0; x < WORKERS; x++ {
		go Worker(x, workerChans[x], ss.Status)
	}

	wg.Wait()
}
