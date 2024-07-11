package service

import (
	"github.com/gin-gonic/gin"
	"go-cloud-storage/app/model"
	"go-cloud-storage/common/errno"
	"go-cloud-storage/common/response"
	"go-cloud-storage/pkg/config"
	"go-cloud-storage/pkg/logger"
	"go-cloud-storage/util"
	"io"
	"os"
	"time"
)

// TODO:未考虑大文件上传

// UploadHandler 上传文件
func UploadHandler(c *gin.Context) {
	// 1.通过表单上传文件时，获取上传文件的头部信息
	header, err := c.FormFile("fileToUpload")
	if err != nil {
		logger.Errorw("")
		response.WriteResponse(c, errno.ErrInvalidParameter, nil)
	}
	// 2.构建文件元信息
	fileMeta := model.FileMeta{
		FileName: header.Filename,
		Location: config.TempPath + header.Filename,
		FileSize: header.Size,
		UploadAt: time.Now().Format("2006-01-02 15:04:05"),
	}
	// 3.保存到临时目录
	if err := c.SaveUploadedFile(header, fileMeta.Location); err != nil {
		response.WriteResponse(c, err, nil)
	}

	file, err := os.Open(fileMeta.FileName)
	if err != nil {
		response.WriteResponse(c, err, nil)
	}
	fileMeta.FileSha1 = util.FileSha1(file)
	// TODO: 存储到OSS

	response.WriteResponse(c, errno.OK, nil)

}
func UploadSuccessHandler(c *gin.Context) {
	_, _ = io.WriteString(c.Writer, "upload success")
}
