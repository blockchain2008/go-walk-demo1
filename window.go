package main

import (
	"archive/zip"
	"fmt"
	"github.com/lxn/walk"
	"github.com/lxn/walk/declarative"
	"io"
	"os"
)

type Window interface {
	ShowWindow() // 展示窗体界面
}

// ComWindow 创建压缩解压缩界面类
type ComWindow struct {
	Window
	*walk.MainWindow // 主窗体
}

// LabWindow 展示压缩，解压成功失败提示信息的界面类
type LabWindow struct {
	Window
}

// Show 创建界面类对象
func Show(WindowType string) {
	var Win Window
	switch WindowType {
	case "main_window":
		Win = &ComWindow{}
	case "lab_window":
		Win = &LabWindow{}
	default:
		fmt.Println("参数传递错误")

	}
	Win.ShowWindow()
}

var labText *walk.Label // 用来展示提示信息的Label.
var Text string         // 保存提示信息的。

// ShowWindow 首先实现ShowWindow方法，展示出空白的窗口
func (comWindow *ComWindow) ShowWindow() {
	var unzipEdit *walk.LineEdit       // 选择“解压文件”文件文本框
	var saveUnZipEdit *walk.LineEdit   // 解压后文件存放路径的文本框
	var zipEdit *walk.LineEdit         // 用来展示要压缩的文件路径的文本框
	var saveZipEdit *walk.LineEdit     // 用来展示压缩后的文件的存放路径的文本框
	var unzipBtn *walk.PushButton      // 选择解压文件按钮
	var saveUnZipBtn *walk.PushButton  // 创建用来选择解压为你教案存放路径的按钮
	var zipBtn *walk.PushButton        // 选择要压缩文件的按钮
	var saveZipBtn *walk.PushButton    // 选择压缩后文件存放路径的按钮就
	var startUnZipBtn *walk.PushButton // 开始解压按钮
	var startZipBtn *walk.PushButton   // 开始压缩按钮

	pathWindow := new(ComWindow)
	err := declarative.MainWindow{
		AssignTo: &pathWindow.MainWindow,                    // 关联主窗体，表明创建主窗体
		Title:    "文件压缩助手",                                  // 窗口的标题名称
		Size:     declarative.Size{Width: 480, Height: 230}, // 指定窗口的宽度与高度
		// Layout:   declarative.VBox{},                        // 必须指定这个，否则报错
		Layout: declarative.HBox{}, // 水平布局方式
		Children: []declarative.Widget{
			// 左边区域
			declarative.Composite{
				Layout: declarative.Grid{Columns: 2, Spacing: 10}, // 左边区域分为两列布局
				Children: []declarative.Widget{
					declarative.LineEdit{ // 表示文本框
						AssignTo:    &unzipEdit, // 将创建好的文本框与变量关联，后面可以根据该变量获取文本框中的值。
						Text:        "要解压的文件路径",
						ToolTipText: "请输入要解压的文件路径",
					},
					// 添加选择解压文件的按钮
					declarative.PushButton{
						AssignTo: &unzipBtn,
						Text:     "选择要解压的文件",
						OnClicked: func() {
							// 匿名函数
							// fmt.Println(unzipBtn.Text())
							// 弹出选择文件对话框
							filePath := pathWindow.OpenFileManager()
							// fmt.Println(filePath)
							unzipEdit.SetText(filePath) // 将放回的文件的路径赋值给文件框。
						},
					},
					// 创建展示解压后文件存放路径的文本框。
					declarative.LineEdit{
						AssignTo:    &saveUnZipEdit,
						Text:        "解压后的文件所在目录",
						ToolTipText: "请输入解压后的文件所在目录",
					},
					// 创建用来选择解压后文件存放路径的按钮。
					declarative.PushButton{
						AssignTo: &saveUnZipBtn,
						Text:     "选择要保存的目录",
						OnClicked: func() {
							folderPath := pathWindow.OpenDirManager()
							// fmt.Println(folderPath)
							saveUnZipEdit.SetText(folderPath)
						},
					},
					// 创建一个用来展示要压缩的文件路径的文本框
					declarative.LineEdit{
						AssignTo:    &zipEdit,
						Text:        "要压缩的文件路径",
						ToolTipText: "请输入要压缩的文件路径",
					},
					declarative.PushButton{
						AssignTo: &zipBtn,
						Text:     "选择要压缩的文件",
						OnClicked: func() {
							filePath := pathWindow.OpenFileManager()
							zipEdit.SetText(filePath)
						},
					},
					// 添加压缩好的文件存放路径展示的文本框。
					declarative.LineEdit{
						AssignTo:    &saveZipEdit,
						Text:        "压缩后的文件所在目录",
						ToolTipText: "请输入压缩后的文件所在目录",
					},
					// 创建一个选择压缩后文件路径的按钮
					declarative.PushButton{
						AssignTo: &saveZipBtn,
						Text:     "选择要保存的目录",
						OnClicked: func() {
							folderPath := pathWindow.OpenDirManager()
							saveZipEdit.SetText(folderPath)
						},
					},
					// 用于展示压缩和解压缩后相应的提示信息
					declarative.Label{
						AssignTo: &labText,
						Text:     "",
					},
				},
			},
			// 右边区域
			declarative.Composite{
				Layout: declarative.Grid{Rows: 2, Spacing: 40},
				Children: []declarative.Widget{
					declarative.PushButton{
						AssignTo: &startUnZipBtn,
						Text:     "开始解压",
						OnClicked: func() {
							// fmt.Println("开始解压")
							// 解压文件，传递了压缩文件的路径和解压后文件存放的路径
							pathWindow.StartToUnZip(unzipEdit.Text(), saveUnZipEdit.Text())
							Text = "文件解压成功"
							Show("lab_window")
							// ASCII
							// 汉字转码：go中汉字采用的是 UTF8 编码。
							// 而 Windows 系统中汉字采用的编码格式是GBK，不转码的话，会出现中文文件名乱码的问题。
						},
					},
					declarative.PushButton{
						AssignTo: &startZipBtn,
						Text:     "开始压缩",
						OnClicked: func() {
							// fmt.Println("开始压缩")
							pathWindow.StartToZip(zipEdit.Text(), saveZipEdit.Text())
							Text = "文件压缩成功"
							Show("lab_window")
						},
					},
				},
			},
		},
	}.Create() // 创建窗口
	if err != nil {
		fmt.Println(err)
	}
	// 窗口的展示，需要通过坐标来指定。
	pathWindow.SetX(650) // x坐标
	pathWindow.SetY(300) // y坐标
	pathWindow.Run()     // 运行窗口，才能将创建的窗口给用户展示出来
}

// OpenFileManager 打开文件选择对话框
func (comWindow *ComWindow) OpenFileManager() (filePath string) {
	// 1: 创建文件对话框对象
	dlg := new(walk.FileDialog)
	dlg.Title = "选择文件"
	dlg.Filter = "所有文件(*.*)|*.*|文本文档(*.txt)|*.txt"
	// 2: 打开文件对话框
	isOpen, err := dlg.ShowOpen(comWindow) // 如果单击对话框中的“打开”按钮，返回true，否则返回false
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(isOpen)
	// 3: 获取选择的文件
	filePath = dlg.FilePath // 获取选中的文件路径
	return filePath
}

// OpenDirManager 打开浏览文件夹的窗口
func (comWindow *ComWindow) OpenDirManager() (folderPath string) {
	// 1: 创建对话框对象
	dlg := new(walk.FileDialog)
	// 2: 打开窗口
	_, err := dlg.ShowBrowseFolder(comWindow) // 展示浏览文件夹的窗口
	if err != nil {
		fmt.Println(err)
	}
	// 3: 获取选中的路径，并且返回
	folderPath = dlg.FilePath // 获取选中的路径
	return folderPath
}

// StartToUnZip 实现文件解压操作
func (comWindow *ComWindow) StartToUnZip(filePath string, saveFilePath string) {
	// 1: 获取第一个文本框中，要解压的文件的路径，并且读取压缩文件中的内容。
	reader, err1 := zip.OpenReader(filePath)
	if err1 != nil {
		fmt.Println("StartToUnZip-err1------>", err1)
	}
	defer reader.Close()
	// 2: 循环遍历压缩包中的文件。
	for _, file := range reader.File {
		rc, err2 := file.Open() // 打开从压缩文件中获取到的文件或文件夹
		if err2 != nil {
			fmt.Println("StartToUnZip-err2------>", err2)
		}
		defer rc.Close()
		// 构建完整的文件夹或者是文件的存放位置（文件夹或者文件存放路径+文件夹或者文件的名称）
		// C:\Test\hh
		newName := saveFilePath + "\\" + file.Name
		newName, err := UTF8ToGBK(newName)
		if err != nil {
			fmt.Println(err)
		}
		// 判断是否为文件夹，IsDir()：如果是文件夹，该方法返回值为true，否则为false。
		if file.FileInfo().IsDir() {
			// 创建文件夹
			err3 := os.MkdirAll(newName, os.ModePerm)
			if err3 != nil {
				fmt.Println("StartToUnZip-err3------>", err3)
			}
		}
		// 判断是否为文件
		if !file.FileInfo().IsDir() {
			f, err4 := os.Create(newName)
			if err4 != nil {
				fmt.Println("StartToUnZip-err4------>", err4)
			}
			defer f.Close()
			// 读取压缩包中文件的内容，然后写入到新创建的文件中。
			// read write
			_, err5 := io.Copy(f, rc)
			if err5 != nil {
				fmt.Println("StartToUnZip-err5------>", err5)
			}
		}
	}
	// 3: 判断一下是否是文件夹，如果是文件夹，则创建。
	// 4: 如果读入出来的是文件，则创建文件。
}

// ShowWindow 将提示信息打印在label上
func (lab *LabWindow) ShowWindow() {
	labText.SetText(Text)
}

// StartToZip 实现文件压缩操作
func (comWindow *ComWindow) StartToZip(filePath string, saveFilePath string) {
	// 1: 获取第四个文本框中，然后创建压缩文件。
	d, err := os.Create(saveFilePath)
	if err != nil {
		fmt.Println(err)
	}
	defer d.Close()
	// 2: 获取第三个文本框中的值，打开该文件。
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	// 3: 将压缩的文件写入到压缩包中。
	// 3.1 要获取要压缩的文件的信息。
	stat, err := file.Stat()
	if err != nil {
		fmt.Println(err)
	}
	header, err := zip.FileInfoHeader(stat)
	if err != nil {
		fmt.Println(err)
	}
	// 3.2 将要压缩的文件写入到压缩包中。
	writer1 := zip.NewWriter(d) // 根据创建的压缩包，创建了一个Writer指针，通过该指针，可以对压缩包进行操作。
	defer writer1.Close()
	writer2, err := writer1.CreateHeader(header)
	if err != nil {
		fmt.Println(err)
	}
	io.Copy(writer2, file)
}
