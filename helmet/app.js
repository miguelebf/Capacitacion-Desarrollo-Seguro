const express = require("express");
const helmet = require("helmet");

const app = express();

// Configurar Helmet con opciones personalizadas
app.use(
    helmet({
        contentSecurityPolicy: {
            directives: {
                defaultSrc: ["'self'"],
                scriptSrc: ["'self'", "https://apis.google.com"],
            },
        },
        hsts: { maxAge: 31536000, includeSubDomains: true },
    })
);

// Eliminar cabecera X-Powered-By
app.disable("x-powered-by");


app.get("/", (req, res) => {
    res.send("Aplicación segura con Helmet!");
});

app.listen(3000, () => {
    console.log("Server Helmet http://localhost:3000");
});