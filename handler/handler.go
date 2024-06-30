package handler

import (
	"go-cloud-storage/meta"
	"go-cloud-storage/util"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"time"
)

// UploadHandler 处理文件上传
func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// 上传html页面
		data, err := os.ReadFile("./static/index.html")
		if err != nil {
			_, _ = io.WriteString(w, "internal server error")
			return
		}
		_, _ = w.Write(data)
	} else if r.Method == "POST" {
		// 接受文件数据流以及存储到本地
		file, header, err := r.FormFile("fileToUpload")
		if err != nil {
			log.Printf("Failed to get data from request, err:%s\n", err.Error())
			return
		}
		defer func(file multipart.File) {
			_ = file.Close()
		}(file)
		// 获取文件元信息
		fileMeta := meta.FileMeta{
			FileName: header.Filename,
			Location: "./tem/" + header.Filename,
			UploadAt: time.Now().Format("2006-01-02 15:04:05"),
		}
		// 创建本地文件
		newFile, err := os.Create(fileMeta.Location)
		if err != nil {
			log.Printf("Failed to create file, err:%s\n", err.Error())
			return
		}
		defer func(newFile *os.File) {
			_ = newFile.Close()
		}(newFile)
		// 拷贝文件
		fileMeta.FileSize, err = io.Copy(newFile, file)
		if err != nil {
			log.Printf("Failed to save file, err:%s\n", err.Error())
			return
		}

		_, _ = newFile.Seek(0, 0)
		fileMeta.FileSha1 = util.FileSha1(newFile)
		meta.UpdateFileMeta(fileMeta)
		http.Redirect(w, r, "/file/upload/success", http.StatusFound)
	}
}
func UploadSuccessHandler(w http.ResponseWriter, r *http.Request) {
	_, _ = io.WriteString(w, "upload success")
}
