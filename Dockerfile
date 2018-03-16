FROM golang:1.9
WORKDIR /
COPY truffletool.go .
RUN go get github.com/spf13/pflag
RUN go build truffletool.go

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
WORKDIR /data
RUN pip install truffleHog
COPY --from=0 /truffletool /usr/bin/truffletool
COPY scan.sh /scan.sh
COPY sensitive-regex.json /sensitive-regex.json
CMD ["/scan.sh"]
