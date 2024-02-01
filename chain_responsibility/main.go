package main

func main() {
	cashier := &Cashier{}

	// Após os medicamentos o paciente deve fazer o pagamento
	medical := &Medical{}
	medical.setNext(cashier)

	// Doutor deve enviar paciente para receber medicamentos
	doctor := &Doctor{}
	doctor.setNext(medical)

	// Receção deve mandar o paciente para o Doutor
	reception := &Reception{}
	reception.setNext(doctor)

	patient := &Patient{name: "Victor"}
	// Paciente começa a jornada
	reception.execute(patient)
}
