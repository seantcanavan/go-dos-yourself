package routine

import (
	"github.com/seantcanavan/go-dos-yourself/target"
)

type SSH struct {
    RunTarget target.Target
}

func (s SSH) SetTarget(target target.Target) {
    s.RunTarget = target
}

func (s SSH) Run() error {
    return nil
}
