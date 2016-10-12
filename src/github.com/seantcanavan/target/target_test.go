package target

import (
    "testing"
)

func TestIPredatorTargetFromFile(t *testing.T) {
    t := TargetFromFile("samples/ping_predator01.json")

    if t.Port != 80 {
        t.Error("t.Port did not unmarshal correctly")
    }

    if t.IPV6 != false {
        t.Error("t.IPV6 did not unmarshal correctly")
    }

    if t.IPV4Address != "194.132.32.32" {
        t.Error("t.IPV4Address did not unmarshal correctly")
    }

    if t.Name != "The public DNS server for IPreadtor.se" {
        t.Error("t.Name did not unmarshal correctly")
    }
}

func TestLocalHostTargetFromFile(t *testing.T) {
    t := TargetFromFile("samples/all_localhost01.json")

    if t.Name != "my 1337 machine" {
        t.Error("t.Name did not unmarshal correctly")
    }

    if t.URI != "127.0.0.1" {
        t.Error("t.URI did not unmarshal correctly")
    }

    if t.Port != 8080 {
        t.Error("t.Port did not unmarshal correctly")
    }

    if t.IPV6 != false {
        t.Error("t.IPV6 did not unmarshal correctly")
    }

    if t.PortRange != false {
        t.Error("t.PortRange did not unmarshal correctly")
    }

    if t.IPRange != false {
        t.Error("t.IPRange did not unmarshal correctly")
    }
}
