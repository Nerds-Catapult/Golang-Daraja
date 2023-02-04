const express = require('express');
const cors = require('cors');
const dotenv = require('dotenv').config();
const PORT = process.env.PORT
const connectDb = require('./database/connect');                                                                                                            


const app = express();
app.use(cors());
app.use(express.json());
app.use(express.urlencoded({ extended: true }));

connectDb();

app.listen(PORT, () => { console.log(`server started on port ${PORT}`)})
