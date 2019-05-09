package utils

import (
	"bytes"
	"encoding/binary"
	"net"
	"os"
	"os/exec"
	"reflect"
	"strconv"
	"strings"
)

// string 转 int
func StringToInt(str string) (int, error) {
	return strconv.Atoi(strings.TrimSpace(str))
}

// string 转 int64
func StringToInt64(str string) (int64, error) {
	return strconv.ParseInt(str, 10, 64)
}

// int 转 string
func IntToString(num int) string {
	return strconv.Itoa(num)
}

// int64 转 string
func Int64ToString(num int64) string {
	return strconv.FormatInt(num, 10)
}

// int to string
func IntToByte(num int) []byte {
	var buffer bytes.Buffer
	err := binary.Write(&buffer, binary.BigEndian, num)
	if err != nil {
		return []byte("")
	}
	return buffer.Bytes()
}

// 执行 shell 命令
func ExecBashShell(s string) (string, error) {

	cmd := exec.Command("/bin/bash", "-c", s)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()

	return out.String(), err
}

// ip 转 long / uint32
func Ip2Long(ipAddr string) uint32 {
	ip := net.ParseIP(ipAddr)
	if ip == nil {
		return 0
	}
	return binary.BigEndian.Uint32(ip.To4())
}

// long / uint32 转 ip
func Long2Ip(properAddress uint32) string {
	ipByte := make([]byte, 4)
	binary.BigEndian.PutUint32(ipByte, properAddress)
	ip := net.IP(ipByte)
	return ip.String()
}

// 判断是否为文件
func IsFile(filename string) bool {
	_, err := os.Stat(filename)
	if err != nil && os.IsNotExist(err) {
		return false
	}
	return true
}

// 判断是否为string
func IsString(val interface{}) bool {
	v := reflect.ValueOf(val)
	switch v.Kind() {
	case reflect.String, reflect.Array:
		return true
	}
	return false
}
