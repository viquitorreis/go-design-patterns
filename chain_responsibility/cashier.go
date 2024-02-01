package main

import "fmt"

type Cashier struct {
	next Department
}

func (c *Cashier) execute(p *Patient) {
	if p.paymentDone {
		fmt.Println("Pagamento já foi feito")
	}

	fmt.Println("Caixeiro está recebendo o dinheiro do paciente")
}

func (c *Cashier) setNext(next Department) {
	c.next = next
}
