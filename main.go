package main

import (
	"CenterTalk/user"
	"CenterTalk/window"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/theme"
	"github.com/flopp/go-findfont"
	"os"
	"strings"
)

func init() {
	// 中文
	fontPaths := findfont.List()
	for _, path := range fontPaths {
		if strings.Contains(path, "STKAITI.TTF") || strings.Contains(path, "simhei.ttf") {
			os.Setenv("FYNE_FONT", path)
			break
		}
	}
}

func main() {
	ct := app.New()
	ct.Settings().SetTheme(theme.LightTheme())

	// 登录窗口
	loginW := ct.NewWindow("ct login")
	loginW.SetContent(window.MakeLoginUI(loginW))
	loginW.Resize(fyne.NewSize(500, 500))

	loginW.Show()
	go func() {
		currentUser := user.GetUser()
		for {
			if currentUser.Status == 1 {
				showW := ct.NewWindow("ct show")
				workUI := window.MakeWorkUI(showW)
				showW.SetContent(workUI)
				showW.Resize(fyne.NewSize(700, 700))
				showW.Show()
				loginW.Close()
				break
			}
		}
	}()
	ct.Run()
}

/*type config struct {
	editWidget    *widget.Entry
	previewWidget *widget.RichText
	currentURI    fyne.URI
	saveMenuItem  *fyne.MenuItem
}

var conf config

func main() {
	a := app.New()
	window := a.NewWindow("MarkDown")

	edit, pre := conf.makeUI()
	conf.makeMenu(window)

	window.SetContent(container.NewHSplit(edit, pre))
	window.Resize(fyne.Size{
		Width:  800,
		Height: 500,
	})

	window.CenterOnScreen()

	window.ShowAndRun()
}

// 页面
func (conf *config) makeUI() (*widget.Entry, *widget.RichText) {
	edit := widget.NewMultiLineEntry()
	preview := widget.NewRichTextFromMarkdown("...")
	conf.editWidget = edit
	conf.previewWidget = preview

	// 监听
	edit.OnChanged = preview.ParseMarkdown
	return edit, preview
}

// 菜单项
func (conf *config) makeMenu(window fyne.Window) {
	// 菜单项
	openMenuItem := fyne.NewMenuItem("Open...", conf.openFunc(window))
	saveMenuItem := fyne.NewMenuItem("Save...", conf.saveFunc(window))
	conf.saveMenuItem = saveMenuItem
	conf.saveMenuItem.Disabled = true
	saveAsMenuItem := fyne.NewMenuItem("Save as...", conf.saveAsFunc(window))

	// 菜单，包括几个菜单项
	fileMenu := fyne.NewMenu("File", openMenuItem, saveMenuItem, saveAsMenuItem)

	// 主菜单，显示
	menu := fyne.NewMainMenu(fileMenu)

	window.SetMainMenu(menu)
}

// 增加打开文件过滤器，过滤拓展名文件
var filter = storage.NewExtensionFileFilter([]string{".md", ".MD"})

// 打开文件
func (conf *config) openFunc(window fyne.Window) func() {
	return func() {
		openDialog := dialog.NewFileOpen(func(reader fyne.URIReadCloser, err error) {
			if err != nil {
				dialog.ShowError(err, window)
				return
			}
			if reader == nil {
				// cancel
				return
			}
			defer reader.Close()

			bytes, err := io.ReadAll(reader)
			if err != nil {
				dialog.ShowError(err, window)
				return
			}
			conf.currentURI = reader.URI()
			conf.editWidget.SetText(string(bytes))

			window.SetTitle(window.Title() + "-" + reader.URI().Name())

			conf.saveMenuItem.Disabled = false
		}, window)

		// 设置打开过滤
		openDialog.SetFilter(filter)
		openDialog.Show()
	}
}

// 保存
func (conf *config) saveFunc(window fyne.Window) func() {
	return func() {
		if conf.currentURI != nil {
			writer, err := storage.Writer(conf.currentURI)
			defer writer.Close()
			if err != nil {
				dialog.ShowError(err, window)
				return
			}
			_, err = writer.Write([]byte(conf.editWidget.Text))
			if err != nil {
				dialog.ShowError(err, window)
				return
			}
		}
	}
}

// 另存为
func (conf *config) saveAsFunc(window fyne.Window) func() {
	return func() {
		saveDialog := dialog.NewFileSave(func(write fyne.URIWriteCloser, err error) {
			if err != nil {
				dialog.ShowError(err, window)
				return
			}

			// 用户选择的存储位置
			if write == nil {
				// user cancel
				return
			}

			// 检查用户输入的文件名是否以md/MD结尾
			if !strings.HasSuffix(strings.ToLower(write.URI().Name()), ".md") {
				dialog.ShowInformation("ERROR", "Please name your file with .md extension", window)
				return
			}

			defer write.Close()
			_, _ = write.Write([]byte(conf.editWidget.Text))
			conf.currentURI = write.URI()

			// 设置window标题
			window.SetTitle(window.Title() + "-" + write.URI().Name())

			// 另存为后让save的菜单项不可点击
			conf.saveMenuItem.Disabled = false
		}, window)

		saveDialog.SetFileName("untitled.md")
		saveDialog.SetFilter(filter)
		// 展示
		saveDialog.Show()
	}
}*/

/*type App struct {
	output *widget.Label
}

var myApp App

func main() {
	a := app.New()
	window := a.NewWindow("MyNewApplication")

	output, entry, button := myApp.makeUI()
	window.SetContent(container.NewVBox(output, entry, button))
	window.Resize(fyne.Size{
		Width:  500,
		Height: 500,
	})
	window.ShowAndRun()
}

func (app *App) makeUI() (*widget.Label, *widget.Entry, *widget.Button) {
	output := widget.NewLabel("Hello World")
	entry := widget.NewEntry()
	btn := widget.NewButton("Enter", func() {
		app.output.SetText(entry.Text)
	})
	btn.Importance = widget.HighImportance
	app.output = output
	return output, entry, btn
}*/
