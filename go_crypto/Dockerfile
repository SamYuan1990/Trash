From golang:1.21

RUN mkdir -p /app
WORKDIR /app
COPY main.go .
COPY main_test.go .
COPY go.mod .

#ENV GOPROXY=https://goproxy.cn,direct

RUN go mod tidy
RUN go build main.go
#RUN ls main

CMD ["go","test","-v","-bench=.","--count=10"]
