FROM golang
LABEL maintainer=niaddice
COPY . /$GOPATH/src/go-starter/
WORKDIR /$GOPATH/src/go-starter/
#设置环境变量，开启go module和设置下载代理
RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct
#增加缺失的包，移除没用的包
RUN go mod tidy
RUN go build main.go
EXPOSE 8092:8092
CMD ["go","run","main.go"]