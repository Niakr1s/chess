package events

import (
	"strings"
)

// SplitChannel splits channel into separate channels, based on prefixes.
func SplitChannel(in <-chan Event) (chatCh <-chan Event, chessCh <-chan Event) {
	chatChannel := make(chan Event)
	chessChannel := make(chan Event)

	chatCh = chatChannel
	chessCh = chessChannel

	go func() {
		for e := range in {
			if strings.HasPrefix(e.Eventname(), "chat") {
				chatChannel <- e
			} else if strings.HasPrefix(e.Eventname(), "chess") {
				chessChannel <- e
			}
		}

		close(chatChannel)
		close(chessChannel)
	}()
	return
}
