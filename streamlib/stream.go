package streamlib

import (
	"fmt"
)

type Stream interface {
	Start() error
}

type stream struct {
}

func (s *stream) Start() error {
	fmt.Println("Start")

	return nil
}
