### Backend
FROM golang:latest AS backend-build

WORKDIR /cardboard-companion

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build

FROM node:latest

RUN apt-get update && apt-get install -y supervisor

ENV GIN_MODE=release
WORKDIR /code

COPY supervisord.conf .
COPY --from=backend-build /cardboard-companion/cardboard-companion .

CMD ["supervisord", "-c", "./supervisord.conf"]
