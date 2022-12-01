import { Service } from "../models/Service";

const ServiceSeeder = async () => {
    await Service.build({
        name: 'FedEx',
        code: 'fedex',
        types: ['fedexAIR', 'fedexGroud'],
        unitsType: 'us'
    }).save()

    await Service.build({
        name: 'UPS',
        code: 'ups',
        types: ['UPSExpress', 'UPS2DAY'],
        unitsType: 'uk'
    }).save()

}

export { ServiceSeeder }