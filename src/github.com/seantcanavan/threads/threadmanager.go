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

func (t thread) start() {
	t.messages <- "start real"
}

func (t thread) stop() {
	t.messages <- "stop real"
}

func (t thread) change_job() {
	t.messages <- "change_job real"
}

func (t thread) print_details() {
	t.messages <- "print_details real"
}

type thread_actions interface {
	start()
	stop()
	change_job()
	print_details()
}
