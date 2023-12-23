FROM golang:1.21.3

WORKDIR /opt/reviews-service

COPY ./ /opt/reviews-service/

RUN make build && \
    mv -v main /usr/local/bin/reviews-service

EXPOSE 8090

CMD [ "/usr/local/bin/reviews-service" ]