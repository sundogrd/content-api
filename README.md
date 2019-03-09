# creator-service(sundogrd)  ![travis-build](https://api.travis-ci.org/sundogrd/content-api.svg?branch=master)
A Go RESTful API server with gin and docker

## 安装
安装最新的[Go](https://golang.org/)，设置好`$GOPATH`，将`$GOPATH/bin`设置为可执行PATH。这里使用的是Mac或者Linux环境。

````
$ cd $GOPATH/src
$ git clone https://github.com/sundogrd/content-api.git
````

## 启动开发环境
建议用docker环境开发和部署，下面构建的docker镜像，会安装此脚本。（不知道要不要装air或者gin这种hot reload）
启动服务：
```
$ make dev
```
测试服务：`http://localhost:8086/hello`

## 构建和启动正式环境
```
$ make build
$ make start
```

## 清理
```
$ make clean
```

## Docker
- 安装[Docker](https://www.docker.com/) (请设置阿里云等镜像加速器, 不用加速器拉镜像速度很慢)
- 安装[Docker Compose](https://github.com/docker/compose/releases) (一般在Windows和Mac下安装了Docker，同时会安装docker-compose)

> 开发环境是通过源代码的方式运行调试；测试和正式环境镜像只有可执行文件和配置文件，无源代码。

### 构建和启动docker开发环境
生成开发用的docker镜像，开发环境将项目根目录作为数据卷挂载到容器中。
```
$ make start-docker-dev
```
或者在项目根文件夹下通过以下命令启动：
````
$ docker-compose up -d
````

第一次启动会稍微花费点时间自动构建docker镜像，成功构建镜像后，下次启动不会再花费时间构建镜像。 

通过docker-compose启动api web, mysql, redis三个容器，api web端口为8086, mysql端口为3306, redis端口为6379。数据库可通过本地客户端工具连接进行操作和调试。

测试接口是否正常启动：请访问`http://localhost:8086/hello`看是否有反馈信息。

### 停止开发环境
````
$ make stop-docker-dev
````
或者在项目根文件夹下运行命令：
````
$ docker-compose down
````

### 生成测试环境docker镜像
修改`Makefile`，可以修改镜像名称
````
$ make docker-image-staging
````

### 生成生产环境docker镜像
修改`Makefile`，可以修改镜像名称
````
$ make docker-image
````

## 包依赖管理
包依赖管理使用Go1.11出的Modules

## 注意
- 本脚手架使用[gin](https://github.com/gin-gonic/gin)作为HTTP框架，使用方式请查看官网。也可以换成其它HTTP框架，使用方式类似。
- [golint](https://github.com/golang/lint)为Go语言的Lint工具。详情请看官网。可以设置代码编辑器支持`golint`动态检测。
- 一些配套工具比如配置文件、日志、数据库、reids等封装库，放到[sundog-api-lib]()，
- 主配置文件采用json格式，文件位置`data/config/app.json`，配置库使用了[viper](https://github.com/spf13/viper)，详情请查看官网。
- 日常开发，在`handler/`下放置路由的处理函数，在`routes/`下配置路由，一般按功能分不同文件分开。
- `docker/development` 为开发环境docker构建文件，`docker/staging`为测试环境docker-compose文件，`docker/production`为生产环境docker-compose文件。
- 一般正式发布下服务器部署只包含`data/`目录（存放配置文件、日志、数据库文件等作为docker数据卷加载），`docker-compose.yml`（启动docker容器的编排文件)。
- 一般正式发布下的docker镜像里的项目目录下，只有`Makefile`(make脚本)，`bin/`(编译过的可执行文件，启动服务的主执行文件)，`data/`(挂载的数据卷，存放配置文件、日志、数据库文件等)，部署出去的是编译好的可执行文件，相对安全独立。



## 私货
只有现在只接鼓掌日志统计（以后日志统计是否可以分库）？