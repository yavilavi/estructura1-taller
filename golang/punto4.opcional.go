package main

import "fmt"

type ContactBook struct {
	contacts map[string]string
}

func (cb *ContactBook) AddContact(name, phone string) {
	cb.contacts[name] = phone
}

func (cb *ContactBook) UpdateContact(name, phone string) {
	if _, exists := cb.contacts[name]; exists {
		cb.contacts[name] = phone
	} else {
		fmt.Println("Contacto no encontrado")
	}
}

func (cb *ContactBook) DeleteContact(name string) {
	delete(cb.contacts, name)
}

func (cb *ContactBook) SearchContact(name string) {
	if phone, exists := cb.contacts[name]; exists {
		fmt.Println("Teléfono:", phone)
	} else {
		fmt.Println("Contacto no encontrado")
	}
}

func (cb *ContactBook) ListContacts() {
	fmt.Println("Contactos:", cb.contacts)
}

func main() {
	contactBook := ContactBook{contacts: make(map[string]string)}
	var option, name, phone string

	for {
		fmt.Println("Elige operación (add, update, delete, search, list, exit):")
		fmt.Scanln(&option)

		switch option {
		case "add":
			fmt.Println("Nombre:")
			fmt.Scanln(&name)
			fmt.Println("Teléfono:")
			fmt.Scanln(&phone)
			contactBook.AddContact(name, phone)
		case "update":
			fmt.Println("Nombre:")
			fmt.Scanln(&name)
			fmt.Println("Nuevo Teléfono:")
			fmt.Scanln(&phone)
			contactBook.UpdateContact(name, phone)
		case "delete":
			fmt.Println("Nombre:")
			fmt.Scanln(&name)
			contactBook.DeleteContact(name)
		case "search":
			fmt.Println("Nombre:")
			fmt.Scanln(&name)
			contactBook.SearchContact(name)
		case "list":
			contactBook.ListContacts()
		case "exit":
			return
		default:
			fmt.Println("Opción no válida")
		}
	}
}
