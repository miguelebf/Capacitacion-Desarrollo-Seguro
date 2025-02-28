const express = require('express');
const { Pool } = require('pg');
const app = express();
const port = 8080;

const pool = new Pool({
    user: 'postgres',
    host: 'localhost',
    database: 'testdb',
    password: 'postgres',
    port: 5432,
});

app.get('/vulnerable', async (req, res) => {
    const userID = req.query.id;
    const query = `SELECT name FROM users WHERE id =  ${userID}`;
    try {
        const result = await pool.query(query);
        const names = result.rows.map(row => row.name);
        res.send(`Names: ${names}`);
    } catch (err) {
        res.status(500).send('Error querying database');
        console.error(err);
    }
});

app.listen(port, () => {
    console.log(`Server started at SQL Injection :${port}`);
});
