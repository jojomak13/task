import {Request, Response} from 'express';
import { BadRequestError } from '../errors/BadRequestError';
import { Service } from '../models/Service';
import { Shipping } from '../models/Shipping';

export const create = async (data: any, _req: Request, res: Response) => {
    const service = await Service.findOne({code: data.service });

    if(!service!.types.includes(data.shippingType))
        throw new BadRequestError('Invalid shipping type');
    
    const shipping = Shipping.build({
        ...data,
        service,
        price: 45,
        shippingId: '45645646546556'
    });
    
    await shipping.save()

    res.status(201).json({
        status: true,
        msg: 'Shipping request created successfully',
        data: shipping,
    })
}