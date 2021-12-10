package job

import (
	"fmt"
	"time"
)

func myDefault(jobDispenser *JobDispenser, jobID string) {
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

/* func myWait(jobDispenser *JobDispenser, jobID string) (err error) {
	return
} */

func myAbort(jobDispenser *JobDispenser, jobID string) {
	defer delete(jobDispenser.jobsInProgress, jobID)

	if len(jobDispenser.jobsInProgress) == 1 {
		fmt.Printf("\rNot Jobs to Abort\n")
		fmt.Print("> ")
		return
	}

	for i := range jobDispenser.jobsInProgress {
		if i != jobID {
			jobDispenser.jobsInProgress[i].Abort <- true
		}
	}
	fmt.Printf("\rAbort All Jobs\n")
	fmt.Print("> ")
}
