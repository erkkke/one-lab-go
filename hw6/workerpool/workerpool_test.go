package workerpool

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestExecute(t *testing.T) {
	tasks := make([]func() error, 0, 100)

	for i := 0; i < 100; i++ {
		if i % 5 == 0 {
			tasks = append(tasks, func() error {
				time.Sleep(200 * time.Millisecond)
				return nil
			})
		} else {
			tasks = append(tasks, func() error {
				time.Sleep(200 * time.Millisecond)
				fmt.Println("error was occurred while executing")
				return errors.New("error")
			})
		}
	}

	result1 := Execute(tasks, 5)
	expected1 := ErrorCountOfErrExceeded

	result2 := Execute(tasks, 200)

	assert.Equal(t, expected1, result1)
	assert.Equal(t, nil, result2)
}
