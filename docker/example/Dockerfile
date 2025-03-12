FROM golang:1.23 as build

WORKDIR /source
COPY . .
RUN make build-example

FROM scratch

EXPOSE 8080
COPY --from=build /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=build source/bin/example .
COPY --from=build /source/config/config.toml /config/config.toml

ENTRYPOINT ["/example"]


