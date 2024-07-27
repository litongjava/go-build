package main

import (
  "fmt"
  "log"
  "os"
)

func init() {
  // 设置Flags为 日期 时间 微秒 文件名:行号
  log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile)
}

func main() {
  var filePath = ".build.txt"
  if len(os.Args) < 2 {
    fmt.Println("use default file .build.txt")
  } else {
    filePath = os.Args[1]
  }

  Build(filePath)
}
