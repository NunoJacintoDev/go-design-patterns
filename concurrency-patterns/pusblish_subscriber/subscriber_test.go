package publish_subscriber

import (
	"fmt"
	"strings"
	"sync"
	"testing"
)

func TestWriter(t *testing.T) {
	sub := NewWriterSubscriber(0, nil)

	msg := "Hello"
	var wg sync.WaitGroup
	wg.Add(1)
	stdoutPrinter := sub.(*writerSubscriber)
	stdoutPrinter.Writer = &mockWriter{
		testingFunc: func(res string) {
			if !strings.Contains(res, msg) {
				t.Fatal(fmt.Errorf("Incorrect string: %s", res))
			}
			wg.Done()
		},
	}

	err := sub.Notify(msg)
	if err != nil {
		wg.Done()
		t.Fatal(err)
	}
	wg.Wait()
	sub.Close()
}

type mockWriter struct {
	testingFunc func(string)
}

func (m *mockWriter) Write(p []byte) (n int, err error) {
	m.testingFunc(string(p))
	return len(p), nil
}
