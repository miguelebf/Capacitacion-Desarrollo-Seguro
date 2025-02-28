const express = require('express');
const app = express();
const port = 8080;

// Datos de usuarios
const users = {
    1: { id: 1, username: "user1", email: "user1@example.com" },
    2: { id: 2, username: "user2", email: "user2@example.com" }
};

// Middleware para parsear el header x-user-id
app.use((req, res, next) => {
    const authenticatedUserID = req.headers['x-user-id'];
    if (!authenticatedUserID) {
        return res.status(400).json({ error: "Invalid user ID" });
    }
    req.authenticatedUserID = parseInt(authenticatedUserID, 10);
    if (isNaN(req.authenticatedUserID)) {
        return res.status(400).json({ error: "Invalid user ID" });
    }
    next();
});

// Ruta para obtener un usuario
app.get('/user', (req, res) => {
    const authenticatedUserID = req.authenticatedUserID;

    const userID = parseInt(req.query.id, 10);
    if (isNaN(userID)) {
        return res.status(400).json({ error: "Invalid user ID" });
    }

    // ------------ Verificación de autorización ------------
    if (userID !== authenticatedUserID) {
        return res.status(401).json({ error: "Unauthorized" });
    }

    const user = users[userID];
    if (!user) {
        return res.status(404).json({ error: "User not found" });
    }

    res.json(user);
});

// Configuración del servidor con timeouts
const server = app.listen(port, () => {
    console.log(`Node Secure Version Server Broken Object Level Authorization started at http://localhost:${port}`);
});

// Configuración de timeouts
server.keepAliveTimeout = 15000; // 15 segundos
server.headersTimeout = 10000; // 10 segundos