package job

import (
	"errors"
)

type JobDispenser struct {
	myJobs  []myJob
	restore chan bool
}

func NewJobDispenser() JobDispenser {
	return JobDispenser{myJobs: []myJob{}, restore: make(chan bool)}
}

func NewJobDispenserFilled() (jobDispenser JobDispenser, err error) {
	jobDispenser = NewJobDispenser()
	err = jobDispenser.AddJob(newMyJob("Print some text", myPrint))
	if err != nil {
		return
	}

	err = jobDispenser.AddJob(newMyJob("Wait to finish jobs", myAbort))
	if err != nil {
		return
	}

	err = jobDispenser.AddJob(newMyJob("Abort all jobs", myAbort))
	if err != nil {
		return
	}

	return
}

func (j *JobDispenser) AddJob(job myJob) (err error) {
	j.myJobs = append(j.myJobs, job)
	return
}

func (j JobDispenser) ExecuteJobByIndex(index int) (err error) {
	if !j.isValidIndex(index) {
		err = errors.New("index invalid")
		return
	}
	go j.myJobs[index-1].myFunc(j.restore)

	return
}

func (j JobDispenser) isValidIndex(index int) (check bool) {
	if len(j.myJobs) < index || index <= 0 {
		return
	}
	check = true
	return
}
