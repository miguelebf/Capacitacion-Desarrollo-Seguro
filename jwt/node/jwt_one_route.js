const express = require('express');
const jwt = require('jsonwebtoken');
const multer = require('multer');
const cookieParser = require('cookie-parser');

const app = express();
const jwtKey = 'your_secret_key';
const jwtExpirySeconds = 86400; // 24 hours

const upload = multer();
app.use(cookieParser());

function generateToken(username) {
    const token = jwt.sign({ username }, jwtKey, {
        algorithm: 'HS256',
        expiresIn: jwtExpirySeconds,
    });
    return token;
}

function validateToken(token) {
    try {
        const payload = jwt.verify(token, jwtKey);
        return payload;
    } catch (e) {
        return null;
    }
}

app.post('/login', upload.none(), (req, res) => {
    const { username, password } = req.body;

    // For demonstration, we're assuming "admin"/"password" as valid credentials.
    const username_database = 'admin';
    const password_database = 'password';

    if (username !== username_database || password !== password_database) {
        console.log(username, password);
        return res.status(401).send('Invalid credentials');
    }

    const token = generateToken(username);
    res.cookie('token', token, { maxAge: jwtExpirySeconds * 1000 });
    res.end();
});

app.get('/protected', (req, res) => {
    const token = req.cookies.token;

    if (!token) {
        return res.status(401).send('No token');
    }

    const payload = validateToken(token);

    if (!payload) {
        return res.status(401).send('Invalid token');
    }

    res.send(`Hello, ${payload.username}!`);
});

app.listen(8080, () => {
    console.log('JWT Server One Route started at :8080');
});
