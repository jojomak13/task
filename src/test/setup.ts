import { MongoMemoryServer } from 'mongodb-memory-server';
import mongoose from 'mongoose';
import { ServiceSeeder } from '../seeders/ServiceSeeder';

let mongo: any;

beforeAll(async () => {
    mongo = await MongoMemoryServer.create();
    const mongoUri = mongo.getUri();

    await mongoose.connect(mongoUri);
});

beforeEach(async () => {
    const collections = await mongoose.connection.db.collections();

    for (let collection of collections) {
        await collection.deleteMany({});
    }

    await ServiceSeeder();
});

afterAll(async () => {
    await mongo.stop();
    await mongoose.connection.close();
});
