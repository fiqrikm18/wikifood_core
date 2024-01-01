create table if not exists users
(
    id         uuid not null primary key DEFAULT uuid_generate_v4(),
    first_name varchar,
    last_name  varchar,
    phone      varchar,
    email      varchar,
    password   varchar,
    provider   int
);

ALTER TABLE users ADD CONSTRAINT fk_user_provider FOREIGN KEY (provider) REFERENCES providers (id);