package main

import (
	"github.com/dustin/go-broadcast"
)

type Message struct {
	UserID string
	RoomID string
	Text   string
}

type Listener struct {
	RoomID string
	Chan   chan any
}

type Manager struct {
	roomChannels map[string]broadcast.Broadcaster
	open         chan *Listener
	close        chan *Listener
	delete       chan string
	messages     chan *Message
}

func NewRoomManager() *Manager {
	manager := &Manager{
		roomChannels: make(map[string]broadcast.Broadcaster),
		open:         make(chan *Listener, 100),
		close:        make(chan *Listener, 100),
		delete:       make(chan string, 100),
		messages:     make(chan *Message, 100),
	}

	go manager.run()
	return manager
}

func (m *Manager) run() {
	for {
		select {
		case listener := <-m.open:
			m.register(listener)
		case listener := <-m.close:
			m.deregister(listener)
		case roomid := <-m.delete:
			m.deleteBroadcast(roomid)
		case message := <-m.messages:
			m.room(message.RoomID).Submit(message.UserID + ": " + message.Text)
		}
	}
}

func (m *Manager) register(listener *Listener) {
	m.room(listener.RoomID).Register(listener.Chan)
}

func (m *Manager) deregister(listener *Listener) {
	m.room(listener.RoomID).Unregister(listener.Chan)
	close(listener.Chan)
}

func (m *Manager) deleteBroadcast(roomid string) {
	b, ok := m.roomChannels[roomid]
	if ok {
		b.Close()
		delete(m.roomChannels, roomid)
	}
}

func (m *Manager) room(roomid string) broadcast.Broadcaster {
	b, ok := m.roomChannels[roomid]
	if !ok {
		b = broadcast.NewBroadcaster(10)
		m.roomChannels[roomid] = b
	}
	return b
}

func (m *Manager) OpenListener(roomid string) chan any {
	listener := make(chan any)
	m.open <- &Listener{
		RoomID: roomid,
		Chan:   listener,
	}
	return listener
}

func (m *Manager) CloseListener(roomid string, channel chan any) {
	m.close <- &Listener{
		RoomID: roomid,
		Chan:   channel,
	}
}

func (m *Manager) DeleteBroadcast(roomid string) {
	m.delete <- roomid
}

func (m *Manager) Submit(userid, roomid, text string) {
	msg := &Message{
		UserID: userid,
		RoomID: roomid,
		Text:   text,
	}
	m.messages <- msg
}
