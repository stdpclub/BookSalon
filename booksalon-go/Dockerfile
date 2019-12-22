FROM golang

WORKDIR /go/src

RUN git clone https://github.com/go-playground/validator.git    ~/go/src/gopkg.in/go-playground/validator.v8
RUN git clone https://github.com/go-yaml/yaml.git               ~/go/src/gopkg.in/yaml.v2
RUN git clone https://github.com/urfave/cli.git                 ~/go/src/github.com/urfave/cli/v2
RUN git clone https://github.com/gin-gonic/gin.git              ~/go/src/github.com/gin-gonic/gin
RUN git clone https://github.com/jinzhu/gorm.git                ~/go/src/github.com/jinzhu/gorm


COPY . .

EXPOSE 80

CMD [ "go build", "-o booksalon ./go/src/internal/", " && ./booksalon"]