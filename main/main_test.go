package main

import (
	"fmt"
	"testing"
)

func TestTheSkyIsBlue(t *testing.T) {
	if 1 != 1 {
		fmt.Println("houston, we have a problem")
		t.Error()
	}
}
