# Building on top of Ubuntu 14.04. The best distro around.
FROM ubuntu:14.04

COPY ./ /opt/
EXPOSE 80

ENTRYPOINT ["/opt/go-ecs-ecr"]
