FROM golang:latest 
RUN mkdir /app 
RUN go get github.com/gin-gonic/gin && go get github.com/go-sql-driver/mysql 
ADD . /app/
WORKDIR /app
RUN go build -o main . 
EXPOSE 8000
CMD ["/app/main"]
