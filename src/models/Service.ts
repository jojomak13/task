import mongoose, {Schema, model} from 'mongoose';

interface ServiceAttrs {
    name: string;
    code: string;
    unitsType: string;
    types: string[];
}

export interface ServiceDocument extends mongoose.Document {
    id: string;
    name: string;
    code: string;
    unitsType: string;
    types: string[];
    version: number;
}

interface ServiceModel extends mongoose.Model<ServiceDocument> {
    build(attrs: ServiceAttrs): ServiceDocument;
}

const ServiceSchema = new Schema({
    name: {
        type: String,
        required: true
    },
    code: {
        type: String,
        required: true
    },
    unitsType: {
        type: String,
        enum: ['uk', 'us'],
        required: true
    },
    types: {
        type: Array,
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

ServiceSchema.statics.build = (attrs: ServiceAttrs) => {
    return new Service({...attrs}); 
};

const Service = model<ServiceDocument, ServiceModel>('Service', ServiceSchema);

export { Service };