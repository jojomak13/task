import app from './src/server';

const start = async () => {
    const port = process.env.PORT || 8080;

    app.listen(port, () => {
        console.log(`[Shipping Service] Running on port ${port}`);
    });
}

start();