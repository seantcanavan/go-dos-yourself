package main

import (
	"github.com/seantcanavan/threads"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestCaseOne(t *testing.T) {
	mockControl := gomock.NewController(t)
	defer mockControl.Finish()
	item := threads.NewMockthread_actions(mockControl)
	item.EXPECT().Start()
	item.Start()


}