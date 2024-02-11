package main

import (
	"context"
	"fmt"
	"time"
	"vinted-bidder/internal/manager"
)

func main() {

	m := manager.New()
	ctx := context.Background()
	m.AddProcess("test", &TestProcess{
		stop: make(chan struct{}),
	})
	m.AddProcess("test2", &TestProcess{
		stop: make(chan struct{}),
	})

	m.StartAll(ctx)

}

type TestProcess struct {
	stop chan struct{}
}

func (t *TestProcess) Start(ctx context.Context) error {
	go t.run()
	time.Sleep(5 * time.Second)
	t.Stop(ctx)
	return nil
}

func (t *TestProcess) run() error {
	for {
		select {
		case <-t.stop:
			return nil
		default:
			fmt.Println("Running")
			time.Sleep(1 * time.Second)
		}
	}
}

func (t *TestProcess) Stop(ctx context.Context) {
	close(t.stop)
}
