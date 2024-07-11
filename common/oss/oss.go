package oss

import "io"

// OSS 对象存储接口
type OSS interface {
	UploadFile(objectKey string, reader io.Reader) (string, error)
	UploadByteFile(objectKey string, fileBuf []byte) (string, error)
	DeleteFile() error
	CreateFileURL(bucketName, endpoint, objectName string) string
	CreateObjectKey(suffix string, directories ...string) string
}
