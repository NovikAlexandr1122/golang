// не законченный пример

package main

type testStruct struct {
	On    bool
	Ammo  int
	Power int
}

func (s testStruct) Shot() bool {

	isTrue := false

	if s.Ammo > 0 {
		s.Ammo = s.Ammo - 1
		isTrue = true
	} else if s.Ammo < 0 {
		isTrue = false
	}
	return isTrue
}

func (r testStruct) RideBike() bool {

	isTrue := false

	if r.Power > 0 {
		r.Power = r.Power - 1
		isTrue = true
	} else if r.Power < 0 {
		isTrue = false
	}
	return isTrue
}

func main() {
	testStruct := new(testStruct)
}
