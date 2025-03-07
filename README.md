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

Получение пользователей федерации с возрастом 48 лет.

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

