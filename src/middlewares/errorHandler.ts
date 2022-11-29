import { Request, Response, NextFunction } from 'express';
import { CustomError } from '../errors/CustomError';

const errorHandler = (
    error: Error,
    _req: Request,
    res: Response,
    _next: NextFunction
) => {
    if (error instanceof CustomError) {
        return res.status(error.statusCode).send(error.serialize());
    }

    console.error(error);

    return res.status(400).send({
        type: 'UnhandeledError',
        errors: [{ message: 'Something went wrong' }],
    });
};

export { errorHandler };
