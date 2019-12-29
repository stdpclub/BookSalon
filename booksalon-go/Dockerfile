FROM golang

WORKDIR /go/src

# 设置go module源。解决国内无法连接问题
RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.io
RUN go env -w GOSUMDB=sum.golang.google.cn

RUN go get -u github.com/gin-gonic/gin
RUN go get github.com/urfave/cli/v2
RUN go get -u github.com/jinzhu/gorm

COPY . .

EXPOSE 80

CMD [ "go build", "-o booksalon . ", " && ./booksalon"]