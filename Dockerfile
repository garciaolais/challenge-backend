FROM golang:alpine AS build
WORKDIR /go/bin
COPY . .
RUN apk add --no-cache make
RUN make build

#################
FROM scratch
COPY --from=build /go/bin /go/bin
ENTRYPOINT [ "/go/bin/challenge" ]