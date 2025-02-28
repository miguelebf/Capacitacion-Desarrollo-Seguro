const express = require('express');
const Joi = require('joi');

const app = express();
app.use(express.json());

// Definir el esquema de validación con Joi
const userSchema = Joi.object({
    name: Joi.string().required(),
    email: Joi.string().email().required(),
    age: Joi.number().integer().min(18).max(100).required()
});

app.post('/users', (req, res) => {
    const { error, value } = userSchema.validate(req.body);
    if (error) {
        return res.status(400).json({ error: error.details[0].message });
    }
    res.status(200).json({ message: 'Usuario creado!', user: value });
});

app.listen(8080, () => {
    console.log('Server running on port 8080');
});
