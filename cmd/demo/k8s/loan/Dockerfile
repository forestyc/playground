FROM m.daocloud.io/docker.io/alpine:3.20.3

WORKDIR /app/loan
RUN mkdir etc

COPY loan .
COPY etc/loan.yaml ./etc

EXPOSE 8080

ENTRYPOINT ["./loan"]