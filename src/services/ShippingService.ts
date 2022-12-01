import { BadRequestError } from "../errors/BadRequestError";
import { Shipping, ShippingDocument } from "../models/Shipping";
import { IShipping } from "./IShipping";

class ShippingService {
    constructor(private data: ShippingDocument) { }

    async create(shippingService: IShipping): Promise<ShippingDocument> {
        const request = shippingService.create(this.data);
        const service = await shippingService.getService();
        
        if(!service)
            throw new BadRequestError(`Cannot load ${shippingService.code} shipping service`);

        const shipping = Shipping.build({
            ...this.data,
            service,
            price: request.price,
            shippingId: request.id
        });
        
        return await shipping.save();
    }
}

export { ShippingService }