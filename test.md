### Около БД

#### add commands

Добавить валюту руками

`gocconv currency add [Name] [Token]`

Получить инфу из интернета

`gocconv currency get [Token]`

Добавить обмен валют

`gocconv rate add [FromToken] [ToToken] [Rate]`

Получить курс из интернета

`gocconv rate get [FromToken] [ToToken]`

### Конвертация

#### convert commands

Конвертировать `ammount` валюты в другую

`gocconv convert [FromToken] [ToToken] [amount]`
