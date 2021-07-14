package product

type Product struct {
	Name string
	Cost float64
}

func (p *Product) setCost(cost float64) {
	p.Cost = cost
}
