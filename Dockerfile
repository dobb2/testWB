#FROM golang:1.20.3-bullseye as gobuild
#WORKDIR /app
#RUN go version
#COPY . ./
## build go app
#RUN go mod download
#RUN go build -o sub-client-app ./cmd/sub-client/main.go
## to reduce the image size with multi-stage assembly
#FROM scratch
#WORKDIR /app
#COPY --from=gobuild /app .
#RUN chmod +x sub-client-app
#
#CMD ["./sub-client-app"]