FROM golang:1.21.3

WORKDIR /opt/load-gen

COPY ./ /opt/load-gen/

RUN go build -o main main.go  && \
    mv -v main /usr/local/bin/load-gen

CMD [ "/usr/local/bin/load-gen" ]