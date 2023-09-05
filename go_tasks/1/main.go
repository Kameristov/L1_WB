package main

type Human struct {
	Weight float64
}

func (h Human) GetWeight() float64 {
	return h.Weight
}

type Action struct {
	Name string
}

func (a Action) GetName() string {
	return a.Name
}

func main() {

}
