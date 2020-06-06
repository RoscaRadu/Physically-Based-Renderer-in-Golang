package net

import (
	"../trace"
	"sync"
	"time"
)

const (
	pongTimeout = 60 * time.Second
	pongWait    = 20 * time.Second
	pingMessage = "a"
)

type RenderingClient struct {
	Conn                *websocket.Conn
	Results             chan *trace.Sampler
	IsRenderingInProgress bool
	CloseChan           chan bool
	OperationId         string
	PongChannel         chan bool
	WaitGroup           sync.WaitGroup
}

func (c *RenderingClient) Initialize() {
	c.WaitGroup = sync.WaitGroup{}

	c.Conn.SetPongHandler(
		func(message string) error {
			c.PongChannel <- true
			return nil
		},
	)

	go pingThread(c.Conn, c.PongChannel)

	c.WaitGroup.Add(2)

	go c.ReadHandler()
	go c.ResultSender()

	c.WaitGroup.Wait()
}