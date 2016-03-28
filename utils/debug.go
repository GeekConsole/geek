package utils

import (
	"bytes"
	"encoding/json"
	"log"
)

const (
	FormatJsonBorder = "------------\n%s\n------------\n"
)

var (
	_dddebuggg = true
)

func SetDebug(b bool) {
	_dddebuggg = b
}

func IsDebug() bool {
	return _dddebuggg
}

func DebugF(format string, v ...interface{}) {
	if IsDebug() {
		log.Printf(format, v...)
	}
}

func DebugJSON(body []byte) {
	if IsDebug() {

		buf := bytes.NewBuffer(make([]byte, 0, len(body)+1024))
		if err := json.Indent(buf, body, "", "    "); err == nil {
			body = buf.Bytes()
		}
		log.Printf(FormatJsonBorder, body)
	}
}
