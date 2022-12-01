import express, { Express } from 'express';
import 'express-async-errors';
import { NotFoundError } from './errors/NotFoundError';
import { errorHandler } from './middlewares/errorHandler';
import routes from './routes';

const app: Express = express();

app.use(express.json());

app.use('/api/shipping', routes);

app.use(() => {
    throw new NotFoundError();
});

app.use(errorHandler);

export default app;
