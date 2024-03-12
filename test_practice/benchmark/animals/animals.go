package animals

type Duck struct {
	Name string
}

func NewDuck(name string) *Duck {
	return &Duck{Name: name}
}

func (d *Duck) Say() string {
	return d.Name + " says quack"
}

func (d *Duck) Eat(food string) string {
	return d.Name + " ate " + food
}
