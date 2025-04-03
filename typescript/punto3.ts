// Función sin parámetros ni retorno



function saludar(): void {
    console.log("¡Hola, mundo!");
}

saludar();


// Función con un parámetro y sin retorno

function saludarUsuario(nombre: string): void {
    console.log(`¡Hola, ${nombre}!`);
}

saludarUsuario("Juan");


// Función con varios parámetros y con retorno

function sumar(a: number, b: number): number {
    return a + b;
}

console.log(sumar(5, 7));


// Función dentro de una función

function externa() {
    console.log("Función externa ejecutada");

    function interna() {
        console.log("Función interna ejecutada");
    }

    interna();
}

externa();


// Uso de funciones nativas del lenguaje

console.log(Math.max(10, 20, 30));


// Variables globales y locales

let globalVar = "Soy una variable global";

function pruebaVariables() {
    let localVar = "Soy una variable local";
    console.log(globalVar); // Accede a la variable global
    console.log(localVar);  // Accede a la variable local
}

pruebaVariables();




// Dificultad extra (opcional):

// Función que recibe dos cadenas de texto y retorna un número

function procesarNumeros(texto1: string, texto2: string): number {
    let contador = 0;
    for (let i = 1; i <= 100; i++) {
        if (i % 3 === 0 && i % 5 === 0) {
            console.log(texto1 + texto2);
        } else if (i % 3 === 0) {
            console.log(texto1);
        } else if (i % 5 === 0) {
            console.log(texto2);
        } else {
            console.log(i);
            contador++;
        }
    }
    return contador;
}

console.log(procesarNumeros("Fizz", "Buzz"));


