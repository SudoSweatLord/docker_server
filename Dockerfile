FROM debian:stable-slim

# COPY source destination
COPY goserver /bin/docker_server

#Automatically starts the server process in the container when it's running
CMD ["/bin/docker_server"]
