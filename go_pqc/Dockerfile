From ghcr.io/samyuan1990/liboqs-go:latest

RUN mkdir -p /app
WORKDIR /app
COPY main.go .
COPY go.mod .
COPY go.sum .
COPY vendor .
#ENV GO111MODULE=auto
ENV GOPROXY=https://goproxy.cn,direct

#RUN go mod tidy
RUN go build main.go
#RUN ls main

CMD [main]
