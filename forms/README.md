# Формы слова

Пакет *forms* используется для лемматизации слов на базе словаря всех известных форм.

## Файл словаря

Словарь форм хранится в файле *data.txt*, который подгружается при помощи *go:embed*.
Файл является текстовым и каждая строка отвечает за одно слово. Формы в каждой строке
разделяются пробельным символами. Самая первая форма в строке является базовой для слова.
Пример словаря ниже:

```plaintext
тест теста тесту тестом тесте
тесты тестов тестам тестами тестах
тесто теста тесту тестом тесте
```

Исходя из этих данных для формы  *тесту*, будет возвращена базовая форма
*тест* и *тесто*, для *тестов* - *тесты*.

## API

### Get

```golang
func Get(src string) ([]string, bool)
```

Функция возвращает массив базовых форм и *true* для исходной формы, если базовые формы определены,
либо *nil* и *false*, если нет данных.

### Each

```golang
func Each(src string, fn EachFunc)
```

Функция *Each* принимает на вход исходную форму слова и callback-функцию, в которую сначала передает
текущую форму слова, а потом все известные базовые. Если callback возвращает *false*, то процесс
сразу перерывается.

callback-функция имеет следующий формат:

```golang
type EachFunc func(string) bool
```