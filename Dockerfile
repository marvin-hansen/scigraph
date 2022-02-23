FROM  debian:stable-slim
ARG DEBIAN_FRONTEND=noninteractive

RUN apt update
RUN apt-get -y -q install wget unzip default-jdk
RUN apt clean

WORKDIR /app
# Download & Install instructions:
# https://infinitegraph.com/install-steps-linux/
RUN wget https://infinitegraph.com/wp-content/uploads/2021/08/zip-linux-gcc53-amd64-202130.zip
RUN unzip -qd /app zip-linux-gcc53-amd64-202130.zip
RUN rm *.zip

COPY scripts/infinite/entrypoint.sh entrypoint.sh
COPY scripts/infinite/restart.sh restart.sh
RUN chmod +x *.sh

EXPOSE 8190 8185

ENTRYPOINT ["./entrypoint.sh"]