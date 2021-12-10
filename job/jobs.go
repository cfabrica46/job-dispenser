package job

import (
	"fmt"
	"time"
)

func myDefault(jobDispenser *JobDispenser, jobID string) {
	defer delete(jobDispenser.jobsInProgress, jobID)

	t := time.NewTicker(time.Second * 2)

	for {
		select {
		case <-jobDispenser.jobsInProgress[jobID].Abort:
			return
		case <-t.C:
			fmt.Printf("JOB %s\n", jobID)
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
