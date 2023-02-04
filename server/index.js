const express = require('express');
const cors = require('cors');
const dotenv = require('dotenv').config();
const PORT = process.env.PORT
const connectDb = require('./database/connect');
const postRoutes = require('./routes/postRoutes');
const dalleRoutes = require('./routes/dalleRoutes');                                                                                                         
const app = express();
app.use('/api/v1/posts', postRoutes);
app.use('/api/v1/dalle', dalleRoutes);


app.use(cors());
app.use(express.json());
app.use(express.urlencoded({ extended: true }));

connectDb();

app.listen(PORT, () => { console.log(`server started on port ${PORT}`)})
