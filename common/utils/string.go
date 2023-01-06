package utils

import (
	"math/rand"
	"strings"
	"time"
)

var (
	r                  *rand.Rand
	randomBytes        = []byte("0123456789abcdefghijklmnopqrstuvwxyz-") //ABCDEFGHIJKLMNOPQRSTUVWXYZ
	randomBytesWithout = []byte("0123456789abcdefghijklmnopqrstuvwxyz")  //ABCDEFGHIJKLMNOPQRSTUVWXYZ
)

func init() {
	r = rand.New(rand.NewSource(time.Now().Unix()))
}

// RandString 生成随机字符串
func RandString(lens int) string {
	result := []byte{}
	for i := 0; i < lens; i++ {
		if i == 0 || i == lens-1 {
			result = append(result, randomBytesWithout[r.Intn(len(randomBytesWithout))])
			continue
		}
		result = append(result, randomBytes[r.Intn(len(randomBytes))])
	}
	return string(result)
}

// JoinString 拼接字符串
func JoinString(s ...string) string {
	// strings.Join(s, "")
	var b strings.Builder
	for _, str := range s {
		b.WriteString(str)
	}
	return b.String()
}
