// Ejemplos de Operadores en TypeScript



// Operadores Aritméticos
let suma = 5 + 3;
let resta = 10 - 4;
let multiplicacion = 6 * 2;
let division = 8 / 2;
let modulo = 10 % 3;
let exponente = 2 ** 3;
let divisionEntera = Math.floor(10 / 3);

// Operadores de Comparación
let esIgual = 5 === 5;
let esDiferente = 5 !== 3;
let esMayor = 6 > 2;
let esMenor = 2 < 8;

// Operadores Lógicos
let esVerdadero = true && false;
let esFalso = true || false;
let negacion = !true;

// Operadores de Asignación
let x = 10;
x += 5;
x -= 2;

// Operadores de Identidad
let esMismoObjeto = x === x;
let esDiferenteObjeto = x !== 15;

// Operadores de Pertenencia
let texto = "Typescript";
let esParte = texto.includes("y");
let noEsParte = !texto.includes("p");

// Operadores de Bits
let bitAnd = 5 & 3;
let bitOr = 5 | 3;
let bitXor = 5 ^ 3;
let bitNot = ~5;
let bitDesplazarIzq = 5 << 1;
let bitDesplazarDer = 5 >> 1;

console.log("suma: ",suma, "resta: ",resta, "multiplicacion: ",multiplicacion, "division: ",division, "modulo: ",modulo, "exponente: ",exponente, "divisionEntera: ",divisionEntera);
console.log("esIgual ", esIgual, "esDiferente ", esDiferente, "esMayor ", esMayor, "esMenor ", esMenor);
console.log("esVerdadero ",esVerdadero, "esFalso ",esFalso, "negacion ",negacion);
console.log("x ",x, "esMismoObjeto ",esMismoObjeto, "esDiferenteObjeto ",esDiferenteObjeto);
console.log("esParte ",esParte, "noEsParte ",noEsParte);
console.log("bitAnd ",bitAnd, "bitOr ",bitOr, "bitXor ",bitXor, "bitNot ",bitNot, "bitDesplazarIzq ",bitDesplazarIzq, "bitDesplazarDer ",bitDesplazarDer);




// Estructuras de Control en TypeScript
console.log("Estructuras de Control en TypeScript");

// Condicionales
console.log("Condicionales");
let numero = 10;
if (numero > 5) {
    console.log("El número es mayor que 5");
} else if (numero === 5) {
    console.log("El número es 5");
} else {
    console.log("El número es menor que 5");
}

// Bucles
console.log("Bucles");
for (let i = 0; i < 5; i++) {
    console.log(`Iteración ${i}`);
}

let contador = 0;
while (contador < 5) {
    console.log(`Contador: ${contador}`);
    contador++;
}

// Manejo de Excepciones
console.log("Manejo de Excepciones");
try {
    let resultado = 10 / 0;
    if (!isFinite(resultado)) throw new Error("División por cero");
} catch (error) {
    console.error("Error: División por cero");
}




// Dificultad extra (opcional)

console.log("Dificultad extra (opcional)")
console.log("Números entre 10 y 55, pares, excluyendo 16 y múltiplos de 3")
// Números entre 10 y 55, pares, excluyendo 16 y múltiplos de 3
for (let num = 10; num <= 55; num++) {
    if (num % 2 === 0 && num !== 16 && num % 3 !== 0) {
        console.log(num);
    }
}

// Números primos del 1 al 1000
function esPrimo(n: number): boolean {
    if (n < 2) return false;
    for (let i = 2; i <= Math.sqrt(n); i++) {
        if (n % i === 0) return false;
    }
    return true;
}

let primos: number[] = [];
for (let num = 1; num <= 1000; num++) {
    if (esPrimo(num)) primos.push(num);
}
console.log("Primos: ", primos);

// Serie de Fibonacci
function fibonacci(n: number): void {
    let a = 0, b = 1;
    for (let i = 0; i < n; i++) {
        console.log(a," ");
        [a, b] = [b, a + b];
    }
}
console.log("Fibonacci: ");
fibonacci(15);

