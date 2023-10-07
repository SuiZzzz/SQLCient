package user

import (
	"CenterTalk/serlization"
	"encoding/json"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"io"
	"log"
	"net/http"
	"strings"
	"sync"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Status   int
	Id       uint
	Nickname string
	Avatar   string
}

var once sync.Once
var user *User

var CurrentUser *User = &User{
	Id: 2,
}

func GetUser() *User {
	once.Do(func() {
		user = &User{}
	})
	return user
}

func (user *User) Login(window fyne.Window) *serlization.Response {
	req := &serlization.LoginReq{
		Username: user.Username,
		Password: user.Password,
	}
	bytes, _ := json.Marshal(req)
	payload := strings.NewReader(string(bytes))
	request, err := http.NewRequest(http.MethodPost, "http://localhost:8083/user/login", payload)
	if err != nil {
		log.Println(err)
		dialog.ShowInformation("ERROR", "Sorry, net has done", window)
		return nil
	}
	request.Header.Add("Content-Type", "application/json")
	res, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Println(err)
		dialog.ShowInformation("ERROR", "Sorry, net has done", window)
		return nil
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
		dialog.ShowInformation("ERROR", "Sorry, net has done", window)
		return nil
	}
	resp := &serlization.Response{}
	err = json.Unmarshal(body, resp)
	if err != nil {
		log.Println(err)
		dialog.ShowInformation("ERROR", "Sorry, net has done", window)
	}
	return resp
}

func (user *User) Register(window fyne.Window, username, password, nickname string) *serlization.Response {
	req := &serlization.LoginReq{
		Username: username,
		Password: password,
		Nickname: nickname,
	}
	bytes, _ := json.Marshal(req)
	payload := strings.NewReader(string(bytes))
	request, err := http.NewRequest(http.MethodPost, "http://localhost:8083/user/register", payload)
	if err != nil {
		log.Println(err)
		dialog.ShowInformation("ERROR", "Sorry, net has done", window)
		return nil
	}
	request.Header.Add("Content-Type", "application/json")
	res, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Println(err)
		dialog.ShowInformation("ERROR", "Sorry, net has done", window)
		return nil
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
		dialog.ShowInformation("ERROR", "Sorry, net has done", window)
		return nil
	}
	resp := &serlization.Response{}
	err = json.Unmarshal(body, resp)
	if err != nil {
		log.Println(err)
		dialog.ShowInformation("ERROR", "Sorry, net has done", window)
	}
	return resp
}
