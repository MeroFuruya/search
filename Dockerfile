ARG GO_VERSION=1.23.0
FROM golang:${GO_VERSION}-alpine AS build

COPY go.mod go.sum ./
RUN go mod download -x

COPY . .
RUN go build -o=/search .

FROM alpine:3.12

COPY --from=build /search /bin/search

ENV SEARCH_HOST=0.0.0.0
ENV SEARCH_PORT=3000
EXPOSE $SEARCH_PORT

HEALTHCHECK --interval=30s --timeout=30s --start-period=5s --retries=3 CMD [ "wget", "-q", "http://localhost:3000/health" ]

CMD [ "search" ]