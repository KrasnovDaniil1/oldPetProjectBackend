<h1>PDFCPU-RESTAPI</h1>

    <span>Проект создан для работы с документами формата pdf. Использует библиотеку <a href="https://pdfcpu.io/" target="\_blank">PdfCpu</a>. Также для удобной работы написан frontend - pdfcpu-frontend </span>

    <h1>Как запустить проект:</h1>
    <ul>
        <li>
            <span>Запустить в докере:</span>
            <ul>
                <li>docker build -t go-docker-image .</li>
                <li>docker run go-docker-image</li>
            </ul>
        </li>
        <li>
            <span>Создать .env файл с параметрами:</span>
            <ul>
                <li>ENDPOINT - ссылка на сервер - storage.yandexcloud.net</li>
                <li>ACCESS_KEY_ID - id ключа / логин</li>
                <li>SECRET_ACCESS_KEY - ключ доступа / пароль</li>
                <li>REGION - регион - ru-central1</li>
                <li>BUCKET_NAME - название бакета - bucketpdf</li>
                <li>PORT - порт для запуска, по умалчанию 3000 - 8080</li>
            </ul>
        </li>
    </ul>
    <h1>Как использовать.</h1>
    <h2>Stamp.</h2>
    <h4>POST /addWatermarks -- создаёт stamp/watermark</h4>
    <ul>
        <li>inFile -- file/pdf файл для изменения.</li>
        <li>mode -- string тип печати text/pdf/img.</li>
        <li>fileMode -- file/pdf/img файл в качестве stamp/watermark, нужен если mode указан как pdf/img.</li>
        <li>pagePdfMode -- integer какую страницу использовать в качестве stamp/watermark, нужен если mode указан как pdf.</li>
        <li>onTop -- bool что использовать для печати stamp/watermark.</li>
        <li>update -- bool обновить печать в файле.</li>
        <li>selectPage -- integer/string на каких страницах поставить печать. all/1,2</li>
        <li>text -- string текст для печати, нужен если mode указан как test.</li>
        <li>description -- string конфигурации для печати. Подробнее смотрите на сайте библиотеки PdfCpu.</li>
    </ul>
    <h4>POST /removeWatermarks -- удаляет stamp/watermark.</h4>
    <ul>
        <li>inFile -- file/pdf файл для изменения.</li>
        <li>selectPage -- integer/string на каких страницах удалить печать.</li>
    </ul>
    <h4>POST /collect -- изменяет расположение страниц/дублирует</h4>
    <ul>
        <li>inFile -- file/pdf файл для изменения.</li>
        <li>selectPage -- integer/string расположение страниц. - 1,1,1,2-l-1</li>
    </ul>
    <h4>POST /rotate -- переварачивает страницы.</h4>
    <ul>
        <li>inFile -- file/pdf файл для изменения.</li>
        <li>selectPage -- integer/string какие страницы перевернуть</li>
        <li>rotate -- integer на сколько градусов развернуть.</li>
    </ul>
    <h4>POST /trim -- режет страницы.</h4>
    <ul>
        <li>inFile -- file/pdf файл для изменения.</li>
        <li>selectPage -- integer/string какие страницы изменить</li>
    </ul>
    <h4>POST /crop -- обрезает страницы.</h4>
    <ul>
        <li>inFile -- file/pdf файл для изменения.</li>
        <li>selectPage -- integer/string какие страницы изменить</li>
        <li>description -- string коррдинаты для обрезания - [0 0 200 200]</li>
    </ul>
    <h4>POST /merge -- объединяет страницы.</h4>
    <ul>
        <li>inFiles -- file/pdf файлы для объединения, минимум 2.</li>
    </ul>
    <h4>POST /optimize -- оптимизирует файл.</h4>
    <ul>
        <li>inFile -- file/pdf файл для оптимизации.</li>
    </ul>
    <h4>POST /split -- режет файл.</h4>
    <ul>
        <li>inFile -- file/pdf файл для оптимизации.</li>
        <li>span -- int шаг для разделения страницы</li>
    </ul>
