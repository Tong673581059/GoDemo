package main

import (
	"./util"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lxn/walk"
	_ "net/http"
	"strings"
)

type MyMainWindow struct {
	*walk.MainWindow
	edit *walk.TextEdit
}

func main() {
	initInterface("hksports.txt")
	//mw := &MyMainWindow{}
	//err := MainWindow{
	//	AssignTo: &mw.MainWindow, //窗口重定向至mw，重定向后可由重定向变量控制控件
	//	Icon:     "test.ico",     //窗体图标
	//	Title:    "文件选择对话框",      //标题
	//	MinSize:  Size{Width: 150, Height: 200},
	//	Size:     Size{300, 400},
	//	Layout:   VBox{}, //样式，纵向
	//	Children: []Widget{ //控件组
	//		TextEdit{
	//			AssignTo: &mw.edit,
	//		},
	//		PushButton{
	//			Text:      "打开",
	//			OnClicked: mw.selectFile, //点击事件响应函数
	//		},
	//	},
	//}.Create() //创建
	//
	//if err != nil {
	//	fmt.Fprintln(os.Stderr, err)
	//	os.Exit(1)
	//}
	//
	//mw.Run() //运行
}

func (mw *MyMainWindow) selectFile() {

	dlg := new(walk.FileDialog)
	dlg.Title = "选择文件"
	dlg.Filter = "TXT文件 (*.txt)"

	mw.edit.SetText("") //通过重定向变量设置TextEdit的Text
	if ok, err := dlg.ShowOpen(mw); err != nil {
		mw.edit.AppendText("Error : File Open\r\n")
		return
	} else if !ok {
		mw.edit.AppendText("Cancel\r\n")
		return
	}

	s := fmt.Sprintf("Select : %s\r\n", dlg.FilePath)
	initInterface(dlg.FilePath)
	mw.edit.AppendText(s)
}

func initInterface(path string) {
	var interfaceList = util.ReadTxt(path)
	//fmt.Println(interfaceList)

	r := gin.Default()

	for _, bean := range interfaceList {
		r.POST(bean.Url, func(context *gin.Context) {
			currentBean := util.GetBean(interfaceList, context.Request.URL.Path)
			fmt.Println(currentBean)
			for _, param := range currentBean.Param {
				if !strings.Contains(param, "!") {
					var p = context.PostForm(param)
					if len(p) <= 0 {
						context.JSON(200, gin.H{
							"code":    10001,
							"message": "入参" + param + "为空",
						})
						return
					}
				}
			}
			context.JSON(200, util.StrToJson(currentBean.Response))
		})
	}
	r.Run()
}
