FROM golang:latest

RUN go get github.com/gin-gonic/gin
RUN go get github.com/jinzhu/gorm
RUN go get github.com/jinzhu/gorm/dialects/postgres
RUN go get github.com/new1990/gosmaple/db
RUN go get github.com/new1990/gosmaple/entity
RUN go get github.com/new1990/gosmaple/service
RUN go get github.com/new1990/gosmaple/controller
RUN go get github.com/new1990/gosmaple/server
RUN go get github.com/derekparker/delve/cmd/dlv

EXPOSE 8080

CMD ["go", "run", "main.go"]