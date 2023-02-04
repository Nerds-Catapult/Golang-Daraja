const mongoose = require('mongoose');
mongoose.set('strictQuery', true);
const connectDb = async () => {
    try {
        const conn = await mongoose.connect(process.env.DATABASE_URI);
        console.log(`MongoDB Connected: ${conn.connection.host}`);
    }
    catch (err) {
        console.error(err);
        process.exit(1);
    }
}

module.exports = connectDb
