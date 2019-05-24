package msg

import (
	"bufio"
	"fmt"
	"os"
	"reflect"

	gamePb "server/msg/gamepb"

	"github.com/name5566/leaf/network/protobuf"
)

var Processor = protobuf.NewProcessor()

var (
	f *os.File
	w *bufio.Writer
)

func init() {
	registerProcessor()

	Processor.Range(pl)
}

func registerProcessor() {
	Processor.Register(&gamePb.UserLogin{})
	Processor.Register(&gamePb.LoginResult{})
}

func pl(id int, t reflect.Type) {
	if id == -1 {
		w.Flush()
		f.Close()
	} else if id == 0 {
		filePath := os.Getenv("GOPATH") + "\\outPut\\GamePb.lua"
		var err error
		f, err = os.Create(filePath)

		if err != nil {
			fmt.Println("msg.go - pl : create file error !")

			return
		}
		w = bufio.NewWriter(f)
		s := fmt.Sprintf("%s = %d;", t.String(), id)
		fmt.Fprintln(w, s)
	} else {
		s := fmt.Sprintf("%s = %d;", t.String(), id)
		fmt.Fprintln(w, s)
	}

}
