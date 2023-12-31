# Домашнее задание 3. Разделение монолита на сервисы

## Цель
В результате выполнения ДЗ вы перенесете бизнес-домен монолитного приложения в отдельный сервис.
В данном задании тренируются навыки:
- декомпозиции предметной области;
- разделение монолитного приложения;
- работа с HTTP;
- работа с REST API и gRPC;

Реализовать функционал:
- Вынести систему диалогов в отдельный сервис.

## Требования
- Взаимодействия монолитного сервиса и сервиса чатов реализовать на Rest API или gRPC.
- Организовать сквозное логирование запросов.
- Предусмотреть то, что не все клиенты обновляют приложение быстро и кто-то может ходить через старое API.

## Отчёт
Сначала я реализовал два метода для отправки и получения сообщения диалога в основном микросервисе backend.
Затем согласно условию ДЗ вынес их в отдельный микросервис dialogs.
Для маршрутизации запросов я добавил компонент api-gateway, который реализуется с помощью traefik. 
Помимо маршрутизации запросов traefik ещё проводит авторизацию запросов и в приватные микросервисы отправляет специальный заголовок `X-Sc-User-Id`, которому микросервисы доверяют.
Так же traefik добавляет заголовок `X-Request-Id` в запросы к микросервисам. Те в свою очередь прокидывают этот заголовок, если обращаются к другим микросервисам. 
Таким образом в логах можно увидеть поле request_id, который будет одинаковый в рамках обработки запроса во всех микросервисах. Это позволяет реализовать сквозной логирование запросов.
Для поддержки обновления клиентов можно будет завести новый метод /api/v1/dialogs/{user_id}/sendV2, которым будут пользоваться новые клиенты. Старый же метод можно будет удалить, когда все клиенты перейдут на новую версию.
Сервис dialogs ходит в backend для проверкти того, что пользователь, которому мы хотим написать сообщение зарегистрирован.

Для локального запуска нужно:
1. скачать проект `git clone git@github.com:itimofeev/social-network.git`
2. запустить `docker compose --file docker-compose.yc.yml up -d --remove-orphans`
3. залогиниться под первый пользователем 
```bash
curl -X POST --location "http://127.0.0.1:80/api/v1/login" \
    -H "Content-Type: application/json" \
    -d '{
          "id": "a086c063-713e-4497-8a07-0b659a48eb41",
          "password": "123456"
        }'
```
В ответ будет получен токен в формате PASETO:
`v4.local.2w-1XbhJshuutel46NHJ1jl5BMyj364LzgVKlVLFNPOcV1YbcfAbCXUW9FYVHnt7jHjBLL-Gv3sVmqYvlRBP_WL-_zOFsYF2Lo8irJfavttnEYiSIT5bPTFsKDyfe2PjWXEZty3c_FU9hfmvmYSm_4rzMdwQmlk5frXZFTt31jFbRzasTi7ckTPLnWx7LrAdeUufahmABO7XOTD79VoHdoTrPgLmB-GL124xOSxe-jwjO7n93dEFXlcedp8mIHD2ZbeRE91_LYDU`
4. затем с этим токеном можно будет отправить сообщение второму пользователю
```
curl -X POST --location "http://127.0.0.1:80/api/v1/dialog/d06ff731-f291-4703-8010-a53c62be5d2b/send" \
    -H "Content-Type: application/json" \
    -H "Authorization: Bearer v4.local.2w-1XbhJshuutel46NHJ1jl5BMyj364LzgVKlVLFNPOcV1YbcfAbCXUW9FYVHnt7jHjBLL-Gv3sVmqYvlRBP_WL-_zOFsYF2Lo8irJfavttnEYiSIT5bPTFsKDyfe2PjWXEZty3c_FU9hfmvmYSm_4rzMdwQmlk5frXZFTt31jFbRzasTi7ckTPLnWx7LrAdeUufahmABO7XOTD79VoHdoTrPgLmB-GL124xOSxe-jwjO7n93dEFXlcedp8mIHD2ZbeRE91_LYDU" \
    -d '{
          "text": "hello, there!"
        }'
```
5. и затем можно будет получить список сообщений
```
curl -X GET --location "http://127.0.0.1:80/api/v1/dialog/d06ff731-f291-4703-8010-a53c62be5d2b/list" \
    -H "Content-Type: application/json" \
    -H "Authorization: Bearer v4.local.2w-1XbhJshuutel46NHJ1jl5BMyj364LzgVKlVLFNPOcV1YbcfAbCXUW9FYVHnt7jHjBLL-Gv3sVmqYvlRBP_WL-_zOFsYF2Lo8irJfavttnEYiSIT5bPTFsKDyfe2PjWXEZty3c_FU9hfmvmYSm_4rzMdwQmlk5frXZFTt31jFbRzasTi7ckTPLnWx7LrAdeUufahmABO7XOTD79VoHdoTrPgLmB-GL124xOSxe-jwjO7n93dEFXlcedp8mIHD2ZbeRE91_LYDU"
```
