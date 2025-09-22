### Frontent
FROM node:latest AS frontend-build

WORKDIR /code

COPY ui/package.json ui/package-lock.json ./
RUN npm install

COPY ui .

RUN npm run build

### Backend
FROM golang:latest AS backend-build

WORKDIR /cardboard-companion

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build

### Deployment
FROM node:latest

RUN apt-get update && apt-get install -y supervisor

ENV GIN_MODE=release
WORKDIR /code

COPY supervisord.conf .
COPY --from=frontend-build /code .
COPY --from=backend-build /cardboard-companion/cardboard-companion .

CMD ["supervisord", "-c", "./supervisord.conf"]
