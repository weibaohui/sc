FROM golang:alpine as builder
LABEL website = github.com/weibaohui/sc
WORKDIR /app/
COPY . .
RUN ls
#RUN go env -w GOPROXY=https://goproxy.cn
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags '-d -w -s ' -a -installsuffix cgo -o sc .
RUN ls

FROM alpine
RUN apk add --no-cache ca-certificates  tzdata git
WORKDIR /app/
COPY --from=builder /app/sc .

ENTRYPOINT ["./sc"]
