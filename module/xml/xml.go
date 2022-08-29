package xml

import (
	"encoding/xml"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
)

const FilePath = "common/config/temp.xml"

type MainNode struct {
	Name string    `xml:"name"`
	Sub  []SubNode `xml:"sub"`
}

type SubNode struct {
	SubName string `xml:"sub-name"`
}

func init() {
	log.Info("module xml")
	setup()
}

func setup() {
	_, err := os.Stat(FilePath)
	if err != nil {
		log.Error("xml配置文件不存在", err)
		createXml()
	} else {
		openXml()
	}
}

func openXml() {
	f, err := os.Open(FilePath)
	if err != nil {
		log.Error("配置文件打开失败", FilePath, err)
		return
	}
	defer f.Close()
	// xml 配置文件不存在，将 XmlConfig 写入到 xml
	wrb, _ := xml.MarshalIndent(MainNode{}, "", "   ")
	wrb = append([]byte(xml.Header), wrb...)
	err = ioutil.WriteFile(FilePath, wrb, 0666)
	if err != nil {
		log.Error("xml 写入失败:%v", err)
	} else {
		log.Info("xml 写入成功")
	}
}

func createXml() {
	_, err := os.Create(FilePath)
	if err != nil {
		log.Error("配置文件创建失败", FilePath, err)
		return
	}
	log.Info("配置文件创建成功")
	openXml()
}
