package main

import "fmt"

type Reception struct {
	next Department
}

func (r *Reception) execute(p *Patient) {
	if p.registrationDone {
		fmt.Println("Registro do paciente já foi feito")
		r.next.execute(p)
		return
	}

	fmt.Println("Recepção está registrando o paciente")
	p.registrationDone = true
	r.next.execute(p)
}

func (r *Reception) setNext(next Department) {
	r.next = next
}
