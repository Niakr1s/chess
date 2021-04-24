package events_test

import (
	"testing"

	"github.com/niakr1s/chess/events"
	"github.com/stretchr/testify/assert"
)

func Test_Json(t *testing.T) {
	t.Run("AuthUsernameEvent", func(t *testing.T) {
		j := []byte(`{"event":"auth:username","data":{"username":"username"}}`)
		e, err := events.JsonToEvent(j, "")
		assert.NoError(t, err)
		expectedE := events.AuthUsernameEvent{Username: "username"}
		assert.Equal(t, expectedE, e)

		encoded, err := events.EventToJson(e)
		assert.NoError(t, err)
		assert.Equal(t, j, encoded)
	})
	t.Run("AuthUsernameEvent", func(t *testing.T) {
		j := []byte(`{"event":"chat:message","data":{"message":"message"}}`)
		e, err := events.JsonToEvent(j, "username")
		assert.NoError(t, err)
		expectedE := events.ChatMessageEvent{Username: "username", Message: "message"}
		assert.NotNil(t, e.(events.ChatMessageEvent).Time)
		expectedE.Time = e.(events.ChatMessageEvent).Time
		assert.Equal(t, expectedE, e)

		_, err = events.EventToJson(e)
		assert.NoError(t, err)
	})
}
