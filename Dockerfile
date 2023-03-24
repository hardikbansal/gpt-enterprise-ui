FROM golang:latest
WORKDIR /server
#CMD ["go","mod","help"]
COPY . .
RUN go build main.go
CMD ["/server/main"]