package main

type Doctor struct {
	next Department
}

func (d *Doctor) execute(p *Patient) {
	if p.doctorCheckDone {
		println("Paciente já foi atendido pelo médico")
		d.next.execute(p)
		return
	}

	println("Médico está atendendo o paciente")
	p.doctorCheckDone = true
	d.next.execute(p)
}

func (d *Doctor) setNext(next Department) {
	d.next = next
}
