FROM golang:1.16-alpine

RUN apk add make gnuplot

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY * ./

RUN make generate-results
RUN make generate-graph

CMD ["cat", "out/results.csv"]