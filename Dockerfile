FROM ubuntu

RUN apt-get update && apt-get upgrade -y

RUN apt-get install -y \
    git \
    golang \
    nodejs \
    npm \
    postgresql \
    postgresql-client

RUN mkdir /app
WORKDIR /app
ENV GOPATH /app/

RUN go get gopkg.in/macaron.v1 

CMD go run main.go
