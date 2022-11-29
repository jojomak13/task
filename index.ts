import mongoose from 'mongoose';
import app from './src/server';

const start = async () => {
    const envKeys = ['MONGO_URI'];

    for (let key of envKeys) {
        if (!process.env[key]) throw new Error(`[${key}] not found`);
    }

    try {
        mongoose.connect(process.env.MONGO_URI!);
    } catch (err) {
        console.log(err);
    }

    const port = process.env.PORT || 8080;

    app.listen(port, () => {
        console.log(`[Shipping Service] Running on port ${port}`);
    });
};

start();
