import express from 'express';
import * as dotenv from 'dotenv';
import cors from 'cors'
dotenv.config();
const PORT = process.env.PORT
const app = express();
app.use(cors());
app.use(express.json({ limit: '50mb' }));


app.listen(PORT, () => {
  console.log(`Server listening on port ${PORT}`);
})
