begin;
--创建更新时间戳trigger
CREATE OR REPLACE FUNCTION func_create_set_timestamp()
    RETURNS TRIGGER AS
$$
BEGIN
    NEW.created_at = NOW();
    new.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION func_update_set_timestamp()
    RETURNS TRIGGER AS
$$
BEGIN
    new.created_at = old.created_at;
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;
--创建用户
drop table if exists "user";
create table if not exists "user"
(
    id         bigserial primary key,
    created_at timestamptz  not null,
    updated_at timestamptz  not null,
    deleted_at timestamptz  null,
    uuid       uuid,
    phone      varchar(20)  not null,
    username   varchar(32)  not null,
    password   varchar(256) not null,
    nickname   varchar(64)  not null,
    avatar     varchar(256) null,
    email      varchar(256) null,
    enabled    bool default false,
    system_id  bigint       not null
);
create index if not exists uuid_index on "user" (uuid);
create unique index if not exists phone_unique on "user" (phone);
create unique index if not exists username_unique on "user" (username);
comment on table "user" is '用户表';
comment on column "user".id is '主键bigint';
comment on column "user".created_at is '创建时间戳';
comment on column "user".updated_at is '更新时间戳';
comment on column "user".deleted_at is '删除时间戳';
comment on column "user".uuid is '用户uuid';
comment on column "user".phone is '手机号';
comment on column "user".username is '用户名';
comment on column "user".password is '密码';
comment on column "user".nickname is '昵称';
comment on column "user".avatar is '头像';
comment on column "user".email is '邮箱地址';
comment on column "user".enabled is '是否启用';


--创建请求记录表
drop table if exists "request_record";
create table if not exists "request_record"
(
    id            bigserial primary key,
    created_at    timestamptz not null,
    updated_at    timestamptz not null,
    deleted_at    timestamptz null,
    ip            inet,
    method        varchar(16),
    path          varchar(256),
    status        int,
    latency       interval second(6),
    agent         varchar(128),
    error_message varchar(256),
    body          varchar(512),
    resp          varchar(512),
    user_id       bigint
);
comment on table "request_record" is '操作记录表';
comment on column "request_record".id is '主键bigint';
comment on column "request_record".created_at is '创建时间戳';
comment on column "request_record".updated_at is '更新时间戳';
comment on column "request_record".deleted_at is '删除时间戳';

CREATE TRIGGER trigger_create_set
    BEFORE INSERT
    ON "user"
    FOR EACH ROW
EXECUTE PROCEDURE func_create_set_timestamp();

CREATE TRIGGER trigger_update_set
    BEFORE UPDATE
    ON "user"
    FOR EACH ROW
EXECUTE PROCEDURE func_update_set_timestamp();

CREATE TRIGGER trigger_create_set
    BEFORE INSERT
    ON "request_record"
    FOR EACH ROW
EXECUTE PROCEDURE func_create_set_timestamp();

CREATE TRIGGER trigger_update_set
    BEFORE UPDATE
    ON "request_record"
    FOR EACH ROW
EXECUTE PROCEDURE func_update_set_timestamp();
commit;