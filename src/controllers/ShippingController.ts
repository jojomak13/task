import {Request, Response} from 'express';
import { ShippingFactory } from '../services/ShippingFactory';

export const create = async (data: any, _req: Request, res: Response) => { 
    res.status(201).json({
        status: true,
        msg: 'Shipping request created successfully',
        data: await ShippingFactory.create(data),
    })
}