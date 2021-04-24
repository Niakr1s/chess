package events

import (
	"time"

	"github.com/niakr1s/chess/chess"
)

type Event interface {
	Eventname() string
}

type AuthUsernameEvent struct {
	Username string `json:"username"`
}

func (aue AuthUsernameEvent) Eventname() string {
	return "auth:username"
}

type ChatMessageEvent struct {
	Message  string    `json:"message"`
	Username string    `json:"username"`
	Time     time.Time `json:"time"`
}

func (cme ChatMessageEvent) Eventname() string {
	return "chat:message"
}

type ChatUserJoinEvent struct {
	Username string    `json:"username"`
	Time     time.Time `json:"time"`
}

func (e ChatUserJoinEvent) Eventname() string {
	return "chat:userJoin"
}

type ChatUserLeaveEvent struct {
	Username string    `json:"username"`
	Time     time.Time `json:"time"`
}

func (e ChatUserLeaveEvent) Eventname() string {
	return "chat:userLeave"
}

type ChessMoveEvent struct {
	From     string `json:"from"`
	To       string `json:"to"`
	Username string `json:"username"`
}

func (cme ChessMoveEvent) Eventname() string {
	return "chess:move"
}

type ChessNewTurnEvent struct {
	chess.Game
}

func (cse ChessNewTurnEvent) Eventname() string {
	return "chess:newTurn"
}

type ErrorEvent struct {
	Message string `json:"message"`
}

func (ee ErrorEvent) Eventname() string {
	return "error"
}
