# Checklist — Глава 01

Перед переходом к главе 02 убедись:

- [ ] Понимаю разницу TCP и HTTP
- [ ] Могу прочитать HTTP request/response в `curl -v`
- [ ] Знаю структуру `http.Request` и `http.ResponseWriter`
- [ ] Умею установить status code и Content-Type
- [ ] Выполнил все задачи ⭐ и ⭐⭐ (минимум задачи 1–6)
- [ ] Код в `workspace/ch01/` запускается без ошибок
- [ ] Сделал git commit

## Самопроверка

1. Что произойдёт если вызвать `WriteHeader` дважды?
2. Какой status code по умолчанию если не вызвать `WriteHeader`?
3. Чем GET отличается от POST на уровне HTTP?
4. Что такое idempotent method?

<details>
<summary>Ответы</summary>

1. Второй вызов игнорируется (panic в некоторых middleware)
2. 200 OK
3. GET не имеет body (обычно), safe, для получения данных. POST может иметь body, создаёт/изменяет
4. Повторный запрос даёт тот же эффект (GET, PUT, DELETE — idempotent; POST — нет)

</details>
