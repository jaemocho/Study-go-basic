package main

import (
	"context"
	"fmt"
)

type Subscriber struct {
	ctx   context.Context
	name  string
	msgCh chan string
}

func NewSubscriber(name string, ctx context.Context) *Subscriber {
	return &Subscriber{
		ctx:   ctx,
		name:  name,
		msgCh: make(chan string),
	}
}

func (s *Subscriber) Subscribe(pub *Publisher) {
	// publisher 에 내 채널 등록
	pub.Subscribe(s.msgCh)
}

func (s *Subscriber) Update() {
	for {
		select {
		// 내 채널에 data가 들어오면 출력
		case msg := <-s.msgCh:
			fmt.Printf("%s got Message:%s\n", s.name, msg)
		case <-s.ctx.Done():
			wg.Done()
			return
		}
	}
}
