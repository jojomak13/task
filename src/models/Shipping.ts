import mongoose, {Schema, model} from 'mongoose';
import { Service, ServiceDocument } from './Service';

interface ShippingAttrs {
    service: ServiceDocument;
    shippingType: String;
    width: Number;
    height: Number;
    length: Number;
    weight: Number;
    price: Number;
    shippingId: String;
}

export interface ShippingDocument extends mongoose.Document {
    id: string;
    service: ServiceDocument;
    shippingType: String;
    width: Number;
    height: Number;
    length: Number;
    weight: Number;
    price: Number;
    shippingId: String;
    version: number;
}

interface ShippingModel extends mongoose.Model<ShippingDocument> {
    build(attrs: ShippingAttrs): ShippingDocument;
}

const ShippingSchema = new Schema({
    service: {
        type: Schema.Types.ObjectId,
        ref: Service,
        required: true
    },
    shippingType: {
        type: String,
        required: true
    },
    weight: {
        type: Number,
        require: true
    },
    width: {
        type: Number,
        require: true
    },
    height: {
        type: Number,
        require: true
    },
    length: {
        type: Number,
        require: true
    },
    price: {
        type: Number,
        required: true,
    },
    shippingId: {
        type: String,
        required: true
    }
}, {
    timestamps: {createdAt: 'created_at', updatedAt: 'updated_at'},
    toJSON: {
        transform(_doc, ret){
            ret.id = ret._id;
            delete ret._id;
        }
    }
});

ShippingSchema.statics.build = (attrs: ShippingAttrs) => {
    return new Shipping({...attrs}); 
};

const Shipping = model<ShippingDocument, ShippingModel>('Shipping', ShippingSchema);

export { Shipping };