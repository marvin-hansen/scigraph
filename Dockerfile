FROM  debian:stable-slim
ARG DEBIAN_FRONTEND=noninteractive

# install the necessary dependencies
RUN apt-get update
RUN apt-get -y -q install unzip curl vim-tiny apt-transport-https ca-certificates wget dirmngr gnupg software-properties-common

# How to Install Java 8 on Debian 10 Linux
# https://linuxize.com/post/install-java-on-debian-10/

# Import the repository’s GPG key using the following wget command:
RUN wget -qO - https://adoptopenjdk.jfrog.io/adoptopenjdk/api/gpg/key/public | apt-key add -

# Add the AdoptOpenJDK APT repository to your system:
# https://linuxize.com/post/install-java-on-debian-10/
RUN add-apt-repository --yes https://adoptopenjdk.jfrog.io/adoptopenjdk/deb/

# Once the repository is enabled, update apt sources and install Java 8 using the following commands:
RUN apt-get update
RUN apt-get -y -q install adoptopenjdk-8-hotspot

# Clean up after install
RUN apt clean

WORKDIR /app
# Download & Install instructions:
# https://infinitegraph.com/install-steps-linux/
RUN wget https://infinitegraph.com/wp-content/uploads/2021/08/zip-linux-gcc53-amd64-202130.zip
RUN unzip -qd /app zip-linux-gcc53-amd64-202130.zip
RUN rm *.zip

# Update env. variable
ENV PATH="/app/bin:$PATH"

COPY scripts/infinite/restApiConfig.xml restApiConfig.xml
COPY scripts/infinite/entrypoint.sh entrypoint.sh
COPY scripts/infinite/restart.sh restart.sh
RUN chmod +x *.sh

EXPOSE 8190 8185

ENTRYPOINT ["./entrypoint.sh"]