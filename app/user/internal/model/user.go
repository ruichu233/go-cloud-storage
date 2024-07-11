package model

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

type User struct {
	BaseModel
	Username string `json:"username" gorm:"uniqueIndex:username;size:40;not null"`
	Password string `json:"password" gorm:"not null"`
}

func (*User) TableName() string {
	return "user"
}

func CreateUserKey(userId uint) string {
	userStr := strconv.FormatInt(int64(userId), 10)
	return "user_info::" + userStr
}

func CreateMapUserInfo(user *User) map[string]interface{} {
	userStr, _ := json.Marshal(user)
	userMap := make(map[string]interface{})
	_ = json.Unmarshal(userStr, &userMap)
	delete(userMap, "created_at")
	delete(userMap, "updated_at")
	delete(userMap, "deleted_at")
	fmt.Println("userMap:", userMap)
	return userMap
}

func Md5Crypt(str string, salt ...interface{}) string {
	if l := len(salt); l > 0 {
		slice := make([]string, l+1)
		str = fmt.Sprintf(str+strings.Join(slice, "%v"), salt...)
	}
	return fmt.Sprintf("%x", md5.Sum([]byte(str)))
}
