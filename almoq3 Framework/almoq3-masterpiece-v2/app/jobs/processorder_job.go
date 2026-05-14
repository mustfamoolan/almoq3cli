package jobs

import (
	"fmt"
)

type ProcessOrderJob struct {
	Payload string
}

func NewProcessOrderJob(payload string) *ProcessOrderJob {
	return &ProcessOrderJob{
		Payload: payload,
	}
}

func (j *ProcessOrderJob) Handle() error {
	// Job logic here
	fmt.Printf("Processing ProcessOrderJob with payload: %s\n", j.Payload)
	return nil
}
