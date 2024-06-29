package handler

import (
	"io"
	"log"
	"net/http"
	"os"
)

// 处理文件上传
func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// 上传html页面
		data, err := os.ReadFile("./static/index.html")
		if err != nil {
			io.WriteString(w, "internal server error")
			return
		}
		w.Write(data)
	} else if r.Method == "POST" {
		// 接受文件数据流以及存储到本地
		file, header, err := r.FormFile("fileToUpload")
		if err != nil {
			log.Printf("Failed to get data from request, err:%s\n", err.Error())
			return
		}
		defer file.Close()
		newFile, err := os.Create("./tem/" + header.Filename)
		if err != nil {
			log.Printf("Failed to create file, err:%s\n", err.Error())
			return
		}
		defer newFile.Close()
		_, err = io.Copy(newFile, file)
		if err != nil {
			log.Printf("Failed to save file, err:%s\n", err.Error())
			return
		}

		http.Redirect(w, r, "/file/upload/success", http.StatusFound)
	}
}
func UploadSuccessHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "upload success")
}
