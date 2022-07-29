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
    enabled    bool default false
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
    id         bigserial primary key,
    created_at timestamptz not null,
    updated_at timestamptz not null,
    deleted_at timestamptz null,
    client_ip  inet,
    method     varchar(16),
    path       varchar(256),
    status_code int,
    latency    interval second(6),
    agent      varchar(128),
    query      json,
    form_data  json,
    body       json,
    resp       json,
    user_id    bigint
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
drop table if exists system;
create table system
(
    id          bigserial primary key,
    created_at  timestamptz  not null,
    updated_at  timestamptz  not null,
    deleted_at  timestamptz  null,
    code        varchar(32)  not null,
    label       varchar(64)  not null,
    description varchar(255) null
);
create unique index if not exists code_unique on system (code);
comment on table "system" is '系统表';
comment on column "system".id is '主键bigint';
comment on column "system".created_at is '创建时间戳';
comment on column "system".updated_at is '更新时间戳';
comment on column "system".deleted_at is '删除时间戳';
comment on column "system".code is '唯一标识';
comment on column "system".label is 'Label值';
comment on column "system".description is '描述';
insert into system (created_at, updated_at, code, label, description)
values (NOW(), NOW(), 'jacCommercial', '江淮商务车', null),
       (NOW(), NOW(), 'jacCivilian', '江淮民用车', null);

CREATE TRIGGER trigger_create_set
    BEFORE INSERT
    ON "system"
    FOR EACH ROW
EXECUTE PROCEDURE func_create_set_timestamp();

CREATE TRIGGER trigger_update_set
    BEFORE UPDATE
    ON "system"
    FOR EACH ROW
EXECUTE PROCEDURE func_update_set_timestamp();

drop table if exists role;
create table role
(
    id          bigserial primary key,
    created_at  timestamptz  not null,
    updated_at  timestamptz  not null,
    deleted_at  timestamptz  null,
    code        varchar(32)  not null,
    label       varchar(64)  not null,
    description varchar(255) null,
    system_id   bigint       not null
);
create unique index if not exists code_unique on role (code);
create index if not exists system_id_index on role (system_id);
comment on table "role" is '角色表';
comment on column "role".id is '主键bigint';
comment on column "role".created_at is '创建时间戳';
comment on column "role".updated_at is '更新时间戳';
comment on column "role".deleted_at is '删除时间戳';
comment on column "role".code is '唯一标识';
comment on column "role".label is 'Label值';
comment on column "role".description is '描述';
CREATE TRIGGER trigger_create_set
    BEFORE INSERT
    ON "role"
    FOR EACH ROW
EXECUTE PROCEDURE func_create_set_timestamp();

CREATE TRIGGER trigger_update_set
    BEFORE UPDATE
    ON "role"
    FOR EACH ROW
EXECUTE PROCEDURE func_update_set_timestamp();

CREATE TYPE system_resource_type AS ENUM ('menu', 'button', 'api');
drop table if exists system_resource;
create table system_resource
(
    id               bigserial primary key,
    created_at       timestamptz          not null,
    updated_at       timestamptz          not null,
    deleted_at       timestamptz          null,
    code             varchar(32)          not null,
    label            varchar(64)          not null,
    description      varchar(255)         null,
    parent_id        bigint               null,
    reference_to_id  bigint               null,
    front_route_path varchar(128)         null,
    icon_path        varchar(255)         null,
    type             system_resource_type not null,
    order_num        int  default 0,
    enabled          bool default false
);
create unique index if not exists code_unique on system_resource (code, parent_id);
comment on table "system_resource" is '系统资源表';
comment on column "system_resource".id is '主键bigint';
comment on column "system_resource".created_at is '创建时间戳';
comment on column "system_resource".updated_at is '更新时间戳';
comment on column "system_resource".deleted_at is '删除时间戳';
comment on column "system_resource".code is '唯一标识';
comment on column "system_resource".label is 'Label值';
comment on column "system_resource".description is '描述';
CREATE TRIGGER trigger_create_set
    BEFORE INSERT
    ON "system_resource"
    FOR EACH ROW
EXECUTE PROCEDURE func_create_set_timestamp();

CREATE TRIGGER trigger_update_set
    BEFORE UPDATE
    ON "system_resource"
    FOR EACH ROW
EXECUTE PROCEDURE func_update_set_timestamp();

drop table if exists permission;
create table permission
(
    id          bigserial primary key,
    created_at  timestamptz  not null,
    updated_at  timestamptz  not null,
    deleted_at  timestamptz  null,
    code        varchar(32)  not null,
    label       varchar(64)  not null,
    description varchar(255) null
);
create unique index if not exists code_unique on "permission" (code);
comment on table "permission" is '权限表';
comment on column "permission".id is '主键bigint';
comment on column "permission".created_at is '创建时间戳';
comment on column "permission".updated_at is '更新时间戳';
comment on column "permission".deleted_at is '删除时间戳';
comment on column "permission".code is '唯一标识';
comment on column "permission".label is 'Label值';
comment on column "permission".description is '描述';

CREATE TRIGGER trigger_create_set
    BEFORE INSERT
    ON "permission"
    FOR EACH ROW
EXECUTE PROCEDURE func_create_set_timestamp();

CREATE TRIGGER trigger_update_set
    BEFORE UPDATE
    ON "permission"
    FOR EACH ROW
EXECUTE PROCEDURE func_update_set_timestamp();

drop table if exists "system_with_system_resource";
create table system_with_system_resource
(
    system_id          bigint not null,
    system_resource_id bigint not null
);
create unique index if not exists unique_index on system_with_system_resource (system_id, system_resource_id);
comment on table "system_with_system_resource" is '系统和系统资源多对多关系';
comment on column "system_with_system_resource".system_id is '系统ID';
comment on column "system_with_system_resource".system_resource_id is '系统资源ID';

drop table if exists "role_with_system_resource";
create table role_with_system_resource
(
    role_id            bigint not null,
    system_resource_id bigint not null
);
create unique index if not exists unique_index on role_with_system_resource (role_id, system_resource_id);
comment on table "role_with_system_resource" is '角色和系统资源多对多关系';
comment on column "role_with_system_resource".role_id is '角色ID';
comment on column "role_with_system_resource".system_resource_id is '系统资源ID';

drop table if exists "system_resource_with_permission";
create table system_resource_with_permission
(
    permission_id      bigint not null,
    system_resource_id bigint not null
);
create unique index if not exists unique_index on system_resource_with_permission (system_resource_id, permission_id);
comment on table "system_resource_with_permission" is '权限和系统资源多对多关系';
comment on column "system_resource_with_permission".permission_id is '权限ID';
comment on column "system_resource_with_permission".system_resource_id is '系统资源ID';

drop table if exists "system_with_user";
create table system_with_user
(
    system_id bigint not null,
    user_id   bigint not null
);
create unique index if not exists unique_index on system_with_user (system_id, user_id);
comment on table "system_with_user" is '系统和用户多对多关系';
comment on column "system_with_user".system_id is '系统ID';
comment on column "system_with_user".user_id is '用户ID';
commit;