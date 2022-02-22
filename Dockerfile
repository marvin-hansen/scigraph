FROM  debian:stable-slim

WORKDIR /app

ARG DEBIAN_FRONTEND=noninteractive

RUN apt update

RUN apt-get -y -q install wget unzip default-jdk

RUN apt clean

# Download & Install instructions
# https://infinitegraph.com/free-download/
# https://infinitegraph.com/install-steps-linux/
RUN wget https://infinitegraph.com/wp-content/uploads/2021/08/zip-linux-gcc53-amd64-202130.zip

RUN unzip -qd /app zip-linux-gcc53-amd64-202130.zip
RUN rm zip-linux-gcc53-amd64-202130.zip

Add scripts/infinite/run.sh .
RUN chmod +x run.sh

EXPOSE 8190

CMD exec ./run.sh

# This kees the container running after the script completed,
# but will exit immediately on Ctr-C or a docker stop signal
CMD exec /bin/sh -c "trap : TERM INT; (while true; do sleep 1000; done) & wait"