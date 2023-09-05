## Устные вопросы

#### Какой самый эффективный способ конкатенации строк?
> **Ответ**:
В языке Golang конкатенацию строк можно выполнять несколькими способами, и эффективный способ зависит от конкретной ситуации. Вот несколько распространенных способов:

1. Использование оператора `+`: Простым способом конкатенации строк является использование оператора `+`. Однако внутри циклов или при частых конкатенациях большого количества строк это может быть неэффективным из-за необходимости создания нового буфера каждый раз при конкатенации:
```go
result := string1 + string2 + string3
```

2. Использование `strings.Join()`: Функция `strings.Join()` позволяет конкатенировать срез строк с использованием разделителя, что может быть эффективным для конкатенации большого количества строк:
```go
result := strings.Join([]string{string1, string2, string3}, "")
```

3. Использование `bytes.Buffer`: Если вам нужно часто изменять или конкатенировать строки в цикле, использование `bytes.Buffer` может быть эффективным. `bytes.Buffer` предоставляет буферизованную операцию записи, которая уменьшает количество выделений памяти при конкатенации строк:
```go
var buffer bytes.Buffer
buffer.WriteString(string1)
buffer.WriteString(string2)
buffer.WriteString(string3)
result := buffer.String()
```

4. Использование `strings.Builder`: В Go 1.10 и выше рекомендуется использовать `strings.Builder`, поскольку это более оптимизированный и потокобезопасный вариант `bytes.Buffer`:
```go
var builder strings.Builder
builder.WriteString(string1)
builder.WriteString(string2)
builder.WriteString(string3)
result := builder.String()
```

Обратите внимание, что эффективность конкатенации строк может зависеть от конкретного случая использования, размера строк и числа операций конкатенации. Рекомендуется проводить тесты производительности для выбора самого эффективного способа в вашем конкретном случае.

#### Что такое интерфейсы, как они применяются в Go?
> **Ответ**:
В Go интерфейс представляет собой набор методов, которые должны быть реализованы в типе, чтобы удовлетворять этому интерфейсу. Интерфейсы предоставляют абстракцию и позволяют программисту создавать обобщенный код, который может работать с различными типами данных без необходимости знать их конкретную реализацию.

Определение интерфейса происходит следующим образом:
```go
type ИмяИнтерфейса interface {
    Метод1()
    Метод2()
    // ...
}
```

Для того чтобы тип реализовал интерфейс, необходимо реализовать все методы, указанные в интерфейсе:
```go
type ТипСтруктуры struct {
    // поля структуры
}

func (s ТипСтруктуры) Метод1() {
    // реализация Метод1
}

func (s ТипСтруктуры) Метод2() {
    // реализация Метод2
}
```

После того, как тип реализовал все методы интерфейса, он автоматически становится удовлетворяющим этому интерфейсу. Затем код может использовать интерфейс вместо конкретного типа, и вызывать его методы без необходимости ссылаться на конкретные реализации.

Пример использования интерфейса:
```go
type Shape interface {
    Area() float64
}

type Rectangle struct {
    width  float64
    height float64
}

func (r Rectangle) Area() float64 {
    return r.width * r.height
}

type Circle struct {
    radius float64
}

func (c Circle) Area() float64 {
    return math.Pi * c.radius * c.radius
}

func main() {
    var shape Shape
    shape = Rectangle{width: 4, height: 5}
    fmt.Println("Площадь прямоугольника:", shape.Area())

    shape = Circle{radius: 3}
    fmt.Println("Площадь круга:", shape.Area())
}
```

В данном примере интерфейс `Shape` определяет метод `Area()`, который возвращает площадь фигуры. Структуры `Rectangle` и `Circle` реализуют этот метод, поэтому они могут быть присвоены переменной типа `Shape`. Это позволяет вызывать метод `Area()` на переменной `shape` без явного указания на конкретную реализацию.

####  Чем отличаются RWMutex от Mutex?
> **Ответ**:
RWMutex и Mutex - это два различных типа блокировок для синхронизации доступа к общим ресурсам в многопоточной среде в языке Go. Они имеют ряд отличий в функциональности и использовании:

1. Mutex:
   - Mutex является базовым типом блокировки в Go и предоставляет эксклюзивную блокировку (только для чтения или только для записи).
   - Только один поток может захватить блокировку в определенный момент времени, и другие потоки будут ожидать до освобождения блокировки.
   - Mutex выполняет роль простой двоичной блокировки для обеспечения согласованности доступа к ресурсу.

2. RWMutex:
   - RWMutex предоставляет блокировку для чтения и записи (readers-writers lock) и позволяет множественным потокам параллельно читать общий ресурс, но только одному потоку писать в него.
   - Несколько потоков могут одновременно вызывать метод RLock() для захвата блокировки на чтение, пока нет активных блокировок на запись.
   - Однако, если уже захвачена блокировка на запись (с помощью метода Lock()), операции на чтение ждут освобождения блокировки на запись.
   - RWMutex полезен в ситуациях, где часто выполняется чтение данных, но запись происходит реже.

Общим для Mutex и RWMutex является то, что обе блокировки должны быть захвачены и освобождены тем же потоком, чтобы избежать состояния гонки. Также стоит отметить, что RWMutex является немного более сложным в использовании и может иметь некоторые ограничения и накладные расходы по сравнению с Mutex, поэтому его следует использовать тогда, когда есть явная необходимость и польза от разделяемого параллельного чтения и ограниченного единичного записи.

####  Чем отличаются буферизированные и не буферизированные каналы?
> **Ответ**:
В языке Go каналы представляют собой механизм для обмена данными между горутинами. Каналы могут быть буферизированными и не буферизированными. Вот их основные отличия:

1. Небуферизированные каналы:
   - Небуферизированные каналы имеют емкость 0.
   - Запись в небуферизированный канал блокируется, пока данные не будут прочитаны из канала другой горутиной.
   - Чтение из небуферизированного канала блокируется, пока данные не будут записаны в канал другой горутиной.
   - Взаимодействие через небуферизированный канал вызывает синхронизацию между горутинами, горутины ожидают друг друга.

2. Буферизированные каналы:
   - Буферизированные каналы имеют емкость больше 0, которая указывается при их создании.
   - Буферизированный канал позволяет записывающей горутине отправить данные в канал и продолжить работу, если буфер не заполнен.
   - Чтение из буферизированного канала блокируется только в том случае, если буфер пуст. Горутина может продолжить выполнение, если есть достаточно данных в канале.
   - Буферизированные каналы позволяют асинхронное взаимодействие между горутинами.

Выбор между использованием буферизованных и небуферизованных каналов зависит от конкретной ситуации. Небуферизованные каналы часто используются для точной синхронизации между горутинами, когда горутины должны взаимодействовать близким образом. Буферизированные каналы удобны в случаях, когда требуется некоторая гибкость и возможность буферизации данных на короткое время.

Примеры создания небуферизованного и буферизованного каналов:
```go
// небуферизованный канал
ch := make(chan int)

// буферизованный канал с емкостью 10
ch := make(chan int, 10)
```
Обратите внимание, что размер буфера (второй аргумент в `make(chan int, bufferSize)`) определяет емкость канала. При отсутствии буфера (0) канал является небуферизованным.

####  Какой размер у структуры struct{}{}?
> **Ответ**:
В Go структура `struct{}{}` (пустая структура) является минималистичной структурой без полей. Такая структура не занимает никакое место в памяти. Это связано с тем, что в Go пустая структура не имеет полей и, следовательно, не требует выделения памяти для хранения данных.

Однако, даже если структура не имеет полей и не занимает место в памяти, она имеет размер отличный от нуля. В Go, каждая уникальная структура имеет свой собственный размер в памяти, независимо от ее содержимого. Минимальный размер структуры в Go составляет 0 байт.

Кроме того, размер пустой структуры может быть больше нуля, если он выровнен по размеру, необходимому для целевой архитектуры. Но даже в этом случае размер будет очень малым и будет зависеть от компилятора и платформы, на которой выполняется код.

В общем случае не стоит беспокоиться о размере пустой структуры, так как он обычно не имеет существенного значения.

####  Есть ли в Go перегрузка методов или операторов?
> **Ответ**:
В языке Go не поддерживается перегрузка методов или операторов, как в некоторых других языках программирования. Перегрузка возникает в ситуациях, когда существует несколько методов или операторов с одним и тем же именем, но с разными типами параметров или арностью.

Вместо перегрузки, в Go используется концепция набора интерфейсов и пустых идентификаторов (`interface{}`) для достижения подобного эффекта. Go поддерживает полиморфизм через интерфейсы, позволяющие вызывать методы на объекте, если он реализует определенный интерфейс, независимо от его конкретного типа.

Для операторов в Go также отсутствует перегрузка. Вместо этого, операторы для встроенных типов уже предопределены, и при необходимости можно определить свои пользовательские методы для собственных типов.

Отсутствие перегрузки методов и операторов в Go является частью его дизайна, который стремится к ясности и простоте языка. Однако можно использовать другие паттерны и техники (например, функциональные параметры) для эффективной работы с различными типами данных и достижения желаемой функциональности.

####  В какой последовательности будут выведены элементы map[int]int?
Пример:
```
m[0]=1
m[1]=124
m[2]=281
```
> **Ответ**: В случайной последовательности так как это m принадлежит типу map.
Поскольку map в Go является неупорядоченной коллекцией, порядок элементов в map не гарантирован и может быть разным при каждом выполнении программы. Поэтому порядок вывода элементов map[int]int может быть разным при каждом запуске программы.

В вашем примере, возможны следующие варианты вывода элементов map:

1. `{0: 1, 1: 124, 2: 281}`
2. `{0: 1, 2: 281, 1: 124}`
3. `{1: 124, 0: 1, 2: 281}`
4. `{1: 124, 2: 281, 0: 1}`
5. `{2: 281, 0: 1, 1: 124}`
6. `{2: 281, 1: 124, 0: 1}`

Как видите, порядок элементов может меняться при каждом запуске программы. Поэтому не стоит полагаться на определенный порядок элементов в map в своей программе. Если вам нужно сохранить порядок элементов, вам следует использовать другую структуру данных, такую как срез или структуру с явным определением порядка.

####  В чем разница make и new?
> **Ответ**:
В Go операторы `make` и `new` используются для выделения памяти, но они имеют разные цели и применяются для различных типов данных.

Оператор `new` используется для выделения памяти под указанный тип и возвращает указатель на эту область памяти. Синтаксис выглядит так: `new(T)`, где `T` - тип данных. Оператор `new` выделяет нулевое значение указанного типа и возвращает указатель на эту область памяти. При использовании `new` память выделяется, но не инициализируется нулевыми значениями для типа данных.

Пример использования оператора `new`:
```go
var s *string
s = new(string)
fmt.Println(s) // выводит адрес в памяти
```

Оператор `make` используется только для создания и инициализации срезов, карт (map) и каналов (channel). Синтаксис выглядит так: `make(T, size)`, где `T` - тип данных, а `size` - размер памяти для выделения (не для всех типов). Оператор `make` выделяет и инициализирует память для указанного типа данных и возвращает его значение (не указатель). Это означает, что `make` возвращает инициализированные значения для типа данных.

Пример использования оператора `make`:
```go
var s []int
s = make([]int, 5)
fmt.Println(s) // выводит [0 0 0 0 0]
```

Таким образом, различие между `make` и `new` заключается в том, что `new` используется для выделения нулевой памяти без инициализации, в то время как `make` используется для выделения памяти и инициализации для срезов, карт и каналов.

####  Сколько существует способов задать переменную типа slice или map?

**Ответ**: Проинициализировать переменную slice или map можно:
В языке Golang существует несколько способов задать переменную типа slice или map.

Для создания пустой переменной типа slice используется следующий синтаксис:
```go
var sliceName []elementType
```
Например:
```go
var numbersSlice []int
```

Если необходимо инициализировать slice с известными значениями, можно использовать оператор литерала slice:
```go
sliceName := []elementType{value1, value2, value3, ...}
```
Например:
```go
numbersSlice := []int{1, 2, 3, 4, 5}
```

Для создания пустой переменной типа map используется следующий синтаксис:
```go
var mapName map[keyType]valueType
```
Например:
```go
var studentGrades map[string]int
```

Также можно использовать оператор литерала map для инициализации переменной с известными парами ключ-значение:
```go
mapName := map[keyType]valueType{key1: value1, key2: value2, key3: value3, ...}
```
Например:
```go
studentGrades := map[string]int{
    "Alice": 90,
    "Bob": 85,
    "Charlie": 95,
}
```

Обратите внимание, что при использовании оператора литерала map, запятая ставится после каждой пары ключ-значение, кроме последней.

####  Что выведет данная программа и почему?

```
func update(p *int) {
  b := 2
  p = &b
}

func main() {
  var (
     a = 1
     p = &a
  )
  fmt.Println(*p)
  update(p)
  fmt.Println(*p)
}
```
> **Ответ**: 
Данная программа выведет:
```
1
1
```

В функции `update()`, переменная `b` инициализируется значением 2, и затем указатель `p` присваивается адресу переменной `b`. Однако, изменения, внесенные в `p`, остаются локальными и не влияют на значение `p` в функции `main()`. Это связано с тем, что в Go передача аргументов в функцию происходит по значению, и значение указателя `p` в `main()` не изменяется после вызова функции `update()`.

В `main()`, значение `*p` выводится дважды. В первом случае, оно равно 1, так как указатель `p` указывает на переменную `a`, которая имеет значение 1. Во втором случае, оно также равно 1, потому что изменения адреса указателя `p` внутри `update()` не влияют на исходный указатель `p` в `main()`.

####  Что выведет данная программа и почему?

```
func main() {
  wg := sync.WaitGroup{}
  for i := 0; i < 5; i++ {
     wg.Add(1)
     go func(wg sync.WaitGroup, i int) {
        fmt.Println(i)
        wg.Done()
     }(wg, i)
  }
  wg.Wait()
  fmt.Println("exit")
}
```
> **Ответ**:
Данная программа вероятно выведет следующий результат:

```
2
0
4
3
1
exit
```

Порядок вывода значений может отличаться при каждом запуске программы из-за горутин и их конкурентной природы.

В программе используется WaitGroup, чтобы дождаться завершения всех горутин, прежде чем продолжить выполнение. Для каждой итерации цикла создается новая горутина, которая выводит значение `i`.

Так как горутины выполняются параллельно, порядок вывода значений может быть произвольным. Однако, вероятно, что значения будут выведены в порядке от `0` до `4`, так как каждая горутина обеспечивается своим значением `i`.

После завершения работы всех горутин, программа выводит "exit".

####  Что выведет данная программа и почему?

```
func main() {
  n := 0
  if true {
     n := 1
     n++
  }
  fmt.Println(n)
}
```

> **Ответ**:

Данная программа выведет:
```
0
```

При объявлении переменной `n` внутри блока условия `if`, используется оператор `:=` для создания новой переменной `n`, локальной для этого блока. 

Внутри блока условия `n` увеличивается на 1, но это изменение применяется только к локальной переменной `n` внутри блока условия, а не к внешней переменной `n` в функции `main`.

При вызове `fmt.Println(n)` после блока условия, программа обратится к внешней переменной `n`, которая имеет значение 0, именно это значение и будет выведено.

####  Что выведет данная программа и почему?

```
func someAction(v []int8, b int8) {
  v[0] = 100
  v = append(v, b)
}

func main() {
  var a = []int8{1, 2, 3, 4, 5}
  someAction(a, 6)
  fmt.Println(a)
}
```
> **Ответ**:
Данная программа выведет:
```
[100 2 3 4 5]
```

В функции `someAction`, значение элемента `v[0]` изменяется на 100. Это изменение отражается на оригинальном срезе `a` в функции `main`, так как при передаче среза в функцию передается его указатель и любые изменения, сделанные внутри функции, будут отображаться на оригинальном срезе.

Однако, при вызове `append` в функции `someAction` и добавлении элемента `b` в срез `v`, происходит создание нового среза `v` с добавленным элементом `b`. Это новое значение среза `v` является локальной переменной внутри функции `someAction` и не влияет на оригинальный срез `a` в функции `main`.

Таким образом, после вызова `someAction(a, 6)`, оригинальный срез `a` будет иметь значения `[100 2 3 4 5]`.

####  Что выведет данная программа и почему?

```
func main() {
  slice := []string{"a", "a"}

  func(slice []string) {
     slice = append(slice, "a")
     slice[0] = "b"
     slice[1] = "b"
     fmt.Print(slice)
  }(slice)
  fmt.Print(slice)
}
```
> **Ответ**:
Данная программа выведет:
```
[b b a]
[a a]
```

При вызове анонимной функции, срез `slice` передается по значению, что означает, что внутри функции создается копия среза и все изменения происходят на этой копии, не затрагивая оригинальный срез в функции `main`.

Внутри анонимной функции, элементы среза `slice` изменяются на значения "b" по индексам 0 и 1, и после этого в срез `slice` добавляется элемент "a" с помощью `append`, создавая новый срез. Таким образом, внутри функции `fmt.Print` будет выведен срез `["b" "b" "a"]`.

Однако, после вызова анонимной функции, значение оригинального среза `slice` в функции `main` остается без изменений, поэтому во втором вызове `fmt.Print` будет выведен исходный срез `["a" "a"]`.