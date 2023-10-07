package data

import (
	"CenterTalk/serlization"
	"encoding/json"
	"errors"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"github.com/gorilla/websocket"
)

func GetConn(window fyne.Window) *websocket.Conn {
	conn, _, err := websocket.DefaultDialer.Dial("ws://localhost:8083/ws", nil)
	if err != nil {
		dialog.ShowError(errors.New("websocket connected err"), window)
		return nil
	}
	return conn
}

func SendReq(window fyne.Window, conn *websocket.Conn, req serlization.BaseReq) {
	bytes, err := json.Marshal(req)
	if err != nil {
		dialog.ShowError(errors.New("json marshal err"), window)
		return
	}
	err = conn.WriteMessage(websocket.TextMessage, bytes)
	if err != nil {
		dialog.ShowError(errors.New("websocket write err"), window)
		return
	}
}

func ListenToNotifiedResp(window fyne.Window, conn *websocket.Conn) (audit, message []string) {
	resp := &serlization.NotificationResp{}
	_, bytes, err := conn.ReadMessage()
	if err != nil {
		dialog.ShowError(errors.New("websocket read err"), window)
		return nil, nil
	}
	err = json.Unmarshal(bytes, resp)
	if err != nil {
		dialog.ShowError(errors.New("json unmarshal err"), window)
		return nil, nil
	}
	return resp.Audit, resp.Message
}

func ListenToGroupResp(window fyne.Window, conn *websocket.Conn) (nickname, message []string) {
	resp := &serlization.GroupRes{}
	_, bytes, err := conn.ReadMessage()
	if err != nil {
		dialog.ShowError(errors.New("websocket read err"), window)
		return nil, nil
	}
	err = json.Unmarshal(bytes, resp)
	if err != nil {
		dialog.ShowError(errors.New("json unmarshal err"), window)
		return nil, nil
	}
	return resp.Nickname, resp.Message
}
