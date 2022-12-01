import Joi from 'joi';

const CreateShippingRequest = Joi.object({
    service: Joi.any()
        .valid(...['fedex', 'ups'])
        .required(),
        
    shippingType: Joi.string().required(),

    width: Joi.number().positive().required(),

    height: Joi.number().positive().required(),

    length: Joi.number().positive().required(),

    weight: Joi.number().positive().required(),
});

export { CreateShippingRequest };