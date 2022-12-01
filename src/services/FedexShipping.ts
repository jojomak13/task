import { ShippingDocument } from "../models/Shipping";
import { IShipping } from "./IShipping";
import { v4 as uuidv4 } from 'uuid';
import { Service, ServiceDocument } from "../models/Service";

class FedexShipping implements IShipping {
    
    readonly code: String = 'fedex';

    async getService(): Promise<ServiceDocument | null>  {
        return await Service.findOne({code: this.code });
    }

    // the third party logic should be replaced here
    create(data: Partial<ShippingDocument>): any {
        const response = {
            id: uuidv4().toString(),
            price: Math.floor(Math.random() * 100)
        } 

        return response;
    }
}

export { FedexShipping }