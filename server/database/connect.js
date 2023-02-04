const mongoose = require('mongoose');

const connectDB = (url) => {
    mongoose.set('strictQuery', true);

    mongoose.connect(url, {
        
    });
}
