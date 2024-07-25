package main

import (
  "bufio"
  "fmt"
  "log"
  "os"
  "os/exec"
  "runtime"
  "strings"
)

func init() {
  // 设置Flags为 日期 时间 微秒 文件名:行号
  log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile)
}

func main() {
  if len(os.Args) < 2 {
    fmt.Println("Usage: deploy file")
    return
  }

  filePath := os.Args[1]
  if runtime.GOOS != "windows" {
    log.Fatalln("not support current os")
    return
  }

  file, err := os.Open(filePath)
  if err != nil {
    log.Fatalln("Error opening file:", err)
  }
  defer file.Close()

  envVariables := []string{}

  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    line := scanner.Text()

    // 检查是否是设置环境变量的命令
    if strings.HasPrefix(line, "set ") {
      value := line[4:]
      envVariables = append(envVariables, value)
      log.Println("add env variable:", value)
      continue // 跳过执行此命令
    }

    executeCommand(line, envVariables)
  }

  if err := scanner.Err(); err != nil {
    log.Fatalln("Error reading file:", err)
  }
}

// executeCommand 在指定目录下执行一条命令，并应用所有以前设置的环境变量
func executeCommand(commandStr string, envVariables []string) {
  log.Println("Executing in", ":", commandStr)

  cmd := exec.Command("cmd", "/C", commandStr)

  // 添加之前设置的环境变量到命令中
  currEnv := os.Environ()
  for _, env := range envVariables {
    currEnv = append(currEnv, env)
  }
  cmd.Env = currEnv

  cmd.Stdout = os.Stdout
  cmd.Stderr = os.Stderr
  err := cmd.Run()
  if err != nil {
    log.Fatal("Error executing command:", err)
  }
}
