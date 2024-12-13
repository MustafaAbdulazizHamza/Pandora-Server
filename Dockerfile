FROM debian:latest
WORKDIR /root/
COPY Pandora .
COPY Pandora.db .
COPY server.crt .
COPY server.key .
EXPOSE 8080
CMD ["./Pandora"]