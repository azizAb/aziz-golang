#build stage

FROM golang:alpine AS build-env
RUN apk --no-cache add build-base git bzr mercurial gcc
ADD . /go/src/ms-account
WORKDIR /go/src/ms-account
RUN go get ms-account
RUN go install
RUN cd /go/src/ms-account && go build -o ms-account

 # final stage

FROM alpine
WORKDIR /go/src/ms-account
COPY --from=build-env /go/src/ms-account /go/src/ms-account
ENTRYPOINT ./ms-account
EXPOSE 9000