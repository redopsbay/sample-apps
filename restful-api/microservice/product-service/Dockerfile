FROM golang:1.21.3

WORKDIR /opt/product-service

COPY ./ /opt/product-service/

RUN make build && \
    mv -v main /usr/local/bin/product-service

EXPOSE 8000

CMD [ "/usr/local/bin/product-service" ]