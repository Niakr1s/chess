package events_test

import (
	"testing"

	"github.com/niakr1s/chess/events"
	"github.com/stretchr/testify/assert"
)

func Test_Split(t *testing.T) {
	ch := make(chan events.Event)

	chatCh, chessCh := events.SplitChannel(ch)

	go func() {
		ch <- events.AuthUsernameEvent{}
		ch <- events.ChatMessageEvent{}
		ch <- events.ChessMoveEvent{}
		close(ch)
	}()

	assert.IsType(t, events.ChatMessageEvent{}, <-chatCh)
	assert.IsType(t, events.ChessMoveEvent{}, <-chessCh)

	for _, ch := range []<-chan events.Event{chatCh, chessCh} {
		_, ok := <-ch
		assert.False(t, ok)
	}
}
