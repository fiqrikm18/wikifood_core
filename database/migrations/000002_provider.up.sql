create sequence providers_id_seq;
create table if not exists providers
(
    id   int not null primary key default nextval('providers_id_seq'),
    name varchar
);
ALTER SEQUENCE providers_id_seq OWNED BY providers.id;