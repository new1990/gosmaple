FROM golang:latest


EXPOSE 8080

CMD ["go", "run", "main.go"]