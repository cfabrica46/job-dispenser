package job

type myJob struct {
	ID   int
	Name string
	Func func(args []string) (err error)
}

func (j myJob) Execute(args []string) (err error) {
	err = j.Func(args)
	if err != nil {
		return
	}
	return
}

var MyJobs = []myJob{
	{ID: 1, Name: "print", Func: myPrint},
}
