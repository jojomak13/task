import express, { Express } from 'express';
import 'express-async-errors';
import { NotFoundError } from './errors/NotFoundError';
import { errorHandler } from './middlewares/errorHandler';

const app: Express = express();

app.use(express.json());

app.get('/', (req, res) => {
    res.json({ msg: 'welcome' });
});

app.use(() => {
    throw new NotFoundError();
});

app.use(errorHandler);

export default app;
