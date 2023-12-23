FROM golang:1.21.3

WORKDIR /opt/ratings-service

COPY ./ /opt/ratings-service/

RUN make build && \
    mv -v main /usr/local/bin/ratings-service

EXPOSE 8080

CMD [ "/usr/local/bin/ratings-service" ]