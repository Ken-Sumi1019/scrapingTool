package set

type Set struct {
	Data map[interface{}]int
}

func MakeSet() *Set {
	return &Set{Data: map[interface{}]int{}}
}

func (v *Set) Add(s interface{})  {
	v.Data[s] = 0
}

func (v *Set) Exist(s interface{}) bool {
	if _,ok := v.Data[s];ok {
		return true
	}
	return false
}

func (v *Set) Erase(s interface{}) bool {
	if v.Exist(s) {
		delete(v.Data,s)
		return true
	}
	return false
}

