if exists (select 1 from information_schema.tables where table_name = 'test_gomssqlrunner')
drop table test_gomssqlrunner

create table test_gomssqlrunner
(
    testfield1 varchar(200)
)