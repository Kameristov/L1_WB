/*
Реализовать паттерн «адаптер» на любом примере.
*/
package main

import "fmt"

func main() {
	repoPostgres := &Postgres{}
	repoMongo := &Adapter{&MongoDB{}}

	repoPostgres.Set()
	repoMongo.Set()
}

//----------------------------------------------------------------
// Интерфейс работы с репозиторием
type Repo interface {
	Set()
}

//----------------------------------------------------------------
// Структура репозитория с Postgres
type Postgres struct{}

// Метод  удовлетворящющий интерфейсу Repo
func (p *Postgres) Set() {
	fmt.Printf("Отправка данных в Postgres\n")
}

//----------------------------------------------------------------
// Структура препозитория с MongoDB
type MongoDB struct{}

func (m *MongoDB) SetToMongo() {
	fmt.Printf("Отправка данных в MongoDB\n")
}

//----------------------------------------------------------------
// Адаптер для реализации MongoDB
type Adapter struct {
	mdb *MongoDB
}
// Метод Адаптера удовлетворящющий интерфейсу Repo и вызывающий метод структуры MongoDB
func (a *Adapter) Set() {
	a.mdb.SetToMongo()
}

//----------------------------------------------------------------
