# go-build

## 简介
`go-build` 是一个工程构建工具，用于完成复杂的工程构建和打包工作，支持 Windows、Linux 和 macOS 操作系统。该工具旨在简化多平台下的构建流程，使得不同操作系统下的工程构建变得更加一致和高效。

## 配置文件示例
以下是一个配置文件示例 `.build.txt`，包含不同操作系统的环境变量和构建命令：

```
[win.env]
set JAVA_HOME=D:\java\jdk1.8.0_121

[win.build]
mvn clean package -DskipTests -Pproduction

[mac.env]
export JAVA_HOME=~/java/jdk1.8.0_121

[mac.build]
mvn clean package -DskipTests -Pproduction

[linux.env]
export JAVA_HOME=~/usr/java/jdk1.8.0_121

[linux.build]
mvn clean package -DskipTests -Pproduction
```

## 问题描述
在开发工程中，不同的操作系统、不同的编程语言和项目管理工具有各自的构建命令，导致构建过程复杂且容易出错。

## 解决办法
1. 将构建命令放到配置文件中。
2. 编写一个程序，逐行读取文件中的命令并执行。
3. 处理环境变量设置和执行错误的情况。

## 如何使用
该工具已经开发完成并开源。以下是使用步骤：

1. 访问开源地址: [go-build GitHub 仓库](https://github.com/litongjava/go-build)
2. 下载最新版本: [v1.0.0 版本下载](https://github.com/litongjava/go-build/releases/tag/v1.0.0)
3. 将 `build` 命令添加到系统的 `PATH` 环境变量中，例如 `d:\bin` 目录。
4. 创建配置文件 `build.txt`，内容示例如上。
5. 执行命令 `build .build.txt` 进行构建和打包。

## 例子
假设您的配置文件名为 `.build.txt`，文件内容如下：

```
[win.env]
set JAVA_HOME=D:\java\jdk1.8.0_121

[win.build]
mvn clean package -DskipTests -Pproduction

[mac.env]
export JAVA_HOME=~/java/jdk1.8.0_121

[mac.build]
mvn clean package -DskipTests -Pproduction

[linux.env]
export JAVA_HOME=~/usr/java/jdk1.8.0_121

[linux.build]
mvn clean package -DskipTests -Pproduction
```

您可以在命令行中运行：

```sh
build .build.txt
```

该命令将读取配置文件中的环境变量和构建命令，根据当前操作系统设置环境变量并执行相应的构建命令。

## 贡献
欢迎对 `go-build` 工具提出意见和建议，您可以通过 GitHub 提交 issue 或 pull request。

感谢您使用 `go-build`，希望它能为您的开发工作带来便利！