package events

import (
	"encoding/json"
	"fmt"
	"time"
)

type eventJson struct {
	Eventname string          `json:"event"`
	Data      json.RawMessage `json:"data"`
}

func JsonToEvent(bytes []byte, username string) (Event, error) {
	e := eventJson{}
	if err := json.Unmarshal(bytes, &e); err != nil {
		return nil, err
	}

	switch e.Eventname {
	case "auth:username":
		res := AuthUsernameEvent{}
		err := json.Unmarshal(e.Data, &res)
		if err != nil {
			return nil, err
		}
		return res, nil

	case "chat:message":
		res := ChatMessageEvent{}
		err := json.Unmarshal(e.Data, &res)
		if err != nil {
			return nil, err
		}
		res.Time = time.Now()
		res.Username = username
		return res, nil

	case "chess:move":
		res := ChessMoveEvent{}
		err := json.Unmarshal(e.Data, &res)
		if err != nil {
			return nil, err
		}
		res.Username = username
		return res, nil
	}
	return nil, fmt.Errorf("not known event")
}

func EventToJson(e Event) ([]byte, error) {
	data, err := json.Marshal(e)
	if err != nil {
		return nil, err
	}

	j := eventJson{Eventname: e.Eventname(), Data: data}
	return json.Marshal(j)
}
