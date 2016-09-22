package job

import (
	"github.com/seantcanavan/config"
	"github.com/seantcanavan/routine"
	"github.com/seantcanavan/target"
	"path/filepath"
	"testing"
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
