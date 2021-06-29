# dicts

Пакет *dicts* реализует логику по работе со словарями, которые в дальнейшем используется в
алгоритмах выделения ключевых понятий из текста.

## Хранение словарей

Каждая категория слов хранится в своем словаре. Имя файла словаря имеет следующий формат *dict_NAME.txt*
В дальнейшем при получении словаря для конкретного слова будет возвращено *NAME*.

Словари встраиваются в библиотеку при помощи *go:embed*, а загрузка производится на этапе инициализации
пакета *dicts*.

## Формат файла словаря

Файлы словарей имеют текстовый формат. В каждой строке файла определяется одно слово из словаря.
Пример файла словаря ниже:

```plaintext
слово1
слово2
слово3
...
словоN
```

## API

### Each

```golang
func Each(src string, fn EachFunc)
```

Функция *Each* последовательно передает в функцию *fn* все словари, в которые включено исходное слово,
до тех пор пока функция *fn* не вернет *false* или не закончатся словари. При этом функция *fn* должна соответствовать
следующему типу:

```golang
type EachFunc func(string) bool
```

### Get

```golang
func Get(src string) ([]string, bool)
```

Функция *Get* возвращает массив словарей, в которые включено слово, если такие имеются. Вторым параметров
возвращается флаг того, что слово включено хотя бы в один словарь.

### InDict

```golang
func InDict(dict string, name string) bool
```

Функция *InDict* проверяет попадает ли переданное слово *name* в словарь *dict*. В случае успешной проверки
возвращается *true*, в противном случае - *false*.

### Size

```golang
func Size() int
```

Функция *Size* возвращает общее число уникальных слов во всех словарях.