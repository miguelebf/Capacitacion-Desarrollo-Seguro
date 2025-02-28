const crypto = require('crypto');

// Función para cifrar texto con AES-GCM
function encryptAES(plaintext, key) {
    const iv = crypto.randomBytes(12); // GCM usa un nonce de 12 bytes
    const cipher = crypto.createCipheriv('aes-256-gcm', key, iv);

    let encrypted = cipher.update(plaintext, 'utf8', 'base64');
    encrypted += cipher.final('base64');
    const authTag = cipher.getAuthTag().toString('base64');

    return iv.toString('base64') + encrypted + authTag;
}

// Función para descifrar texto con AES-GCM
function decryptAES(ciphertext, key) {
    const iv = Buffer.from(ciphertext.slice(0, 16), 'base64');
    const authTag = Buffer.from(ciphertext.slice(-24), 'base64');
    const encrypted = ciphertext.slice(16, -24);

    const decipher = crypto.createDecipheriv('aes-256-gcm', key, iv);
    decipher.setAuthTag(authTag);

    let decrypted = decipher.update(encrypted, 'base64', 'utf8');
    decrypted += decipher.final('utf8');

    return decrypted;
}

const key = crypto.randomBytes(32); // Clave de 32 bytes (AES-256)
const plaintext = "Hola, mundo! 1234";

const encrypted = encryptAES(plaintext, key);
console.log("Texto cifrado:", encrypted);

const decrypted = decryptAES(encrypted, key);
console.log("Texto descifrado:", decrypted);
