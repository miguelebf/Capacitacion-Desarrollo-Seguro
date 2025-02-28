const crypto = require('crypto');

// Función para hashear con SHA-256
function hashSHA256(data) {
    return crypto.createHash('sha256').update(data).digest('hex');
}

const text = "Hola, mundo! 1234";
const hashedText = hashSHA256(text);

console.log("Texto original:", text);
console.log("Hash SHA-256:", hashedText);
