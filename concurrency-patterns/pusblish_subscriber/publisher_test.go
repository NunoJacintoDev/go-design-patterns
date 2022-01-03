package publish_subscriber

import (
	"errors"
	"fmt"
	"sync"
	"testing"
)

func TestPublisher(t *testing.T) {
	msg := "Hello"
	var wgFunc sync.WaitGroup
	var wgClose sync.WaitGroup

	// Create Publisher and start listening to channels
	p := NewPublisher()

	// Create Mock Subscriber
	sub := &mockSubscriber{
		notifyTestingFunc: func(msg interface{}) {
			defer wgFunc.Done()
			s, ok := msg.(string)
			if !ok {
				t.Fatal(errors.New("Could not assert result"))
			}
			fmt.Println("[Mock-Subscriber] Got message from publisher!", s)
			if s != msg {
				t.Fail()
			}
		},
		closeTestingFunc: func() {
			fmt.Println("[Mock-Subscriber] Closing")
			wgClose.Done()
		},
	}

	// Add Subscriber
	p.AddSubscriberCh() <- sub

	// Publish Message
	wgFunc.Add(1)
	p.PublishingCh() <- msg
	wgFunc.Wait()

	pubCon := p.(*publisher)
	if len(pubCon.subscribers) != 1 {
		t.Error("Unexpected number of subscribers")
	}

	wgClose.Add(1)
	p.RemoveSubscriberCh() <- sub
	wgClose.Wait()

	//Number of subscribers is restored to zero
	if len(pubCon.subscribers) != 0 {
		t.Error("Expected no subscribers")
	}

	// Stop Publisher
	p.Stop()
}

// Subscriber Mock
type mockSubscriber struct {
	notifyTestingFunc func(msg interface{})
	closeTestingFunc  func()
}

func (m *mockSubscriber) Close() {
	m.closeTestingFunc()
}
func (m *mockSubscriber) Notify(msg interface{}) error {
	m.notifyTestingFunc(msg)
	return nil
}
