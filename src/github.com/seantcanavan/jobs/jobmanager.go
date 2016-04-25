package jobs

type job struct {
	name string
	id int
}

func NewJob(name string, id int) *job {
	j := new(job)
	j.name = name
	j.id = id
	return j
}

type jobmanager interface {

}