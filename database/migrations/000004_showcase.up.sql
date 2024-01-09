create table if not exists showcase
(
    id         uuid not null primary key DEFAULT uuid_generate_v4(),
    title varchar,
    body text,
    image varchar,
    sort int
);
