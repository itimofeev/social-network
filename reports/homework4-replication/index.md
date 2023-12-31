# Домашнее задание 2. Разделение монолита на сервисы

## Цель
Цель:
В результате выполнения ДЗ вы настроите синхронную репликацию, протестируете ее влияние на производительность системы и убедитесь, что теперь вы не теряете транзакции в случае аварии.
В данном задании тренируются навыки:
- обеспечение отказоустойчивости проекта;
- администрирование PostgreSQL;
- настройка репликации;
- проведение нагрузочных тестов.

##Описание/Пошаговая инструкция выполнения домашнего задания:
1. Настраиваем асинхронную репликацию.
2. Выбираем 2 запроса на чтение (/user/get/{id} и /user/search из спецификации) и переносим их на чтение со слейва.
3. Делаем нагрузочный тест по методам (/user/get/{id} и /user/search из спецификации), которые перевели на слейв до и после репликации. Замеряем нагрузку мастера (CPU, la, disc usage, memory usage).

4. Настроить 2 слейва и 1 мастер.
5. Включить и настроить синхронную репликацию.
6. Создать нагрузку на запись в любую тестовую таблицу. На стороне, которой нагружаем считать, сколько строк мы успешно записали.
7. С помощью kill -9 убиваем мастер PosgtreSQL.
8. Заканчиваем нагрузку на запись.
9. Выбираем самый свежий слейв. Промоутим его до мастера. Переключаем на него второй слейв.
10. Проверяем, есть ли потери транзакций.
 
Результатом сдачи ДЗ будет в виде исходного кода на github и отчета в текстовом виде, где вы отразите как вы выполняли каждый из пунктов.

Guide: https://github.com/OtusTeam/highload/blob/master/lessons/02/05/live/guide.md

## Критерии оценки:
Оценка происходит по принципу зачет/незачет.
Требования:
- В отчете корректно описано, как настроена репликация.
- 2 запроса переведено на чтение со слейва.
- Нагрузочное тестирование показало, что нагрузка перешла на слейв.
- В отчете описано как включить синхнронную и асинхронную репликацию
- Проведен эксперимент по потере и непотере транзакций при аварийной остановке master.

## Отчёт 1. Перенос читающей нагрузки на replica-у

Подготовка: запускаем два постгреса через docker compose `db` и `db-async` и настраиваем асинхронную репликацию.
<details>
  <summary>Раскрыть</summary>

    1. Создаем сеть, запоминаем адрес
    ```bash
    docker network inspect social-network_scnet | grep Subnet # Запомнить маску сети
    # "Subnet": "172.26.0.0/16",
    ```
    2. Меняем postgresql.conf на мастере
    ```
    ssl = off
    wal_level = replica
    max_wal_senders = 4 # expected slave num
    ```
    
    3. Подключаемся к мастеру и создаем пользователя для репликации
    ```bash
    docker exec -it social-network-db-1 psql -U admin social-network
    create role replicator with login replication password 'pass';
    exit
    ```
    
    4. Добавляем запись в pgmaster/pg_hba.conf с subnet с первого шага
    ```
    host    replication     replicator       172.26.0.0/16          md5
    ```
    
    5. Перезапустим мастер
    ```bash
    docker compose restart db
    ```
    
    6. Сделаем бэкап для реплик
    ```bash
    docker exec -it social-network-db-1 bash
    mkdir /pgasync
    pg_basebackup -h db -D /pgasync -U replicator -v -P --wal-method=stream
    exit
    ```
    
    7. Копируем директорию себе
    ```bash
    docker cp social-network-db-1:/pgasync tools/data/volumes/pgasync/
    ```
    
    8. Создадим файл, чтобы реплика узнала, что она реплика
    ```bash
    touch tools/data/volumes/pgasync/standby.signal
    ```
    
    9. Меняем postgresql.conf на реплике pgasync
    ```
    primary_conninfo = 'host=db port=5432 user=replicator password=pass application_name=pgasync'
    ```
    
    10. Запускаем реплику pgasync
    ```bash
    docker compose up db-async
    ```
    
    11. Убеждаемся что реплика работает в асинхронном режиме на pgmaster
    ```bash
    docker exec -it social-network-db-1 psql -U admin social-network
    select application_name, sync_state from pg_stat_replication;
    exit;
    ```
    Получаем:
    ```
    social-network=# select application_name, sync_state from pg_stat_replication;
    application_name | sync_state
    ------------------+------------
    pgasync          | async
    (1 row)
    ```
</details>

Далее запускаем приложение, в котором вся нагрузка идёт на мастер `db` и проводим нагрузочный тест на ручку поиска пользователей.
Видим, что CPU нагрузка именно на мастере:
```
➜  ~ docker stats  social-network-db-1 social-network-db-async-1
CONTAINER ID   NAME                        CPU %     MEM USAGE / LIMIT     MEM %     NET I/O          BLOCK I/O        PIDS
40ea7d8e2f2d   social-network-db-1         7.00%     180.6MiB / 7.668GiB   2.30%     173MB / 525MB    221kB / 47.7MB   18
e8531f7bd5ac   social-network-db-async-1   0.12%     159.3MiB / 7.668GiB   2.03%     496MB / 3.46MB   897kB / 4.1kB    7
```

Вносим изменения в приложение, чтобы читающая нагрузка (search users) выполнялась на реплике и повторяем замеры. Видим, что теперь CPU и память нагружены сильнее на ней:
```
➜  ~ docker stats  social-network-db-1 social-network-db-async-1
CONTAINER ID   NAME                        CPU %     MEM USAGE / LIMIT     MEM %     NET I/O          BLOCK I/O        PIDS
40ea7d8e2f2d   social-network-db-1         0.04%     163.6MiB / 7.668GiB   2.08%     174MB / 589MB    221kB / 47.7MB   9
e8531f7bd5ac   social-network-db-async-1   7.21%     175.6MiB / 7.668GiB   2.24%     496MB / 20.9MB   897kB / 4.1kB    16
```










## Отчёт 2. Возможная потеря транзакций в асинхронной реплике и отсутствие потерь в синхронной
Запустим docker-compose.replication.yml, в котором запускается pgmaster и две реплики pgsync и pgasync.
В команде запуска мастера указаны настройки:
```
-c synchronous_commit=on
-c synchronous_standby_names='pgsync'
```
которые говорят, что pgsync реплика является синхронной.

Запросом на мастер проверяем, что реплики имеют корректные настройки:
```
social-network=# select application_name, sync_state from pg_stat_replication;
application_name | sync_state
------------------+------------
pgasync          | async
pgsync           | sync
(2 rows)
```

Дальше запускаем пишущую нагрузку в табличку users. Через некоторое время убиваем pgmaster командой `docker kill --signal=9 social-network-pgmaster-1`.

Проверяем количество записей в табличке users `select count(1) from users;` на pgsync:
2684

На pgasync: 2684.

Эксперимент показал, что количество вставленных записей одинаковое в синхронной и асинхронной реплике, хотя мы ожидали, что в асинхронной будут потери.

Оказалось, что процесс перекачки WAL в реплики работает очень быстро! И в режиме, когда все контейнеры запущены на одной машине и благодаря этому имеют супер быструю сеть, воспроизвести ситуацию с потерей транзакций очень сложно.

Поэтому пришлось прибегнуть к утилите tc для выставления сетевых задержек контейнеру pgasync.
```bash
docker exec --user=root social-network-pgasync-1 tc qdisc add dev eth0 root netem delay 200ms
```
Повторив эксперимент можно видеть, что в реплике async записей значительно меньше, т.к. из-за задержек репликация не успевала переносить изменения на реплику.
Если же выставить задержки на контейнере pgsync, то потерей транзакций по сравнению с мастером не будет, просто значительно снизится пропускная способность pgmaster из-за ожидания применний изменний на синхронной реплике.

Это учебный эксперимент и в реальной сети задержки могут быть выше. Так же может просто не повести и мастер взорвётся, не успев передать изменения реплике, что так же приведёт к потере транзакций.
