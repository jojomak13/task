import { BadRequestError } from "../errors/BadRequestError";
import { Service } from "../models/Service";
import { ShippingDocument } from "../models/Shipping";
import { FedexShipping } from "./FedexShipping";
import { ShippingService } from "./ShippingService";
import { UpsShipping } from "./UpsShipping";

class ShippingFactory {
    static async create(data: any): Promise<ShippingDocument> {     
        const service = await Service.findOne({code: data.service });

        if(!service)
            throw new BadRequestError(`Cannot load ${data.service} shipping service`);
        
        if(!service.types.includes(data.shippingType))
            throw new BadRequestError('Invalid shipping type');
        
        const shippingService = new ShippingService(data);

        if(service.code === 'fedex')
            return await shippingService.create(new FedexShipping());
        else
            return await shippingService.create(new UpsShipping());

    }
}

export { ShippingFactory }