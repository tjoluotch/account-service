# Go version
FROM golang:1.13 AS build-env
RUN mkdir /service

ENV USER=trevorjo
ENV UID=10001

# create a sytstem group dev with no password, no home directory set, and no shell so prevents the user form
# being a login account and reduces the attack vector
RUN adduser \
--disabled-password \
--gecos "" \
--home "/nonexistent" \
--shell "/sbin/nologin" \
--no-create-home \
--uid "${UID}" \
${USER}

WORKDIR /service
COPY . /service
# change ownership of all /service content to created user
RUN chown -R trevorjo /service
RUN go mod download && \
go mod verify && \
CGO_ENABLED=0 go build -o app -mod vendor -trimpath
USER trevorjo

FROM scratch AS run-env
WORKDIR /build
COPY --from=build-env /service/app /build/
ENTRYPOINT ["/build/app"]

