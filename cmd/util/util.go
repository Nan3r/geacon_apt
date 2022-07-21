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
   q := ""
   for _,value := range data{
        q += string((int(value)>>4) + int('a'))+string((int(value)&0xf + int('a')))
   }
    return q
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
    c := 0
    a := 0
    resa := []byte{}
   for index,value := range data{
       //fmt.Println(int(value))
        if ((index+1) % 2 == 0){
            c = (int(value)-int('a')) &0xf
            a += c 
           resa = append(resa, byte(a))
        }else{
           a = (int(value)-int('a')) << 4 
        }
   }
   return resa
}
