<?php

namespace Database\Seeders;

use App\Models\Provider;
use App\Transformers;
use Illuminate\Database\Seeder;

class ProviderSeeder extends Seeder
{
    /**
     * Run the database seeds.
     */
    public function run(): void
    {
        Provider::create([
            'key' => 'paypal',
            'name' => 'Paypal',
            'transformer' => Transformers\PaypalTransformer::class,
        ]);

        Provider::create([
            'key' => 'stripe',
            'name' => 'Stripe',
            'transformer' => Transformers\StripeTransformer::class,
        ]);
    }
}
