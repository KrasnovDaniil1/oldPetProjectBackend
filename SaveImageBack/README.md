<h1>save-Image-Back</h1>
<span>Сервер для сохранения картинок по тэгам. Мой первый относительно крупный сервер. Возможно потом переделую.</span>

<h2>Быстрый запуск</h2>
<span>docker run -it --name some-postgres -e POSTGRES_PASSWORD=pass -p 5432:5432 -e POSTGRES_USER=user -e POSTGRES_DB=db postgres</span>
<span>go run ./main.go</span>
<h3>Вход в базу</h3>
<span>docker exec -it ID_КОНТЕЙНЕРА psql -U user -d db</span>

<h3>GET /cards - получить карточки</h3>
<span>--params:</span>
<ul>
<li>search - поиск карточки по тегам</li>
<li>offset - пропустить количестов карточек</li>
<li>limit - получить количество карточек</li>
</ul>

<h3>POST /cards - добавить карточку</h3>
<span>--form:</span>
<ul>
<li>image - загрузить картинка для карточки</li>
<li>tags - теги для карточки перечисленные в виде текста, через запятую, без пробела.</li>
</ul>

<h3>GET /tags - получить все тэги</h3>
<span>--params:</span>
<ul>
<li>search - поиск тега</li>
</ul>

<h3>DELETE /cards/:id - удалить карточку по id(получаем вместе с карточками)</h3>
