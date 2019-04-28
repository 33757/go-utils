package utils

import (
    "bytes"
    "os/exec"
    "strconv"
)


func StringToInt(str string) (int, error)  {
    return strconv.Atoi(str)
}

func StringToInt64(str string) (int64, error) {
    return strconv.ParseInt(str,10,64)
}

func IntToString(num int) string {
    return strconv.Itoa(num)
}

func Int64ToString(num int64) string {
    return strconv.FormatInt(num,10)
}

func ExecBashShell(s string) (string, error){

    cmd := exec.Command("/bin/bash", "-c", s)
    var out bytes.Buffer
    cmd.Stdout = &out
    err := cmd.Run()

    return out.String(), err
}
