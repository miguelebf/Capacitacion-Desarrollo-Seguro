const dotenv = require('dotenv');

dotenv.config();

console.log("DATABASE_URL:", process.env.DATABASE_URL);
console.log("PORT:", process.env.PORT);
console.log("DEBUG:", process.env.DEBUG);

