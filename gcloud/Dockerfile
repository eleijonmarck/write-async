FROM python:3-slim-buster

SHELL ["/bin/bash", "-o", "pipefail", "-c"]

RUN mkdir -p /usr/share/man/man1 && \
    apt-get update && \
    apt-get install -qy apt-transport-https ca-certificates curl gnupg openjdk-11-jre openssh-client sudo && \
    rm -rf /var/lib/apt/lists/*

# Add non root user
RUN addgroup --gid 1000 dev && \
    adduser --disabled-password --ingroup dev --uid 1000 dev && \
    echo "dev ALL=(ALL) NOPASSWD: ALL" >>/etc/sudoers.d/50-dev && \
    echo 'Defaults env_keep += "DEBIAN_FRONTEND"' >>/etc/sudoers.d/env_keep
USER dev
ENV WORKDIR=/home/dev
WORKDIR $WORKDIR

# Install Cloud SDK
RUN curl -s https://dl.google.com/dl/cloudsdk/channels/rapid/downloads/google-cloud-sdk-333.0.0-linux-x86_64.tar.gz | tar xvz
RUN ./google-cloud-sdk/install.sh --quiet --command-completion=true --path-update=true --usage-reporting=false
ENV PATH="${WORKDIR}/google-cloud-sdk/bin:${PATH}"

# Install gcloud components
RUN gcloud components update --quiet && \
    gcloud components install --quiet beta pubsub-emulator

EXPOSE 5000

