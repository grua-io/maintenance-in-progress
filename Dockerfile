FROM --platform=$BUILDPLATFORM golang:1.27.1-alpine AS build

ARG TARGETARCH

WORKDIR /src

COPY go.mod ./
COPY *.go index.html ./

RUN go test ./...
RUN CGO_ENABLED=0 GOOS=linux GOARCH=${TARGETARCH} go build -trimpath -ldflags='-s -w' -o /out/maintenance-in-progress .

#

FROM gcr.io/distroless/static-debian12:nonroot AS runtime

COPY --from=build /out/maintenance-in-progress /usr/bin/maintenance-in-progress

EXPOSE 8080

ENTRYPOINT ["/usr/bin/maintenance-in-progress"]
