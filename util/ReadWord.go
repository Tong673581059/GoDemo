package util

import (
	"fmt"
	"github.com/unidoc/unioffice/document"
	"io/ioutil"
	"log"
	"strings"
)

type InterfaceBean struct {
	Title    string
	Url      string
	Param    []string
	Response string `json:"response"`
}

func NewInterfaceBean() InterfaceBean {
	return InterfaceBean{
		"",
		"",
		[]string{},
		"",
	}
}

func ReadWord() []InterfaceBean {
	doc, err := document.Open("testData.docx")
	if err != nil {
		log.Fatalf("error opening document: %s", err)
	}
	//doc.Paragraphs()得到包含文档所有的段落的切片
	var interfaceList []InterfaceBean
	interfaceBean := InterfaceBean{}
	for i, para := range doc.Paragraphs() {
		//run为每个段落相同格式的文字组成的片段
		for j, run := range para.Runs() {
			fmt.Println(i, j, run.Text())
			fmt.Println("bean=======", interfaceBean)
			if i%4 == 0 {
				if i != 0 {
					interfaceList = append(interfaceList, interfaceBean)
					interfaceBean.Title = ""
					interfaceBean.Url = ""
					interfaceBean.Param = []string{}
					interfaceBean.Response = ""
				}
				interfaceBean.Title = run.Text()
			} else if i%4 == 1 {
				interfaceBean.Url = run.Text()
			} else if i%4 == 2 {
				interfaceBean.Param = strings.Split(run.Text(), " ")
			} else if i%4 == 3 {
				interfaceBean.Response = interfaceBean.Response + run.Text()
			}
		}
	}
	interfaceList = append(interfaceList, interfaceBean)
	return interfaceList
}

func ReadTxt(path string) []InterfaceBean {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("Open file error!", err)
	}

	var interfaceList []InterfaceBean
	interfaceBean := InterfaceBean{}
	for i, lineTxt := range strings.Split(string(file), "\r\n") {
		fmt.Println(lineTxt)
		if !strings.Contains(lineTxt, "#") {
			if i%4 == 0 {
				interfaceBean.Title = lineTxt
			} else if i%4 == 1 {
				interfaceBean.Url = lineTxt
			} else if i%4 == 2 {
				interfaceBean.Param = strings.Split(lineTxt, " ")
			} else if i%4 == 3 {
				interfaceBean.Response = lineTxt
				interfaceList = append(interfaceList, interfaceBean)
				interfaceBean.Title = ""
				interfaceBean.Url = ""
				interfaceBean.Param = []string{}
				interfaceBean.Response = ""
			}
		}
	}
	return interfaceList
}

func GetBean(list []InterfaceBean, url string) InterfaceBean {
	for _, bean := range list {
		if bean.Url == url {
			return bean
		}
	}
	return InterfaceBean{}
}
