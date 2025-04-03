// Ejemplos de operaciones con cadenas en TypeScript

const str = "Hello, TypeScript!";

// Acceso a caracteres específicos
console.log(str[0]); // "H"
console.log(str.charAt(1)); // "e"

// Longitud
console.log(str.length); // 17

// Subcadenas
console.log(str.substring(0, 5)); // "Hello"
console.log(str.slice(-10, -1)); // "TypeScrip"

// Concatenación
console.log(str + " is awesome!");
console.log(str.concat(" Let's code!"));

// Repetición
console.log("Ha! ".repeat(3)); // "Ha! Ha! Ha! "

// Recorrido
for (const char of str) {
    console.log(char);
}

// Conversión a mayúsculas y minúsculas
console.log(str.toUpperCase());
console.log(str.toLowerCase());

// Reemplazo
console.log(str.replace("TypeScript", "JavaScript"));

// División y unión
const words = str.split(" ");
console.log(words);
console.log(words.join("-"));

// Interpolación
const name = "John";
console.log(`Hello, ${name}!`);

// Verificaciones
console.log(str.includes("Type")); // true
console.log(str.startsWith("Hello")); // true
console.log(str.endsWith("!")); // true

// Dificultad extra

function isPalindrome(word: string): boolean {
    const cleaned = word.toLowerCase().replace(/\s/g, "");
    return cleaned === cleaned.split('').reverse().join('');
}

function isAnagram(word1: string, word2: string): boolean {
    const normalize = (word: string) => word.toLowerCase().replace(/\s/g,"").split('').sort().join('');
    return normalize(word1) === normalize(word2);
}

function isIsogram(word: string): boolean {
    const letters = new Set<string>();
    for (const char of word.toLowerCase()) {
        if (letters.has(char)) return false;
        letters.add(char);
    }
    return true;
}

// Pruebas
console.log("isPalindrome", isPalindrome("radar")); // true
console.log("isAnagram", isAnagram("listen", "silent")); // true
console.log("isIsogram", isIsogram("background")); // true

