package job

import (
	"errors"

	"github.com/google/uuid"
)

type JobDispenser struct {
	myJobs         []myJob
	jobsInProgress map[string]myJob
}

func NewJobDispenser() JobDispenser {
	return JobDispenser{myJobs: []myJob{}, jobsInProgress: make(map[string]myJob)}
}

func NewJobDispenserFilled() (jobDispenser JobDispenser, err error) {
	jobDispenser = NewJobDispenser()
	jobDispenser.AddJob(newMyJob("Job Default", myDefault))
	jobDispenser.AddJob(newMyJob("Abort all jobs", myAbort))
	return
}

func (j *JobDispenser) AddJob(job myJob) {
	j.myJobs = append(j.myJobs, job)
}

func (j *JobDispenser) ExecuteJobByIndex(index int) (err error) {
	if !j.isValidIndex(index) {
		err = errors.New("index invalid")
		return
	}

	jobID := uuid.NewString()

	if j.jobsInProgress[jobID].myFunc != nil {
		err = errors.New("error with asign id")
		return
	}

	newJob := j.myJobs[index]
	newJob.Abort = make(chan bool)
	j.jobsInProgress[jobID] = newJob

	go j.jobsInProgress[jobID].myFunc(j, jobID)
	return
}

func (j JobDispenser) isValidIndex(index int) (check bool) {
	if index >= len(j.myJobs) || index < 0 {
		return
	}
	check = true
	return
}
