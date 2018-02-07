package main

import (
	"archive/zip"
	"bytes"
	"log"
	"os"
)

func main() {
	//创建一个buffer保存压缩内容
	buf := new(bytes.Buffer)
	//创建一个压缩文档
	w := zip.NewWriter(buf)

	var files = []struct {
		Name, Body string
	}{
		{"aa.txt", "测试测试"},
		{"bb.txt", "wowowo"},
	}
	for _, file := range files {
		f, err := w.Create(file.Name)
		if err != nil {
			log.Fatal(err)
		}
		_, err = f.Write([]byte(file.Body))
		if err != nil {
			log.Fatal(err)
		}
	}

	// 关闭压缩文档
	err := w.Close()
	if err != nil {
		log.Fatal(err)
	}

	// 将压缩文档内容写入文件
	f, err := os.OpenFile("file.zip", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	buf.WriteTo(f)
}
