FROM golang

ARG app_env
ENV APP_ENV $app_env

RUN apt-get update && apt-get upgrade;\
	apt-get -y install libssl-dev;
COPY . /go/src/github.com/sovrinbloc/kairos
WORKDIR /go/src/github.com/sovrinbloc/kairos

RUN go get ./
RUN go build *.go

CMD ./server; 
  
EXPOSE 8074
