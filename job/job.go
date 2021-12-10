package job

type MyFunc func(j *JobDispenser, jobID string)

type myJob struct {
	name   string
	myFunc MyFunc
	Abort  chan bool
}

func newMyJob(name string, myFunc MyFunc) myJob {
	return myJob{name: name, myFunc: myFunc}
}
