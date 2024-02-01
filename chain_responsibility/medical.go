package main

import "fmt"

type Medical struct {
	next Department
}

func (m *Medical) execute(p *Patient) {
	if p.medicationDone {
		fmt.Println("Paciente já foi medicado")
		m.next.execute(p)
		return
	}

	fmt.Println("Medicando o paciente")
	p.medicationDone = true
	m.next.execute(p)
}

func (m *Medical) setNext(next Department) {
	m.next = next
}
