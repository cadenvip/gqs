# Go Quick Start

## Prepare

### Install

#### vagrant

使用 vagrant 来作为虚拟机开发，统一环境（与生产环境也尽量一致）


#### Go

```
安装 Go
$ brew install go

创建工作目录
$ mkdir -p mygo/src mygo/bin mygo/pkg

将工作目录添加到环境里面，使其生效，如果是用的 bash，则讲 zshrc 替换成 bashrc
$ echo "export GOPATH=$HOME/mygo" >> ~/.zshrc
$ echo "export PATH=$PATH:$GOPATH/bin" >> ~/.zshrc
$ source ~/.zshrc

输出 go 环境看一下是否正确
$ go env
```

#### Echo 

下载这个的时候需要开启代理。

```
go get -u github.com/labstack/echo
```

#### gRPC
```
$ wget https://github.com/google/protobuf/releases/download/v3.2.0/protoc-3.2.0-osx-x86_64.zip
$ unzip protoc-3.2.0-osx-x86_64.zip
$ cd protoc-3.2.0-osx-x86_64
$ echo "$PATH:`pwd`/bin" >> ~/.zshrc
$ go get google.golang.org/grpc
$ go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
```

### 架构

#### App

在 `apps` 下面分布不同的 App

#### services

在 `services` 下面有有提供的 `gRPC` server 端

#### public

`public` 下有HTML

#### static

`static` 下面 js，css 这些文件。






