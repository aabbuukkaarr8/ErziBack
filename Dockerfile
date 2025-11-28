# docker/Dockerfile.postgres
FROM postgres:16.6-alpine3.20

# Копируем миграции (они уже пронумерованы: 000001_…, 000002_… и т.д.)
COPY db/migrations/*.up.sql /docker-entrypoint-initdb.d/

# Копируем insert.sql и даём ему имя так,
# чтобы он выполнялся после всех миграций (например, 999_insert.sql)
COPY db/insert/insert.sql /docker-entrypoint-initdb.d/999_insert.sql

# (По желанию) задаём рабочую директорию
WORKDIR /docker-entrypoint-initdb.d
