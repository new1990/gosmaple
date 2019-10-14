FROM golang:latest

RUN go get github.com/gin-gonic/gin
RUN go get github.com/jinzhu/gorm
RUN go get github.com/jinzhu/gorm/dialects/postgres
RUN go get github.com/new1990/gosmaple/db
RUN go get github.com/new1990/gosmaple/entity

EXPOSE 8080

CMD ["go", "run", "main.go"]