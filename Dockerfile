FROM golang

ENV GOOS=linux
ENV GOARCH=amd64
ENV CGO_ENABLED=0

RUN mkdir -p /workspace
WORKDIR /workspace
ADD go.mod go.sum ./
RUN go mod download
ADD . .
RUN go build -trimpath -o .build/qrserver -ldflags "-w -s" .

# ---------------------------------------------------------------------------------

FROM alpine

RUN mkdir -p /app
WORKDIR /app

COPY --from=0 /workspace/.build/* ./
ENTRYPOINT ["/app/qrserver"]
