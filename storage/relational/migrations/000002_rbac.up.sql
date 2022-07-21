begin;
drop table if exists role;
create table role
(
    id         bigserial primary key,
    created_at timestamptz not null,
    updated_at timestamptz not null,
    deleted_at timestamptz null
);

create table system_resource
(
    id         bigserial primary key,
    created_at timestamptz not null,
    updated_at timestamptz not null,
    deleted_at timestamptz null
);

create table permission
(
    id         bigserial primary key,
    created_at timestamptz not null,
    updated_at timestamptz not null,
    deleted_at timestamptz null
);

commit;