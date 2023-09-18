[Назад](/README.md)

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