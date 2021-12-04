package job

import (
	"errors"
	"fmt"
)

type JobDispenser struct {
	myJobs map[int]myJob
}

func NewJobDispenser() JobDispenser {
	return JobDispenser{myJobs: make(map[int]myJob)}
}

func NewJobDispenserFilled() (jobDispenser JobDispenser, err error) {
	jobDispenser = NewJobDispenser()
	err = jobDispenser.AddJob(1, "print", myPrint)
	if err != nil {
		return
	}
	return
}

func (j JobDispenser) AddJob(id int, name string, myFunc MyFunc) (err error) {
	if j.myJobs[id].myFunc != nil {
		err = errors.New("ID job is already ocupied ")
	}
	j.myJobs[id] = myJob{id: id, name: name, myFunc: myFunc}
	return
}

func (j JobDispenser) ExecuteJobByID(id int, c chan<- bool) (err error) {
	return j.myJobs[id].myFunc(c)
}

func (j JobDispenser) PrintAllJobs() (result string) {
	for i := range j.myJobs {
		result += fmt.Sprintf("%d. %s\n", j.myJobs[i].id, j.myJobs[i].name)
	}
	return
}

func (j JobDispenser) VerifyIsExistByID(id int) bool {
	return j.myJobs[id].myFunc != nil
}
