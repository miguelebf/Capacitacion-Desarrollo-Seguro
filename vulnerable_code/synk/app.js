const express = require("express");
const bodyParser = require("body-parser");
const sqlite3 = require("sqlite3").verbose();

const app = express();
const db = new sqlite3.Database(":memory:");

// Configuración básica
app.use(bodyParser.json());

// Crear una tabla de usuarios en la base de datos
db.serialize(() => {
    db.run("CREATE TABLE users (id INTEGER PRIMARY KEY, username TEXT, password TEXT)");
    db.run("INSERT INTO users (username, password) VALUES ('admin', 'admin123')");
});

// Ruta vulnerable a inyección SQL
app.get("/users", (req, res) => {
    const userId = req.query.id;
    const query = `SELECT * FROM users WHERE id = ${userId}`; // Inyección SQL

    db.all(query, (err, rows) => {
        if (err) {
            return res.status(500).send("Error en la base de datos");
        }
        res.json(rows);
    });
});

// Ruta vulnerable a XSS (Cross-Site Scripting)
app.post("/comment", (req, res) => {
    const comment = req.body.comment;
    res.send(`<p>Comentario recibido: ${comment}</p>`); // XSS
});

// Ruta con dependencia vulnerable (usando una versión antigua de lodash)
app.get("/process", (req, res) => {
    const _ = require("lodash");
    const data = { user: req.query.user };
    const processedData = _.merge({}, data); // Uso de lodash
    res.json(processedData);
});

// Iniciar el servidor
app.listen(3000, () => {
    console.log("API escuchando en http://localhost:3000");
});