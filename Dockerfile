FROM golang:1.24-alpine as builder
WORKDIR /app

ARG APP_NAME
ENV APP_NAME ${APP_NAME}
ARG APP_VERSION
ENV APP_VERSION ${APP_VERSION}
ARG BUILD_TIME
ENV BUILD_TIME ${BUILD_TIME}
ARG GIT_REVISION
ENV GIT_REVISION ${GIT_REVISION}

COPY . .
RUN GO111MODULE=on go mod download \
  && go build -ldflags "-X github.com/isayme/websockify-go/websockify.Name=${APP_NAME} \
  -X github.com/isayme/websockify-go/websockify.Version=${APP_VERSION} \
  -X github.com/isayme/websockify-go/websockify.BuildTime=${BUILD_TIME} \
  -X github.com/isayme/websockify-go/websockify.GitRevision=${GIT_REVISION}" \
  -o ./dist/websockify main.go

FROM alpine
WORKDIR /app

ARG APP_NAME
ENV APP_NAME ${APP_NAME}
ARG APP_VERSION
ENV APP_VERSION ${APP_VERSION}

COPY --from=builder /app/dist/websockify ./

CMD ["/app/websockify", "server"]
