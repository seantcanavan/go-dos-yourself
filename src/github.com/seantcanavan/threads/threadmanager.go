package threads


type thread struct {
	messages chan string
	id int
	name string
}

func NewThread(messages chan string, id int, name string) *thread {
	t := new(thread)
	t.messages = messages
	t.id = id
	t.name = name
	return t
}

func (t thread) Start() bool {
	t.messages <- "start real"
	return true
}

func (t thread) Stop() bool {
	t.messages <- "stop real"
	return true
}

func (t thread) ChangeJob(jobName string) {
	t.messages <- "change_job real"
}

func (t thread) PrintDetails() {
	t.messages <- "print_details real"
}

type thread_actions interface {
	Start() bool
	Stop() bool
	ChangeJob(jobName string)
	PrintDetails()
}

