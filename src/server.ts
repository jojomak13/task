import express, { Express } from 'express';
import 'express-async-errors';

const app: Express = express();

app.use(express.json());

app.get('/', (req, res) => {
    res.json({msg: 'welcome'});
});

export default app;