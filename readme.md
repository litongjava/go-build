# go-build

## Introduction
`go-build` is an engineering build tool designed to handle complex engineering build and packaging tasks. It supports Windows, Linux, and macOS operating systems. The tool aims to simplify the build process across multiple platforms, making engineering builds more consistent and efficient.

## Configuration File Example
Below is an example of a configuration file `.build.txt`, which includes environment variables and build commands for different operating systems:

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

## Problem Description
In engineering development, different operating systems, programming languages, and project management tools have their own build commands, making the build process complex and error-prone.

## Solution
1. Place the build commands in a configuration file.
2. Write a program to read and execute the commands line by line.
3. Handle environment variable settings and execution errors.

## How to Use
The tool has been developed and is open source. Follow these steps to use it:

1. Visit the open source repository: [go-build GitHub Repository](https://github.com/litongjava/go-build)
2. Download the latest version: [v1.0.0 Release](https://github.com/litongjava/go-build/releases/tag/v1.0.0)
3. Add the `build` command to your system `PATH` environment variable, e.g., `d:\bin`.
4. Create a configuration file `build.txt` with content similar to the example above.
5. Execute the command `build .build.txt` to build and package your project.

## Example
Assume your configuration file is named `.build.txt`, with the following content:

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

You can run the following command in the command line:

```sh
build .build.txt
```

This command will read the environment variables and build commands from the configuration file, set the environment variables based on the current operating system, and execute the corresponding build commands.

## Contributing
We welcome feedback and suggestions for the `go-build` tool. You can submit issues or pull requests through GitHub.

Thank you for using `go-build`. We hope it makes your development work easier and more efficient!