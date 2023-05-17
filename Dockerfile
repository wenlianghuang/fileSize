FROM golang:1.19-buster

WORKDIR /app 

COPY go.* ./
RUN go mod download

COPY main.go .

#RUN go build -v -o main 

CMD ["go","run","main.go"]