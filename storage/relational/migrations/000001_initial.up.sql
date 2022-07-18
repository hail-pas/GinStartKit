BEGIN;
--创建用户
drop table if exists "user";
create table if not exists "user"
(
    id         bigint primary key,
    created_at timestamptz not null,
    updated_at timestamptz not null,
    deleted_at timestamptz null
);
comment on table "user" is '用户表';
comment on column "user".id is '主键bigint';
comment on column "user".created_at is '创建时间戳';
comment on column "user".updated_at is '更新时间戳';
comment on column "user".deleted_at is '删除时间戳';

--创建请求记录表
drop table if exists "operation_record";
create table if not exists "operation_record"
(


);
comment on table "operation_record" is '操作记录表';
comment on column "operation_record".id is '主键bigint';
comment on column "operation_record".created_at is '创建时间戳';
comment on column "operation_record".updated_at is '更新时间戳';
comment on column "operation_record".deleted_at is '删除时间戳';
COMMIT;

--创建更新时间戳trigger
CREATE OR REPLACE FUNCTION func_set_create_timestamp()
    RETURNS TRIGGER AS
$$
BEGIN
    NEW.created_at = NOW();
    new.updated_at = now();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION func_set_update_timestamp()
    RETURNS TRIGGER AS
$$
BEGIN
    new.created_at = old.created_at;
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trigger_set_update_timestamp
    BEFORE INSERT
    ON "user"
    FOR EACH ROW
EXECUTE PROCEDURE func_set_create_timestamp();

CREATE TRIGGER trigger_set_update_timestamp
    BEFORE UPDATE
    ON "user"
    FOR EACH ROW
EXECUTE PROCEDURE func_set_update_timestamp();

CREATE TRIGGER trigger_set_update_timestamp
    BEFORE INSERT
    ON "operation_record"
    FOR EACH ROW
EXECUTE PROCEDURE func_set_create_timestamp();

CREATE TRIGGER trigger_set_update_timestamp
    BEFORE UPDATE
    ON "operation_record"
    FOR EACH ROW
EXECUTE PROCEDURE func_set_update_timestamp();
