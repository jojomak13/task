import { ServiceDocument } from "../models/Service";
import { ShippingDocument } from "../models/Shipping";

export interface IShipping {
    readonly code: String;

    getService(): Promise<ServiceDocument | null>;

    create(data: Partial<ShippingDocument>): {id: String, price: Number};
}