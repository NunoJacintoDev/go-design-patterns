package publish_subscriber

import (
	"fmt"
)

type Publisher interface {
	start()
	AddSubscriberCh() chan<- Subscriber
	RemoveSubscriberCh() chan<- Subscriber
	PublishingCh() chan<- interface{}
	Stop()
}

type publisher struct {
	subscribers []Subscriber
	addSubCh    chan Subscriber
	removeSubCh chan Subscriber
	in          chan interface{}
	stop        chan struct{}
}

func NewPublisher() Publisher {
	p := &publisher{
		addSubCh:    make(chan Subscriber),
		removeSubCh: make(chan Subscriber),
		in:          make(chan interface{}),
		stop:        make(chan struct{}),
	}
	p.start()
	return p
}

func (p *publisher) start() {
	go func() {
		for {
			select {
			case msg := <-p.in:
				fmt.Println("[Publish] Going to publish message:", msg)
				for _, sub := range p.subscribers {
					sub.Notify(msg)
				}
			case sub := <-p.addSubCh:
				fmt.Println("[Publish] Adding new subscriber", sub)
				p.subscribers = append(p.subscribers, sub)
			case sub := <-p.removeSubCh:
				fmt.Println("[Publish] Removing subscriber", sub)
				for i, candidate := range p.subscribers {
					if candidate == sub {
						p.subscribers = append(p.subscribers[:i], p.subscribers[i+1:]...)
						candidate.Close()
						break
					}
				}
			case <-p.stop:
				fmt.Println("[Publish] Closing all")
				for _, sub := range p.subscribers {
					sub.Close()
				}
				close(p.addSubCh)
				close(p.in)
				close(p.removeSubCh)
				return
			}
		}
	}()

}

func (p *publisher) AddSubscriberCh() chan<- Subscriber {
	return p.addSubCh
}

func (p *publisher) RemoveSubscriberCh() chan<- Subscriber {
	return p.removeSubCh
}

func (p *publisher) PublishingCh() chan<- interface{} {
	return p.in
}

func (p *publisher) Stop() {
	close(p.stop)
}
