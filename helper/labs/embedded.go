package labs

type Dog struct {
	Name string
}

type Animal struct {
	*Dog
	Age int
}

// 构造函数
func NewAnimal(name string, age int) *Animal {
	return &Animal{
		Dog: &Dog{Name: name},
		Age: age,
	}
}
