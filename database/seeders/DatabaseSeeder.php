<?php

namespace Database\Seeders;

// use Illuminate\Database\Console\Seeds\WithoutModelEvents;

use App\Models\User;
use Illuminate\Database\Seeder;

class DatabaseSeeder extends Seeder
{
    /**
     * Seed the application's database.
     */
    public function run(): void
    {
        User::factory()->create([
            'name' => 'Test User',
            'email' => 'jojo@test.com',
            'role' => User::USER_TYPE,
        ]);

        User::factory()->create([
            'name' => 'Test Admin',
            'email' => 'admin@test.com',
            'role' => User::ADMIN_TYPE,
        ]);

        $this->call(CategorySeeder::class);

        $this->call(ProductSeeder::class);
    }
}
