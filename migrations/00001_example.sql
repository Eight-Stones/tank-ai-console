-- +goose Up
create table if not exists example
(
    id   SERIAL PRIMARY KEY,
    code varchar,
    name varchar,
    meta varchar
);

insert into example (code, name, meta) values ('test-code-1','test-name-1','test-meta-1');
insert into example (code, name, meta) values ('test-code-2','test-name-2','test-meta-2');
insert into example (code, name, meta) values ('test-code-3','test-name-3','test-meta-3');

-- +goose Down