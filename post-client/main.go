package main

import (
	"bytes"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"os"
)

func main() {
	// // 単純なPOST
	// values := url.Values{
	// 	"test": {"value"},
	// }
	// resp, err := http.PostForm("http://localhost:18888", values)
	// if err != nil {
	// 	panic(err)
	// }
	// log.Println("Status:", resp.Status)

	// // 任意のボディを送信
	// file, err := os.Open("main.go")
	// if err != nil {
	// 	panic(err)
	// }
	// resp, err := http.Post("http://localhost:18888", "text/plain", file)
	// if err != nil {
	// 	panic(err)
	// }
	// log.Println("Status:", resp.Status)

	// // multipart/form-data形式でファイル送信
	// var buffer bytes.Buffer
	// writer := multipart.NewWriter(&buffer)
	// writer.WriteField("name", "Michael Jackson")

	// fileWrite, err := writer.CreateFormFile("thumbnail", "photo.jpg")
	// if err != nil {
	// 	panic(err)
	// }
	// readFile, err := os.Open("photo.jpg")
	// if err != nil {
	// 	panic(err)
	// }
	// defer readFile.Close()
	// io.Copy(fileWrite, readFile)
	// writer.Close()

	// resp, err := http.Post("http://localhost:18888", writer.FormDataContentType(), &buffer)
	// if err != nil {
	// 	panic(err)
	// }
	// log.Println("Status:", resp.Status)

	// multipart/form-data形式でファイル送信
	// かつ、送信するファイルに任意のMIMEタイプを設定
	var buffer bytes.Buffer
	writer := multipart.NewWriter(&buffer)
	writer.WriteField("name", "Michael Jackson")

	part := make(textproto.MIMEHeader)
	part.Set("Content-Type", "image/jpeg")
	part.Set("Content-Disposition", `form-data; name="thumbnail"; filename="photo.jpg"`)
	fileWriter, err := writer.CreatePart(part)
	if err != nil {
		panic(err)
	}
	readFile, err := os.Open("photo.jpg")
	if err != nil {
		panic(err)
	}
	io.Copy(fileWriter, readFile)

	writer.Close()

	resp, err := http.Post("http://localhost:18888", writer.FormDataContentType(), &buffer)
	if err != nil {
		panic(err)
	}
	log.Println("Status:", resp.Status)
}
