FROM debian:stable-slim

ENV PORT=8991

COPY goserver /bin/docker_server

CMD ["/bin/docker_server"]

