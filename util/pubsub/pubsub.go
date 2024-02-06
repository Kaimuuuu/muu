package pubsub

import "sync"

type PubSub struct {
	lock     *sync.RWMutex
	listener []chan struct{}
}

func New() *PubSub {
	return &PubSub{
		lock:     new(sync.RWMutex),
		listener: make([]chan struct{}, 0),
	}
}

func (p *PubSub) Publish() {
	p.lock.RLock()
	defer p.lock.RUnlock()

	for _, c := range p.listener {
		c <- struct{}{}
	}

}

func (p *PubSub) Subscribe() (<-chan struct{}, func()) {
	p.lock.Lock()
	defer p.lock.Unlock()

	c := make(chan struct{}, 1)
	p.listener = append(p.listener, c)
	return c, func() {
		p.lock.Lock()
		defer p.lock.Unlock()

		for i, ch := range p.listener {
			if ch == c {
				p.listener = append(p.listener[:i], p.listener[i+1:]...)
				close(c)
				return
			}
		}
	}
}
