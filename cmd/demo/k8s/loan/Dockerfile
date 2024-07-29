FROM alpine

WORKDIR /app/houseloan
RUN mkdir etc

COPY houseloan .
COPY etc/houseloan.yaml ./etc

EXPOSE 8080

ENTRYPOINT ["./houseloan"]