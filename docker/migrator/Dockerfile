FROM golang:1.23 as build

WORKDIR /source

COPY . .

RUN make build-migrator

FROM scratch

COPY --from=build /usr/share/zoneinfo /usr/share/zoneinfo

COPY --from=build source/bin/migrator .
COPY --from=build source/config/migrator.toml config/migrator.toml
COPY --from=build source/migrations migrations

CMD ["/migrator", "-config", "config/migrator.toml"]



