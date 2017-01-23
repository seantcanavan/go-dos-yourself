package job

import (
	"path/filepath"
	"testing"

	"github.com/seantcanavan/go-dos-yourself/config"
	"github.com/seantcanavan/go-dos-yourself/routine"
	"github.com/seantcanavan/go-dos-yourself/target"
)

func TestJobPass(t *testing.T) {

	tar, err := target.TargetFromFile(filepath.Join("..", "target", "samples", "all_localhost01.json"))
	if err != nil {
		t.Error(err)
	}

	rou := routine.Ping{RunTarget: tar}
	conf := config.DefaultConfig()

	job := &Job{RunRoutine: rou, RunConfig: conf}

	job.Start()
}
