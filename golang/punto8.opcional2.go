package main

import "fmt"

type PrintQueue struct {
	queue Queue8
}

func (p *PrintQueue) AddDocument(document string) {
	p.queue.Enqueue(document)
	fmt.Println("Documento agregado:", document)
}

func (p *PrintQueue) PrintDocument() {
	if doc, ok := p.queue.Dequeue(); ok {
		fmt.Println("Imprimiendo:", doc)
	} else {
		fmt.Println("No hay documentos en la cola")
	}
}

func main() {
	printer := PrintQueue{}
	printer.AddDocument("Reporte1.pdf")
	printer.AddDocument("Tesis.docx")
	printer.PrintDocument() // Imprimiendo Reporte1.pdf
	printer.PrintDocument() // Imprimiendo Tesis.docx
}
