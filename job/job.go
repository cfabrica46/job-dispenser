package job

type MyFunc func(c chan<- bool) error

type myJob struct {
	id     int
	name   string
	myFunc MyFunc
}
