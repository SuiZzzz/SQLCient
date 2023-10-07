package window

import (
	"CenterTalk/data"
	"CenterTalk/ext"
	"CenterTalk/serlization"
	"CenterTalk/user"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/gorilla/websocket"
	"image/color"
)

type Config struct {
	Label  *canvas.Text
	List   *widget.List
	Conn   *websocket.Conn
	Window fyne.Window
	Layout *fyne.Container
}

var (
	module = &Config{}
)

func MakeWorkUI(window fyne.Window) *container.AppTabs {
	// 建立连接
	module.Window = window
	module.Conn = data.GetConn(window)
	// 中间选项卡
	center := getCenter()
	// 右边显示，默认打开通知
	right := getRight(window)
	split := container.NewWithoutLayout(center, right)
	// 2标签
	world := widget.NewLabel("222")

	// 左主选项卡
	tabs := container.NewAppTabs(
		container.NewTabItem("通知与群组", split),
		container.NewTabItem("AI", world),
	)
	tabs.SetTabLocation(container.TabLocationLeading)
	tabs.Resize(fyne.Size{Width: 100})
	return tabs
}

func getCenter() *fyne.Container {
	// background
	centerBackground := canvas.NewRasterWithPixels(func(x, y, w, h int) color.Color {
		return color.Black
	})
	centerBackground.Resize(fyne.Size{Width: 220, Height: 1000})

	// label
	funcLabel := canvas.NewText("  Function", color.NRGBA{
		R: 150,
		G: 211,
		B: 255,
		A: 255,
	})
	funcLabel.Alignment = fyne.TextAlignLeading
	funcLabel.TextSize = 20
	funcLabel.Move(fyne.Position{Y: 5})

	// 中选择
	notifyBtn := widget.NewButton("通知", getNotification())
	notifyBtn.Resize(fyne.Size{
		Width:  200,
		Height: 45,
	})
	notifyBtn.Move(fyne.Position{Y: 50, X: 10})
	groupBtn := widget.NewButton("群组成员", getGroup())
	groupBtn.Resize(fyne.Size{
		Width:  200,
		Height: 45,
	})
	groupBtn.Move(fyne.Position{Y: 100, X: 10})

	newBtn := widget.NewButton("+ New Request", getGroup())
	newBtn.Resize(fyne.Size{
		Width:  200,
		Height: 45,
	})
	newBtn.Move(fyne.Position{Y: 500, X: 10})

	center := container.NewWithoutLayout(centerBackground, funcLabel, notifyBtn, groupBtn, newBtn)
	return center
}

func getRight(window fyne.Window) *fyne.Container {
	// 背景
	raster := canvas.NewRasterWithPixels(
		func(x, y, w, h int) color.Color {
			return color.NRGBA{
				R: 244,
				G: 245,
				B: 238,
				A: 255,
			}
		})
	// 标题
	label := canvas.NewText("Notification", color.Black)
	label.TextSize = 26
	label.Alignment = fyne.TextAlignLeading

	// 分割线
	separator := widget.NewSeparator()
	separator.Resize(fyne.Size{
		Width:  1000,
		Height: 2,
	})
	separator.Move(fyne.Position{Y: 42})

	// 表格
	// websocket 发送消息
	data.SendReq(module.Window, module.Conn, serlization.BaseReq{
		UserId: user.CurrentUser.Id,
		Type:   1,
	})
	// 接收消息
	audit, message := data.ListenToNotifiedResp(module.Window, module.Conn)
	list := widget.NewList(
		func() int {
			return len(audit)
		},
		func() fyne.CanvasObject {
			tappableLabel := ext.TappableLabel{}
			tappableLabel.OnTapped = func() {
				dialog.ShowInformation(tappableLabel.Text, tappableLabel.Message, module.Window)
			}
			return &tappableLabel
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*ext.TappableLabel).Text = audit[i]
			o.(*ext.TappableLabel).Message = message[i]
			o.(*ext.TappableLabel).Refresh()
		})
	list.Move(fyne.Position{X: 50, Y: 50})
	list.Resize(fyne.Size{
		Width:  300,
		Height: 300,
	})

	//
	module.Label = label
	module.List = list

	// 布局
	right := container.NewBorder(nil, nil, nil, nil, raster)
	right.Resize(fyne.Size{
		Width:  1500,
		Height: 1000,
	})

	layout := container.NewWithoutLayout(right, label, separator, list)
	layout.Move(fyne.Position{X: 220})
	module.Layout = layout
	return layout
}

func getNotification() func() {
	return func() {
		label := module.Label
		label.Text = "Notification"
		label.Refresh()
		// websocket 发送消息
		data.SendReq(module.Window, module.Conn, serlization.BaseReq{
			UserId: user.CurrentUser.Id,
			Type:   1,
		})
		// 接收消息
		audit, message := data.ListenToNotifiedResp(module.Window, module.Conn)
		// 消息展示
		list := module.List
		// 删除原先列表，重新获取新列表
		module.Layout.Remove(list)
		list = widget.NewList(
			func() int {
				return len(audit)
			},
			func() fyne.CanvasObject {
				tappableLabel := ext.TappableLabel{}
				tappableLabel.OnTapped = func() {
					dialog.ShowInformation(tappableLabel.Text, tappableLabel.Message, module.Window)
				}
				return &tappableLabel
			},
			func(i widget.ListItemID, o fyne.CanvasObject) {
				o.(*ext.TappableLabel).Text = audit[i]
				o.(*ext.TappableLabel).Message = message[i]
				o.(*ext.TappableLabel).Refresh()
			})
		list.Move(fyne.Position{X: 50, Y: 50})
		list.Resize(fyne.Size{
			Width:  300,
			Height: 300,
		})
		module.List = list
		module.Layout.Add(list)
	}
}

func getGroup() func() {
	return func() {
		label := module.Label
		label.Text = "Group"
		label.Refresh()

		// websocket
		data.SendReq(module.Window, module.Conn, serlization.BaseReq{
			UserId: user.CurrentUser.Id,
			Type:   2,
		})
		// 接收消息
		nickname, message := data.ListenToGroupResp(module.Window, module.Conn)
		// 删除原先位置的列表
		// 删除原先列表，重新获取新列表
		module.Layout.Remove(module.List)
		// 用户列表
		list := widget.NewList(
			func() int {
				return len(nickname)
			},
			func() fyne.CanvasObject {
				tappableLabel := ext.TappableLabel{}
				tappableLabel.OnTapped = func() {
					dialog.ShowInformation(tappableLabel.Text, tappableLabel.Message, module.Window)
				}
				return &tappableLabel
			},
			func(i widget.ListItemID, o fyne.CanvasObject) {
				o.(*ext.TappableLabel).Text = nickname[i]
				o.(*ext.TappableLabel).Message = message[i]
				o.(*ext.TappableLabel).Refresh()
			})
		list.Move(fyne.Position{X: 50, Y: 50})
		list.Resize(fyne.Size{
			Width:  300,
			Height: 300,
		})
		module.List = list
		module.Layout.Add(list)
	}
}

func newRequest() func() {
	return func() {
		// 下拉框

	}
}
