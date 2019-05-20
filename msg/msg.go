package msg

import (
	"github.com/name5566/leaf/network/protobuf"
)

var Processor = protobuf.NewProcessor()

func init() {
	// Processor.Register(&Hello{})
	// Processor.Register(&UserLogin{})

	Processor.Register(&Hello{})
}

// type Hello struct {
// 	Name string
// }

// 账号协议
type UserLogin struct {
	LoginName string
	LoginPwd  string
}
