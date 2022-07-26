BEGIN;
drop table if exists "request_record";
drop table if exists "user";
drop function if exists "func_update_set_timestamp";
drop function if exists "func_create_set_timestamp";
COMMIT;