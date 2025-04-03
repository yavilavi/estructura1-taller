// Ejemplos de variables "por valor" y "por referencia"

// Por Valor





// Cuando se asigna una variable "por valor", se crea una copia independiente del valor original.



// Los tipos de datos primitivos (números, cadenas, booleanos, etc.) se asignan por valor.

let num1: number = 10;
let num2: number = num1; // num2 recibe una copia de num1

num2 = 20; // Modificamos num2

console.log(num1); // num1 sigue siendo 10
console.log(num2); // num2 ahora es 20

let str1: string = "Hola";
let str2: string = str1; // str2 recibe una copia de str1

str2 = "Mundo"; // Modificamos str2

console.log(str1); // str1 sigue siendo "Hola"
console.log(str2); // str2 ahora es "Mundo"



// Por Referencia





// Cuando se asigna una variable "por referencia", se crea una referencia a la ubicación de memoria del valor original.



// Los tipos de datos no primitivos (objetos, arreglos, etc.) se asignan por referencia.

let arr1: number[] = [1, 2, 3];
let arr2: number[] = arr1; // arr2 referencia el mismo arreglo que arr1

arr2[0] = 10; // Modificamos el arreglo a través de arr2

console.log(arr1); // arr1 también se ve modificado: [10, 2, 3]
console.log(arr2); // arr2: [10, 2, 3]

let obj1: { prop: number } = { prop: 1 };
let obj2: { prop: number } = obj1; // obj2 referencia el mismo objeto que obj1

obj2.prop = 2; // Modificamos el objeto a través de obj2

console.log(obj1); // obj1 también se ve modificado: { prop: 2 }
console.log(obj2); // obj2: { prop: 2 }

// Funciones y Paso de Parámetros

// Por Valor





// Cuando se pasa un parámetro "por valor" a una función, se crea una copia local de ese valor dentro de la función.



// Las modificaciones dentro de la función no afectan el valor original fuera de la función.

    function modificarValor(num: number): void {
    num = 20;
}
let miNumero: number = 10;
modificarValor(miNumero);
console.log(miNumero); // miNumero sigue siendo 10

// Por Referencia:





// Cuando se pasa un parámetro "por referencia" a una función, la función recibe una referencia a la ubicación de memoria del valor original.



// Las modificaciones dentro de la función afectan el valor original fuera de la función.

function modificarReferencia(arr: number[]): void {
    arr[0] = 10;
}

let miArreglo: number[] = [1, 2, 3];
modificarReferencia(miArreglo);

console.log(miArreglo); // miArreglo ahora es [10, 2, 3]



// Dificultad extra

// En el ejemplo de "intercambioPorValor", los valores originales no cambian, mientras que en "intercambioPorReferencia", el arreglo original se modifica.

// Intercambio por valor
    function intercambioPorValor(a: number, b: number): [number, number] {
    let temp = a;
    a = b;
    b = temp;
    return [a, b];
}

let val1 = 5;
let val2 = 10;
let nuevosValoresValor = intercambioPorValor(val1, val2);

console.log("Originales:", val1, val2); // Originales: 5 10
console.log("Nuevos:", nuevosValoresValor[0], nuevosValoresValor[1]); // Nuevos: 10 5

// Intercambio por referencia
function intercambioPorReferencia(arr: number[]): void {
    let temp = arr[0];
    arr[0] = arr[1];
    arr[1] = temp;
}

let refArr = [15, 20];
intercambioPorReferencia(refArr);

console.log("Referencia:", refArr[0], refArr[1]); // Referencia: 20 15

