package job

import (
	"fmt"
	"time"
)

func myDefault(jobDispenser *JobDispenser, jobID string) (err error) {
	defer delete(jobDispenser.jobsInProgress, jobID)
	for {
		select {
		case <-jobDispenser.jobsInProgress[jobID].Abort:
			return
		default:
			fmt.Printf("\rJOB %s\n", jobID)
			fmt.Print(">")
			time.Sleep(time.Second * 2)
		}
	}
}

func myWait(jobDispenser *JobDispenser, jobID string) (err error) {
	return
}

func myAbort(jobDispenser *JobDispenser, jobID string) (err error) {
	defer delete(jobDispenser.jobsInProgress, jobID)

	fmt.Println("Abort All Jobs")

	for i := range jobDispenser.jobsInProgress {
		if i != jobID {
			jobDispenser.jobsInProgress[i].Abort <- true
		}
	}
	return
}
