FROM  debian:stable-slim

WORKDIR /app

ARG DEBIAN_FRONTEND=noninteractive

RUN apt update

RUN apt-get -y -q install wget unzip default-jdk

RUN apt clean

RUN wget https://infinitegraph.com/wp-content/uploads/2021/08/zip-linux-gcc53-amd64-202130.zip

RUN unzip -qd /app zip-linux-gcc53-amd64-202130.zip

Add docker_run.sh .

RUN mv docker_run.sh run.sh

RUN chmod +x run.sh

EXPOSE 8190

CMD ["./run.sh"]