FROM golang:1.8.1

ENV PLUME_HOME=/opt/pyroclast-plumes

RUN git clone https://github.com/onyx-platform/pyroclast-plumes.git /opt/pyroclast-plumes

WORKDIR "$PLUME_HOME/file-plume"

ENV TOPIC_FILE=topic.yml
ENV CONFIG_FILE=config.yml

RUN go get github.com/elastic/beats || true
RUN go build plume.go

ENTRYPOINT /opt/pyroclast-plumes/file-plume/plume -strict.perms=false -e -c /opt/pyroclast-plumes/config/$TOPIC_FILE -c /opt/pyroclast-plumes/config/$CONFIG_FILE
