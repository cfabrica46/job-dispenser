package job

import "log"

var myJobs = make(map[int]myJob)

type MyFunc func() error

func init() {
	err := AddJobIntoMyJobs(1, "print", myPrint)
	if err != nil {
		log.Fatal(err)
	}
}

type myJob struct {
	id     int
	name   string
	myFunc MyFunc
}
