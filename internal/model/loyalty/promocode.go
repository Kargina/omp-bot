package loyalty

import "fmt"

var AllPromocodes = make([]Promocode, 0)

type Promocode struct {
	Promocode   string
	Description string
}

func (p *Promocode) String() string {
	return fmt.Sprintf("Promocode: %s\nDescription: %s", p.Promocode, p.Description)
}
