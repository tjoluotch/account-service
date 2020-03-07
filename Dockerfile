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
#RUN groupadd -r dev && \
#useradd -r -s /bin/false -g dev trevorjo sudo
WORKDIR /service
COPY . /service
# change ownership of all /service content to created user
RUN chown -R trevorjo /service
#RUN echo "trevorjo ALL=(root) NOPASSWD:ALL" > /etc/sudoers.d/user && \
#chmod 0440 /etc/sudoers.d/user
# GOCACHE disable as get a permission denied error due to running as non root user
RUN go mod download && \
go mod verify && \
CGO_ENABLED=0 GOOS=linux go build -o app -mod vendor -trimpath
USER trevorjo

FROM scratch AS run-env
WORKDIR /build
COPY --from=build-env /service/app /build/
ENTRYPOINT ["/build/app"]

