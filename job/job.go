package job

type MyFunc func(c chan bool) error

type myJob struct {
	name   string
	myFunc MyFunc
}

func newMyJob(name string, myFunc MyFunc) myJob {
	return myJob{name: name, myFunc: myFunc}
}
