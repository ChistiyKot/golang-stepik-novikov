/*
В рамках этого урока мы постарались представить себе уже привычные нам переменные и функции, как объекты из реальной жизни. Чтобы закрепить результат мы предлагаем вам небольшую творческую задачу...
*/
package main

type Terminator struct {
	On    bool
	Ammo  int
	Power int
}

func (o *Terminator) Shoot() bool {
	if !o.On || o.Ammo == 0 {
		return false
	}

	o.Ammo--

	return true
}

func (o *Terminator) RideBike() bool {
	if !o.On || o.Power == 0 {
		return false
	}

	o.Power--

	return true
}

func main() {
	testStruct := new(Terminator)

	testStruct.Shoot()
	testStruct.RideBike()
}
