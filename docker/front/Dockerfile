FROM node:14

WORKDIR /go/src/github.com/usagiga/Incipit

ADD . .

RUN cd ./front && npm i

CMD ["make", "front_run"]
