# go-build

## 简介
build 是一个工程构建工具,用于完成复杂的工程构建和打包工作,支持windows,linux and macos
## 配置文件示例
.build.txt
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
在开发工程中遇到的一个问题是工程构建,不同的操作系统,不同的代码语言,不同的项目管理工具有不同的构建命令

## 解决办法
- 1.将命令放到文件中
- 2.写一个程序,读取文件中的命令,一行一行执行
- 3.处理一下环境变量和执行错误的情况

## 如何使用
已经开发完成,开源地址https://github.com/litongjava/go-build
- 下载https://github.com/litongjava/go-build/releases/tag/v1.0.0
- 将build.exe 添加到PATH目录,笔者是d:\bin
- 添加build-win.txt笔者的内容如下
- 执行build . 命令进行部署

## 支持的平台
目前仅仅支持windows平台,没有必要支持linux和macos,它们的shell script脚本已经足够好用

## 其他
对于其他需要命令流程的工作本工具依然适用

## how to use

build-prod.txt
```text
set JAVA_HOME=D:\java\jdk1.8.0_121
mvn clean build -DskipTests -Pproduction
fly build
```

```shell
build build-prod.txt
```

output
```log
E:\code\project\imaginix\imaginix-kimi-service-monitoring>build .
2024/03/06 23:59:30.981964 main.go:49: add env variable: JAVA_HOME=D:\java\jdk1.8.0_121
2024/03/06 23:59:31.002341 main.go:63: Executing in . : mvn clean build -DskipTests -Pproduction
[INFO] Scanning for projects...
[INFO]
[INFO] -----------< com.imaginix:imaginix-kimi-service-monitoring >------------
[INFO] Building imaginix-kimi-service-monitoring 1.0.0
[INFO] --------------------------------[ jar ]---------------------------------
[INFO]
[INFO] --- maven-clean-plugin:2.5:clean (default-clean) @ imaginix-kimi-service-monitoring ---
[INFO] Deleting E:\code\project\imaginix\imaginix-kimi-service-monitoring\target
[INFO]
[INFO] --- maven-resources-plugin:2.6:resources (default-resources) @ imaginix-kimi-service-monitoring ---
[INFO] Using 'UTF-8' encoding to copy filtered resources.
[INFO] Copying 3 resources
[INFO]
[INFO] --- maven-compiler-plugin:3.1:compile (default-compile) @ imaginix-kimi-service-monitoring ---
[INFO] Changes detected - recompiling the module!
[INFO] Compiling 10 source files to E:\code\project\imaginix\imaginix-kimi-service-monitoring\target\classes
[INFO]
[INFO] --- maven-resources-plugin:2.6:testResources (default-testResources) @ imaginix-kimi-service-monitoring ---
[INFO] Using 'UTF-8' encoding to copy filtered resources.
[INFO] Copying 0 resource
[INFO]
[INFO] --- maven-compiler-plugin:3.1:testCompile (default-testCompile) @ imaginix-kimi-service-monitoring ---
[INFO] Changes detected - recompiling the module!
[INFO] Compiling 6 source files to E:\code\project\imaginix\imaginix-kimi-service-monitoring\target\test-classes
[INFO]
[INFO] --- maven-surefire-plugin:2.12.4:test (default-test) @ imaginix-kimi-service-monitoring ---
[INFO] Tests are skipped.
[INFO]
[INFO] --- maven-jar-plugin:2.4:jar (default-jar) @ imaginix-kimi-service-monitoring ---
[INFO] Building jar: E:\code\project\imaginix\imaginix-kimi-service-monitoring\target\imaginix-kimi-service-monitoring-1.0.0.jar
[INFO]
[INFO] --- spring-boot-maven-plugin:2.7.4:rebuild (default) @ imaginix-kimi-service-monitoring ---
[INFO] Replacing main artifact with rebuildd archive
[INFO] ------------------------------------------------------------------------
[INFO] BUILD SUCCESS
[INFO] ------------------------------------------------------------------------
[INFO] Total time:  3.354 s
[INFO] Finished at: 2024-03-06T23:59:35-10:00
[INFO] ------------------------------------------------------------------------
2024/03/06 23:59:35.386504 main.go:63: Executing in . : fly build
==> Verifying app config
Validating E:\code\project\imaginix\imaginix-kimi-service-monitoring\fly.toml
✓ Configuration is valid
--> Verified app config
==> Building image
Waiting for remote builder fly-builder-frosty-rain-1801...
Remote builder fly-builder-frosty-rain-1801 ready
Waiting for remote builder fly-builder-frosty-rain-1801...
Remote builder fly-builder-frosty-rain-1801 ready
==> Building image with Docker
--> docker host: 20.10.12 linux x86_64
[+] Building 34.2s (8/8) FINISHED
 => [internal] load build definition from Dockerfile                                                                                                                                                        0.8s
 => => transferring dockerfile: 32B                                                                                                                                                                         0.8s
 => [internal] load .dockerignore                                                                                                                                                                           0.8s
 => => transferring context: 2B                                                                                                                                                                             0.8s
 => [internal] load metadata for docker.io/litongjava/jre:8u391-stable-slim                                                                                                                                 3.9s
 => [internal] load build context                                                                                                                                                                          29.1s
 => => transferring context: 9.69MB                                                                                                                                                                        29.1s
 => [1/3] FROM docker.io/litongjava/jre:8u391-stable-slim@sha256:199bf2b76a3b8ad68faf6e5c108b1de3baa49366fcba8415e836f02368b855c9                                                                           0.0s
 => CACHED [2/3] WORKDIR /app                                                                                                                                                                               0.0s
 => [3/3] COPY target/imaginix-kimi-service-monitoring-1.0.0.jar /app/                                                                                                                                      0.0s
 => exporting to image                                                                                                                                                                                      0.0s
 => => exporting layers                                                                                                                                                                                     0.0s
 => => writing image sha256:ab2a0b2fa392e7cb1f7e4db7af99002aeb51c1a612fcbea584e58d61b7b12d0f                                                                                                                0.0s
 => => naming to registry.fly.io/imaginix-kimi-service-monitoring:buildment-01HRC5NDB2QFHVFAGKKB8RJFZE                                                                                                     0.0s
--> Building image done
==> Pushing image to fly
The push refers to repository [registry.fly.io/imaginix-kimi-service-monitoring]
858c1983017e: Pushed
0f5bc9e38484: Layer already exists
08187314d29d: Layer already exists
2b4482819176: Layer already exists
buildment-01HRC5NDB2QFHVFAGKKB8RJFZE: digest: sha256:00f942713a6b5a8e2d02a55ec113441536b108f8c64e09f89b17566f9a0da8dd size: 1159
--> Pushing image done
image: registry.fly.io/imaginix-kimi-service-monitoring:buildment-01HRC5NDB2QFHVFAGKKB8RJFZE
image size: 361 MB

Watch your buildment at https://fly.io/apps/imaginix-kimi-service-monitoring/monitoring

Updating existing machines in 'imaginix-kimi-service-monitoring' with rolling strategy
> Updating 32874d3df72768 [app]
> Updating 32874d3df72768 [app]
> Waiting for 32874d3df72768 [app] to have state: started
> Machine 32874d3df72768 [app] has state: started
> Checking that 32874d3df72768 [app] is up and running
> Waiting for 32874d3df72768 [app] to become healthy: 0/1

> Waiting for 32874d3df72768 [app] to become healthy: 1/1

✔ Machine 32874d3df72768 [app] update succeeded
Checking DNS configuration for imaginix-kimi-service-monitoring.fly.dev

Visit your newly builded app at https://imaginix-kimi-service-monitoring.fly.dev/
```
