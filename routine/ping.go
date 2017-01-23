package routine

import (
	"fmt"
	"net"
	"time"

	"github.com/seantcanavan/target"
	"github.com/tatsushid/go-fastping"
)

type Ping struct {
	RunTarget target.Target
}

func (p Ping) SetTarget(target target.Target) {
	p.RunTarget = target
}

func (p Ping) Run() error {
	fp := fastping.NewPinger()
	var ip *net.IPAddr
	var err error

	if p.RunTarget.IPV6 {
		ip, err = net.ResolveIPAddr("ip6:icmp", p.RunTarget.URI)
	} else {
		ip, err = net.ResolveIPAddr("ip4:icmp", p.RunTarget.URI)
	}

	if err != nil {
		return err
	}

	fp.AddIPAddr(ip)

	fp.OnRecv = func(addr *net.IPAddr, rtt time.Duration) {
		fmt.Print("IPAdd:", addr.String(), "\t")
		fmt.Println("Delay:", rtt)
	}

	fp.OnIdle = func() {
		fmt.Println("ping finished executing")
	}

	err = fp.Run()
	return err
}
