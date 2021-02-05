FROM 10.253.84.44:1121/nodered-ci/golang:1.15.3 as builder
WORKDIR /app/
COPY . .
RUN ls
RUN go env -w GOPROXY=https://goproxy.cn
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags '-d -w -s ' -a -installsuffix cgo -o app .
RUN ls

FROM 10.253.84.44:1121/nodered-ci/busybox
WORKDIR /app/
ENV INK8S=true
COPY --from=builder /app/app .


CMD ["./app"]