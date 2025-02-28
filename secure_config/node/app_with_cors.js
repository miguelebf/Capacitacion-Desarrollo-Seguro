const express = require("express");
const cors = require("cors");

const app = express();

// Configurar CORS con opciones personalizadas
const corsOptions = {
    origin: "https://example.com",
    methods: "GET,HEAD,PUT,PATCH,POST,DELETE",
    preflightContinue: false,
    optionsSuccessStatus: 204,
};

app.use(cors(corsOptions));


app.get("/", (req, res) => {
    res.send("Aplicación segura con CORS!");
});

app.listen(8080, () => {
    console.log("Server CORS: 8080");
});
