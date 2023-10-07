package window

import (
	"CenterTalk/user"
	"errors"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"image/color"
)

func MakeLoginUI(window fyne.Window) *fyne.Container {
	//var currentUser *user.User
	// 标题
	welcomeTest := canvas.NewText("Welcome to login", color.White)
	welcome := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), welcomeTest, layout.NewSpacer())

	// 账号
	usernameBind := binding.NewString()
	username := widget.NewEntryWithData(usernameBind)
	username.SetPlaceHolder("username")

	// 密码
	passwordBind := binding.NewString()
	password := widget.NewEntryWithData(passwordBind)
	password.Password = true
	password.SetPlaceHolder("password")

	//登录按钮
	loginButton := widget.NewButton("login", func() {
		u, _ := usernameBind.Get()
		p, _ := passwordBind.Get()
		currentUser := user.GetUser()
		currentUser.Status = 0
		currentUser.Username = u
		currentUser.Password = p
		resp := currentUser.Login(window)
		if resp == nil {
			currentUser.Status = 0
			return
		}
		if resp.Code == 0 {
			dialog.ShowInformation("succeed", "登录成功", window)
			currentUser.Status = 1
		} else {
			dialog.ShowInformation("error", "账号或密码错误", window)
			currentUser.Status = 0
		}
	})

	// 注册按钮
	register := widget.NewButton("register", func() {
		// 账号
		usernameBind := binding.NewString()
		username := widget.NewEntryWithData(usernameBind)

		// 密码
		passwordBind := binding.NewString()
		password := widget.NewEntryWithData(passwordBind)
		password.Password = true

		// 用户名
		nicknameBind := binding.NewString()
		nickname := widget.NewEntryWithData(nicknameBind)

		form := &widget.Form{
			Items: []*widget.FormItem{{Text: "账号", Widget: username}, {Text: "密码", Widget: password},
				{Text: "用户名", Widget: nickname}},
		}
		confirm := dialog.NewCustomConfirm("注册信息", "提交", "取消", form, func(b bool) {
			registerUser := user.GetUser()
			u, _ := usernameBind.Get()
			p, _ := passwordBind.Get()
			n, _ := nicknameBind.Get()
			register := registerUser.Register(window, u, p, n)
			if register.Code == 0 {
				dialog.ShowInformation("成功", "恭喜您！注册成功！", window)
			}
			switch register.Code {
			case 0:
				dialog.ShowInformation("成功", "恭喜您！注册成功！", window)
			case 1:
				dialog.ShowError(errors.New("账号已存在"), window)
			default:
				dialog.ShowError(errors.New("服务器异常，请稍后重试"), window)
			}
		}, window)
		confirm.Show()
	})

	box := container.NewVBox(welcome, username, password, loginButton, register)
	return box
}
