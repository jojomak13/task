import { Request, Response, Router } from 'express';
import { RequestValidationError } from '../errors/RequestValidationError';
import { CreateShippingRequest } from '../requests/CreateShippingRequest';
import * as ShippingController from '../controllers/ShippingController';

const router = Router();

router.get('/', async (_req, res) => {
    res.json({ msg: 'welcome' });
});

router.post('/', async (req: Request, res: Response) => {
    const data = await CreateShippingRequest.validateAsync(req.body, {
        abortEarly: false,
        stripUnknown: true,
    }).catch(err => {
        throw new RequestValidationError(err);
    });

    await ShippingController.create(data, req, res); 
});

export default router;