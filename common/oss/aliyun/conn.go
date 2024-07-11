package aliyun

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"go-cloud-storage/pkg/logger"
	"io"
)

type Options struct {
	Endpoint        string
	AccessKeyId     string
	AccessKeySecret string
	BucketName      string
}

type OSS struct {
	options Options
}

// UploadFile 上传文件到OSS存储桶。
// objectKey 是上传到OSS的目标文件名。
// reader 是包含要上传的数据的读取器。
// 返回值是上传后的文件URL和可能的错误。
func (o *OSS) UploadFile(objectKey string, reader io.Reader) (string, error) {
	// 创建一个新的OSS存储桶实例。
	bucket, err := o.newBucket()
	if err != nil {
		// 如果创建存储桶实例失败，返回错误信息。
		return "", err
	}
	// 将数据从reader上传到指定的objectKey。
	// 上传文件的bytes。
	err = bucket.PutObject(objectKey, reader)
	if err != nil {
		// 如果上传失败，返回错误信息。
		return "", err
	}
	// 根据存储桶名称、Endpoint和objectKey生成文件的URL。
	url := o.CreateFileURL(o.options.BucketName, o.options.Endpoint, objectKey)
	// 返回上传后的文件URL和nil错误。
	return url, nil
}

func (o *OSS) UploadByteFile(objectKey string, fileBuf []byte) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (o *OSS) DeleteFile() error {
	//TODO implement me
	panic("implement me")
}

func (o *OSS) CreateFileURL(bucketName, endpoint, objectName string) string {
	//TODO implement me
	panic("implement me")
}

func (o *OSS) CreateObjectKey(suffix string, directories ...string) string {
	//TODO implement me
	panic("implement me")
}

func Init(opts Options) *OSS {
	return &OSS{options: opts}
}

func (o *OSS) newBucket() (*oss.Bucket, error) {
	// 创建OSSClient实例。
	client, err := oss.New(o.options.Endpoint, o.options.AccessKeyId, o.options.AccessKeySecret)
	if err != nil {
		logger.Errorw("创建OSSClient实例失败", "err", err)
		return nil, err
	}
	// 获取Bucket存储空间。
	bucket, err := client.Bucket(o.options.BucketName)
	if err != nil {
		logger.Errorw("获取Bucket存储空间失败", "err", err)
		return nil, err
	}
	return bucket, nil
}
