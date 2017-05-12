FROM golang:1.8.1

ENV PLUME_HOME=/opt/pyroclast-plumes

RUN git clone https://github.com/onyx-platform/pyroclast-plumes.git /opt/pyroclast-plumes

WORKDIR "$PLUME_HOME/file-plume"
ADD docker-config.yml /opt/pyroclast-plumes

RUN mkdir /data

ENV TOPIC_FILE=topic.yml

RUN go get github.com/elastic/beats || true
RUN go build plume.go

ENTRYPOINT /opt/pyroclast-plumes/file-plume/plume -strict.perms=false -e -c /opt/pyroclast-plumes/topics/$TOPIC_FILE -c /opt/pyroclast-plumes/docker-config.yml
