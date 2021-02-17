FROM golang:alpine as builder
WORKDIR /app/
COPY . .
RUN ls
#RUN go env -w GOPROXY=https://goproxy.cn
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags '-d -w -s ' -a -installsuffix cgo -o app .
RUN ls

FROM alpine
RUN apk add --no-cache ca-certificates  tzdata git
WORKDIR /app/
COPY --from=builder /app/app .


CMD ["./app"]