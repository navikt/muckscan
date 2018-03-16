FROM ubuntu:16.04
WORKDIR /software
RUN apt-get update && \
    apt-get install -y \
        git \
        build-essential \
        python \
        python-pip \
        && \
    rm -rf /var/lib/apt/lists
RUN git clone --depth=1 https://github.com/awslabs/git-secrets
#RUN git clone https://github.com/auth0/repo-supervisor
#RUN git clone --depth=1 https://github.com/dxa4481/truffleHog
WORKDIR /software/git-secrets
RUN make install
WORKDIR /data
RUN pip install truffleHog
