package serlization

import (
	"time"
)

const (
	Notification = 1 << iota
	Group
)

type BaseReq struct {
	UserId uint `json:"user_id"`
	Type   uint `json:"type"`
}

type BaseResp struct {
	Status uint `json:"status"`
	Id     uint `json:"id"`
}

type NotificationResp struct {
	Audit   []string `json:"audit"`
	Message []string `json:"message"`
	Code    uint     `json:"code"`
	Error   string   `json:"error"`
}

type Message struct {
	Sender      string    `json:"sender"`
	Receiver    string    `json:"receiver"`
	Type        byte      `json:"type"`
	Audit       bool      `json:"audit"`
	SQLResult   bool      `json:"sql_result"`
	Text        string    `json:"text"`
	Error       string    `json:"error"`
	CreatedAt   time.Time `json:"created_at"`
	CheckedTime time.Time `json:"checked_time"`
}

type GroupRes struct {
	Message  []string `json:"message,omitempty"`
	Nickname []string `json:"Nickname,omitempty"`
}
type Member struct {
	Nickname string `json:"nickname"`
	Level    string `json:"level"`
}

type User struct {
	Username string `json:"username"`
	Id       uint   `json:"id"`
	Level    uint   `json:"level"`
}

type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Nickname string `json:"nickname"`
}

type Response struct {
	Code    int    `json:"code"`
	Data    any    `json:"data"`
	Message string `json:"message"`
	Error   error  `json:"error"`
}

type DataResp struct {
	User  UserBo `json:"user"`
	Token string `json:"token"`
}

type UserBo struct {
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	Level    byte   `json:"level"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Avatar   string `json:"avatar"`
	Id       uint   `json:"id"`
}
