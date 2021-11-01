package util

import (
	"bytes"
	"regexp"
	"strings"
	"fmt"
)

func BytesCombine(pBytes ...[]byte) []byte {
	return bytes.Join(pBytes, []byte(""))
}


func DebugError() {

}



func NetbiosEncode(data []byte) string{

}

func ParseResponse(data string) string{
    data = strings.Replace(data, "\n", "", -1)
    regexp,_ := regexp.Compile("vE(.*)...l")
    t := regexp.FindString(data)
	if len(t) > 0{
		t = t[2:len(t)-4]
		fmt.Printf("resString: %v\n", t)
		return t
	}else{
		return ""
	}

}

func NetbiosDecode(data string) []byte{

}
