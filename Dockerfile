FROM golang:1.12.7

#ENV GOBIN /go/bin
#WORKDIR /go/src/Trx-service
#ADD . ./

RUN go get github.com/astaxie/beego
RUN go get github.com/beego/bee
RUN go get github.com/lib/pq
RUN go get github.com/joho/godotenv

#RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix .
#RUN bee run

EXPOSE 8080
#ENTRYPOINT ./Trx-service
CMD ["bee", "run"]