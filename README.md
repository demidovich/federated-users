# Federated Database Seeder

### Старт приложения и генерация базы

В генераторе захардкожено 100 федераций.

```shell
make up
go run ./cmd/main.go seed:make --rows 500000 --min-attrs 10 --max-attrs 100

make migrate
make psql
```

Загрузка данных

```sql
drop index if exists person_federation_uuid_idx;
drop index if exists person_attrs_idx;

copy person(uuid, federation_uuid, attrs) from '/seed/person.csv' delimiter ',' csv;

create index person_federation_uuid_idx on person (federation_uuid);
create index person_attrs_idx on person using gin (attrs jsonb_path_value_ops);
```

### SQL запросы

Получение пользователей федерации с возрастом от 30 до 32 лет.

```sql
-- Сначала получаем id какой-нибудь федерации
select federation_uuid from person limit 10;

-- Поиск пользователей которым от 30 до 32 лет
explain analyze select
    uuid,
    attrs->>'Age'       as Age,
    attrs->>'FirstName' as first_name,
    attrs->>'LastName'  as last_name
from person
where
    attrs @> '{"FederationUuid":"31225940-a755-48a9-812c-c5492fbffadd"}' and
    attrs#>'{Age}' >= to_jsonb(30::text) and
    attrs#>'{Age}' <= to_jsonb(32::text)
limit 5;
```

План "непрогретого" в кэше базы запроса.

```
federated_db=# explain analyze select
    uuid,
    attrs->>'Address',
    attrs->>'FirstName' as first_name,
    attrs->>'LastName' as LastName
from person                                                                                
where                                                                                       
    attrs @> '{"FederationUuid":"31225940-a755-48a9-812c-c5492fbffadd"}' and
    attrs @> '{"Age": "50"}'
limit 10;
                                                                    QUERY PLAN                                                                    
--------------------------------------------------------------------------------------------------------------------------------------------------
 Limit  (cost=84.00..88.02 rows=1 width=112) (actual time=76.399..76.442 rows=10 loops=1)
   ->  Bitmap Heap Scan on person  (cost=84.00..88.02 rows=1 width=112) (actual time=76.398..76.440 rows=10 loops=1)
         Recheck Cond: ((attrs @> '{"FederationUuid": "31225940-a755-48a9-812c-c5492fbffadd"}'::jsonb) AND (attrs @> '{"Age": "50"}'::jsonb))
         Heap Blocks: exact=10
         ->  Bitmap Index Scan on person_attrs_idx  (cost=0.00..84.00 rows=1 width=0) (actual time=76.360..76.360 rows=151 loops=1)
               Index Cond: ((attrs @> '{"FederationUuid": "31225940-a755-48a9-812c-c5492fbffadd"}'::jsonb) AND (attrs @> '{"Age": "50"}'::jsonb))
 Planning Time: 0.227 ms
 Execution Time: 76.460 ms
(8 rows)
```

Для этой комбинации в базе есть 116 значений

```
federated_db=# select                
    count(uuid)
from person
where                                 
    attrs @> '{"FederationUuid":"31225940-a755-48a9-812c-c5492fbffadd"}' and
    attrs @> '{"Age": "50"}';
 count 
-------
   151
```
