package job

var myJobs = make(map[int]myJob)

type MyFunc func() error

func init() {
	addJobIntoMyJobs(1, "print", myPrint)
}

type myJob struct {
	id     int
	name   string
	myFunc MyFunc
}
