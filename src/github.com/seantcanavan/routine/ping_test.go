package routine

import (
	"github.com/seantcanavan/target"
	"path/filepath"
	"testing"
)

func TestPingPass(t *testing.T) {
	target, err := target.TargetFromFile(filepath.Join("..", "target", "samples", "ping_google01.json"))

	if err != nil {
		t.Error(err)
	}

	var r Routine
	r = Ping{RunTarget: target}

	err = r.Run()

	if err != nil {
		t.Error(err)
	}
}
