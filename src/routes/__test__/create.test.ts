import request from 'supertest';
import app from '../../server';

it('return with 201 if shipping request created', async () => {

    await request(app)
        .post('/api/shipping')
        .send({
            service: 'fedex',
            shippingType: 'fedexGroud',
            width: 12.5,
            height: 12.5,
            length: 12.5,
            weight: 5.0,
        })
        .expect(201)
})

it('returns with 400 with invalid service type', async () => {
    await request(app)
        .post('/api/shipping')
        .send({
            service: 'hello',
            shippingType: 'fedexGroud',
            width: 12.5,
            height: 12.5,
            length: 12.5,
            weight: 5.0,
        })
        .expect(400);
})

it('returns with 400 with invalid shipping type', async () => {
    await request(app)
        .post('/api/shipping')
        .send({
            service: 'ups',
            shippingType: 'welcome',
            width: 12.5,
            height: 12.5,
            length: 12.5,
            weight: 5.0,
        })
        .expect(400);
})

it('returns with 400 with invalid dimensions', async () => {
    await request(app)
        .post('/api/shipping')
        .send({
            service: 'ups',
            shippingType: 'fedexGroud',
            width: -5,
            height: -12,
            length: 7.0,
            weight: 5.0,
        })
        .expect(400);
})

it('returns with 400 with invalid weight', async () => {
    await request(app)
        .post('/api/shipping')
        .send({
            service: 'ups',
            shippingType: 'fedexGroud',
            width: 12.5,
            height: 12.5,
            length: 12.5,
            weight: -5,
        })
        .expect(400);
})

it('returns with 400 with required fields', async () => {
    await request(app)
        .post('/api/shipping')
        .send({})
        .expect(400);
})