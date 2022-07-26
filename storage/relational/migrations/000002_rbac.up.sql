begin;
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

drop table if exists system_resource;
create table system_resource
(
    id               bigserial primary key,
    created_at       timestamptz  not null,
    updated_at       timestamptz  not null,
    deleted_at       timestamptz  null,
    code             varchar(32)  not null,
    label            varchar(64)  not null,
    description      varchar(255) null,
    parent_id        bigint       null,
    reference_to_id  bigint       null,
    front_route_path varchar(128) null,
    icon_path        varchar(255) null,
    type             char(8)      not null,
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
commit;