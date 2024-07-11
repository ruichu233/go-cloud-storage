package config

var RedisHost = "127.0.0.1"
var RedisPort = "6379"
var RedisPassword = "123456"
var RedisDB = 0

// OSS 配置

var OSSAccessKeyID = "yourAccessKeyId"
var OSSAccessKeySecret = "yourAccessKeySecret"
var OSSEndpoint = "yourBucketName"
var OSSBucketName = "yourBucketName"

// 临时文件配置
var TempPath = "./temp"

// 分块上传配置
var ChunkSize = 5 * 1024 * 1024

type Config struct {
	Server   *Server             `yaml:"server"`
	MySQL    *MySQL              `yaml:"mysql"`
	Redis    *Redis              `yaml:"redis"`
	Etcd     *Etcd               `yaml:"etcd"`
	Services map[string]*Service `yaml:"services"`
	Domain   map[string]*Domain  `yaml:"domain"`
}

type Server struct {
	Name      string `yaml:"name"`
	Addr      string `yaml:"addr"`
	Mode      string `yaml:"mode"`
	TimeOut   int64  `yaml:"timeout"`
	JwtSecret string `yaml:"jwt-secret"`
	Salt      string `yaml:"salt"`
	Version   string `yaml:"version"`
}

type MySQL struct {
	DriverName string `yaml:"driverName"`
	Host       string `yaml:"host"`
	Port       string `yaml:"port"`
	Database   string `yaml:"database"`
	UserName   string `yaml:"username"`
	Password   string `yaml:"password"`
	Charset    string `yaml:"charset"`
}

type Redis struct {
	UserName string `yaml:"userName"`
	Addr     string `yaml:"addr"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}

type Etcd struct {
	Addr string `yaml:"addr"`
}

type Service struct {
	Name        string   `yaml:"name"`
	LoadBalance bool     `yaml:"loadBalance"`
	Addr        []string `yaml:"addr"`
}

type Domain struct {
	Name string `yaml:"name"`
}
