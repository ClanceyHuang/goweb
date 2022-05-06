# goweb

## 项目结构

```markdown
goweb:
.
├── build
│   └── package
│       └── dockerfile
├── cmd
│   ├── GoWeb-linux
│   ├── GoWeb-windows.exe
│   └── main.go
├── configs
│   └── config.json
├── go.mod
├── internal
│   ├── config
│   │   ├── config.go
│   │   └── config_test.go
│   ├── controller
│   │   ├── controller.go
│   │   ├── index.go
│   │   ├── look.go
│   │   └── welcome.go
│   ├── log
│   │   └── log.go
│   ├── middleware
│   │   ├── cross.go
│   │   ├── log.go
│   │   └── timeout.go
│   └── templates
│       └── template.go
├── LICENSE
├── README.md
├── scripts/ # 简单的脚本用于部署与维护
│   ├── linux
│   │   ├── build_to_linux.sh
│   │   ├── build_to_windows.sh
│   │   ├── run_background_linux.sh
│   │   ├── run_linux.sh
│   │   └── stop_run.sh
│   └── windows
│       ├── build_to_linux.bat
│       └── build_to_windows.bat
└── web
    ├── img
    │   ├── favicon.ico
    │   ├── golang-down.png
    │   ├── golang.png
    │   └── golang-right.png
    ├── plugins
    │   ├── bootstrap
    │   │   ├── css
    │   │   │   └── bootstrap.min.css
    │   │   └── js
    │   │       └── bootstrap.min.js
    │   └── jquery
    │       └── js
    │           └── jquery.min.js
    └── templates
        ├── index.tmpl
        ├── look.tmpl
        ├── test.tmpl
        └── welcome.tmpl
```

## 部署指南

### 编译
首先你得确保你安装好了 `Golang` 的环境。

如果你修改了项目，那么你就需要重新编译，得到可执行文件，才能部署运行。

`Golang` 的交叉编译提供了很大的便利，这使得我们可以在 `Windows` 系统下编译出可以在 `Linux` 环境下运行的可执行文件。

* 如果你是 `Windows` 平台下进行编译
  * 项目下 `scripts/windows` 文件夹下运行 `build_to_windows.bat`，将会在项目目录下生成 `GoWeb-windows.exe` 文件。项目根目录中初始带了这个文件，再次编译会覆盖根目录中原有的同名文件。
  * 项目下 `scripts/windows` 文件夹下运行 `build_to_linux.bat`，将会在项目目录下生成 `GoWeb-linux` 文件。该文件可以在 `linux` 系统下运行。再次编译会覆盖项目根目录中原有的同名文件。
* 如果你是 `Linux` 平台下进行编译
  * 项目下 `scripts/linux` 文件夹下运行 `build_to_linux.sh`（设置权限为可运行才能运行），执行后项目根目录下生成 `GoWeb-linux` 可执行文件。

### 配置
本项目使用 `JSON` 文件来进行配置。

查看 `config/conifg.json` 配置语义如下表，根据需求修改。

|配置字段|语义|
|:--|:--|
|static	|静态文件的路径|
|template|	go 模板的路径|
|address|	服务启动的地址 (默认为空，可以实现公网部署)|
|port|	服务启动的端口|
|pprofPort|	性能监控的服务的端口|
|handleTimeoutSecond|	request 处理超时时间（该配置在项目中未使用，作保留）|
|trace|	trace logger 输出日志的路径（项目中未使用，作保留）|
|info|	info logger 输出日志的路径|
|warning|	warning logger 输出日志的路径|
|error|	error logger 输出日志的路径|


### 运行
如果你是 `Windows` 平台下运行

点击项目根目录下的 `GoWeb-windows.exe` 开始运行。

如果你是 `Linux` 平台下运行

`Linux` 下分为一般运行，和后台运行和作为服务运行，在 `scripts/linux` 中准备了前两种运行方式的脚本和停止程序的脚本。如果你想作为服务运行，可以参照简介中提到的视频来完成。

运行 `run_linux.sh` 。将程序运行在终端，随着终端的关闭，服务也就关了。

运行 `run_background_linux.sh` 。将程序在后台运行，即使终端关闭，也将继续运行。

你可以通过 ps -aux | grep GoWeb 来查看后台的服务跑起来没有

运行 stop_run.sh 。会把后台的服务关掉。

### 访问
本机可以访问 http://localhost:8080 或 http://localhost:8081/debug/pprof。

前者是首页，后者是性能监控界面。

如果你是部署在服务器上，想要通过公网访问，记得设置防火墙开放相应的端口。如果是云服务器还要打开安全组。


## 🐋 使用 Docker 运行

develop 分支已更新 dockerfile，并使用 Dorker 对 goweb 的部署运行进行了测试。

以下是 goweb 的 Docker 部署运行指南。

### 🔨 安装 Docker

Docker 是什么、怎么装、怎么用，这里不做详细的讲解介绍。

详细可参见 Docker 官方文档：[https://docs.docker.com/](https://docs.docker.com/)

### 📝 准备镜像

在项目的 `/build/package` 目录下有准备好的 dockerfile 用于生成镜像文件。进入 `/build/package` 目录下，使用以下指令生成镜像：

```bash
docker build -t goweb -f ./dockerfile ../../
```

其中 `-t ` 后面跟的参数 `goweb` 是镜像名，`-f` 后跟的参数是 dockerfile 的位置，`../` 指定构建的工作目录（在本项目中必须指定到项目目录 `/goweb` ）。

使用以下指令查看生成的镜像：

```bash
docker images
```

### ⚙ 运行容器

后台运行容器输入以下指令：

```bash
docker run -d -p 8080:8080 -p 8081:8081 goweb
```

想在控制台与容器交互使用以下指令：

```bash
 docker run -it -p 8080:8080 -p 8081:8081 goweb
```

注意，`-p 8080:8080` 和 `-p 8081:8081` 是将容器的端口映射到本机的端口。**其语义是 `-p <要映射到本机的端口>:<容器端口>`**。

此时（新开一个终端）输入以下指令可以查看正在运行的容器：

```bash
docker ps
```

### 🔍 访问

访问还是照常访问。

本机可以访问 http://localhost:8080 或 http://localhost:8081/debug/pprof。

前者是首页，后者是性能监控界面。

### 实用 Docker 指令

查看所有容器

```bash
docker ps -a
```

停止容器运行

```bash
docker stop [CONTAINER_ID]
```

启动停止的容器

```bash
docker start [CONTAINER_ID]
```

删除容器

```bash
docker rm [CONTAINER_ID]
```

查看所有镜像

```bash
docker images -a
```

删除镜像

```bash
docker rmi [IMAGE_ID]
```

执行容器的 bash（该容器必须得有 bash）以与容器交互

```bash
docker exec -it [CONTAINER_ID] /bin/bash
```

-------
