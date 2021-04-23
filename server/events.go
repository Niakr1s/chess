package server

import (
	"encoding/json"
	"fmt"
	"time"
)

type Event struct {
	Event string          `json:"event"`
	Data  json.RawMessage `json:"data"`
}

type ChatMessageEvent struct {
	Message  string    `json:"message"`
	Username string    `json:"username"`
	Time     time.Time `json:"time"`
}

type AuthUsernameEvent struct {
	Username string `json:"username"`
}

func (c *WsClient) jsonToEvent(bytes []byte) (interface{}, error) {
	e := Event{}
	if err := json.Unmarshal(bytes, &e); err != nil {
		return nil, err
	}

	switch e.Event {
	case "chat:message":
		res := ChatMessageEvent{}
		if err := json.Unmarshal(e.Data, &res); err != nil {
			return nil, err
		}
		res.Username = c.Username
		res.Time = time.Now()
		return res, nil

	case "auth:name":
		res := AuthUsernameEvent{}
		if err := json.Unmarshal(e.Data, &res); err != nil {
			return nil, err
		}
		return res, nil

	default:
		return nil, fmt.Errorf("not known event")
	}
}

func (c *WsClient) eventToJson(e interface{}) ([]byte, error) {
	data, err := json.Marshal(e)
	if err != nil {
		return nil, err
	}

	j := Event{Data: data}
	switch e.(type) {
	case ChatMessageEvent:
		j.Event = "chat:message"

	case AuthUsernameEvent:
		j.Event = "auth:name"

	default:
		return nil, fmt.Errorf("not known event")
	}
	return json.Marshal(j)
}
