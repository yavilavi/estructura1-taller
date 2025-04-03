// Estructuras de Datos en TypeScript

// Arrays

let array: number[] = [3, 1, 4, 1, 5];
array.push(9); // Inserción
array.splice(1, 1); // Borrado
array[2] = 2; // Actualización
array.sort((a, b) => a - b); // Ordenación


// Sets

let set: Set<number> = new Set([3, 1, 4, 1, 5]);
set.add(9); // Inserción
set.delete(1); // Borrado


// Maps

let map: Map<string, number> = new Map([
    ["uno", 1],
    ["dos", 2]
]);
map.set("tres", 3); // Inserción
map.delete("dos"); // Borrado
map.set("tres", 33); // Actualización


// Stack (Pila) con Array

class Stack<T> {
    private items: T[] = [];
    push(item: T) { this.items.push(item); }
    pop(): T | undefined { return this.items.pop(); }
}


// Queue (Cola) con Array

class Queue<T> {
    private items: T[] = [];
    enqueue(item: T) { this.items.push(item); }
    dequeue(): T | undefined { return this.items.shift(); }
}


// Agenda de contactos en terminal

import * as readline from "readline";

interface Contact {
    name: string;
    phone: string;
}

class ContactBook {
    private contacts: Contact[] = [];

    addContact(name: string, phone: string) {
        if (!/^[0-9]{1,11}$/.test(phone)) {
            console.log("Número de teléfono inválido.");
            return;
        }
        this.contacts.push({ name, phone });
        console.log("Contacto agregado.");
    }

    updateContact(name: string, newPhone: string) {
        let contact = this.contacts.find(c => c.name === name);
        if (contact) contact.phone = newPhone;
        else console.log("Contacto no encontrado.");
    }

    deleteContact(name: string) {
        this.contacts = this.contacts.filter(c => c.name !== name);
        console.log("Contacto eliminado.");
    }

    searchContact(name: string) {
        let contact = this.contacts.find(c => c.name === name);
        if (contact) console.log(contact);
        else console.log("No encontrado.");
    }

    listContacts() {
        console.log("Contactos:", this.contacts);
    }
}

const rl = readline.createInterface({ input: process.stdin, output: process.stdout });
const contactBook = new ContactBook();

function mainMenu() {
    rl.question("Elige operación (add, update, delete, search, list, exit): ", (option) => {
        switch (option) {
            case "add":
                rl.question("Nombre: ", name => {
                    rl.question("Teléfono: ", phone => {
                        contactBook.addContact(name, phone);
                        mainMenu();
                    });
                });
                break;
            case "update":
                rl.question("Nombre: ", name => {
                    rl.question("Nuevo Teléfono: ", phone => {
                        contactBook.updateContact(name, phone);
                        mainMenu();
                    });
                });
                break;
            case "delete":
                rl.question("Nombre: ", name => {
                    contactBook.deleteContact(name);
                    mainMenu();
                });
                break;
            case "search":
                rl.question("Nombre: ", name => {
                    contactBook.searchContact(name);
                    mainMenu();
                });
                break;
            case "list":
                contactBook.listContacts();
                mainMenu();
                break;
            case "exit":
                rl.close();
                break;
            default:
                console.log("Opción no válida");
                mainMenu();
        }
    });
}

mainMenu();


