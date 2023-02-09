package main

import "fmt"

type Event interface {
	Register(EventListener)
}

type EventListener interface {
	OnFire()
}

type Mail struct {
	listener EventListener
}

func (m *Mail) Register(listener EventListener) {
	m.listener = listener
}

func (m *Mail) OnRecv() {
	m.listener.OnFire()
}

type Alarm struct {
}

func (a *Alarm) OnFire() {
	fmt.Println("메일이 왔습니다")
}

func main() {
	var mail = &Mail{}
	var listener EventListener

	listener = &Alarm{}

	mail.Register(listener)
	mail.OnRecv()
}

/*
PS C:\Users\조재모\goprojects\example\ex26.eventAlarm> .\ex26.eventAlarm.exe
메일이 왔습니다
*/
