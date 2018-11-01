package debug

import (
	"fmt"
	"io"
	"os"
	"runtime"
	"strconv"
	"strings"
	"sync/atomic"
)

var aid *atomic.Value

func init() {
	aid = new(atomic.Value)
	aid.Store(1)
}

func Println(a ...interface{}) {
	Fprintln(os.Stdout, a...)
}

func Fprintln(w io.Writer, a ...interface{}) {
	id := aid.Load().(int)
	aid.Store(id + 1)

	pc, file, line, ok := runtime.Caller(1)
	if !ok {
		fmt.Println(append([]interface{}{id}, a...)...)
	}

	name := ""
	names := strings.Split(runtime.FuncForPC(pc).Name(), ".")
	if len(names) > 3 {
		name = strings.Join(names[len(names)-2:], ".")
	} else {
		name = strings.Join(names[len(names)-1:], ".")
	}

	fileList := strings.Split(file, "/")
	fmt.Fprintln(w, append([]interface{}{id, fileList[len(fileList)-1] + ":" + strconv.Itoa(line), name}, a...)...)
}

/*

不懂
学习
待续

*/
