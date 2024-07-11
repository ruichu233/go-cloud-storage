package service

import (
	"go-cloud-storage/app/upload/cache"
	"go-cloud-storage/common/errno"
	"go-cloud-storage/common/response"
	"go-cloud-storage/pkg/config"
	"go-cloud-storage/pkg/logger"
	"math"
	"os"
	"strconv"
)
import "github.com/gin-gonic/gin"

// MultipartUploadInfo :初始化分块上传
type MultipartUploadInfo struct {
	FileHash   string `json:"filehash"` // 文件hash
	FileSize   int    `json:"filesize"` // 文件大小
	UploadID   string // 上传id
	ChunkSize  int    // 分块大小
	ChunkCount int    // 分块数量
}

func init() {
	os.MkdirAll("./tmp", 0744)
}

// InitialMultipartUploadHandler :初始化分块上传
func InitialMultipartUploadHandler(c *gin.Context) {
	// 1、解析请求参数
	var uploadInfo MultipartUploadInfo
	c.ShouldBind(&uploadInfo)
	username := c.Request.FormValue("username")
	filehash := c.Request.FormValue("filehash")
	filesize, err := strconv.Atoi(c.Request.FormValue("filesize"))
	if err != nil {
		response.WriteResponse(c, errno.ErrInvalidParameter, nil)
	}
	// 2、生成分块上传信息
	uploadInfo := MultipartUploadInfo{
		FileHash:   filehash,
		FileSize:   filesize,
		UploadID:   username + filehash,
		ChunkSize:  config.ChunkSize,
		ChunkCount: int(math.Ceil(float64(filesize) / float64(config.ChunkSize))),
	}

	mp := map[string]interface{}{
		"filehash":   uploadInfo.FileHash,
		"filesize":   uploadInfo.FileSize,
		"chunkcount": uploadInfo.ChunkCount,
	}

	if err := cache.NewUploadCache().HSetMultipartUploadInfo(c, uploadInfo.UploadID, mp); err != nil {
		logger.Errorw("set upload info error", "err", err)
		response.WriteResponse(c, err, nil)
	}

	response.WriteResponse(c, errno.OK, uploadInfo)
}
