package global

import (
	"go-cloud-storage/util"
)

const SECRETKEY = "fP1YSYw7IKHRd9YANjFLY2XVK5dbeQxakubfErxvszBb73wjcfwAGxs6T0HxyK7j"

var SnowFlake *util.SnowFlake
var Jwt *util.JWT

func init() {
	SnowFlake, _ = util.NewSnowFlake(0, 0)
	Jwt = util.NewJWT(SECRETKEY)
}
