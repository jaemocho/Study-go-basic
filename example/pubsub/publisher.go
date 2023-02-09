package main

import "context"

type Publisher struct {
	ctx context.Context
	// channel 인데 (chan <- string) 을 받는 channel
	// chan <- string : writeOnly 채널(넣기만 가능)
	// <- chan string : readOnly 채널(받기만 가능 )
	subscribeCh chan chan<- string
	publishCh   chan string
	subscribers []chan<- string
}

func NewPublisher(ctx context.Context) *Publisher {
	return &Publisher{
		ctx:         ctx,
		subscribeCh: make(chan chan<- string),
		publishCh:   make(chan string),
		subscribers: make([]chan<- string, 0),
	}
}

func (p *Publisher) Subscribe(sub chan<- string) {
	p.subscribeCh <- sub
}

func (p *Publisher) Publish(msg string) {
	p.publishCh <- msg
}

func (p *Publisher) Update() {
	for {
		select {
		// p.subscribeCh에 data가 들어오면 뽑아서
		// 구독자로 추가
		case sub := <-p.subscribeCh:
			p.subscribers = append(p.subscribers, sub)
			// p.publishCh 에 data가 들어오면 뽑아서
			// 구독자들의 채널에 msg 삽입
		case msg := <-p.publishCh:
			for _, subscriber := range p.subscribers {
				subscriber <- msg
			}
		case <-p.ctx.Done():
			wg.Done()
			return
		}
	}
}
