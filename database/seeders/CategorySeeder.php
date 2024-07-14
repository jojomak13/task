<?php

namespace Database\Seeders;

use App\Models\Category;
use Illuminate\Database\Console\Seeds\WithoutModelEvents;
use Illuminate\Database\Seeder;

class CategorySeeder extends Seeder
{
    /**
     * Run the database seeds.
     */
    public function run(): void
    {
        $categories = [
            ['name' => 'Meat'],
            ['name' => 'Fresh produce'],
            ['name' => 'Dairy'],
            ['name' => 'Bakery'],
            ['name' => 'Pantry'],
            ['name' => 'Beverages'],
            ['name' => 'Frozen foods'],
            ['name' => 'Household'],
        ];

        foreach($categories as $cat) {
            Category::create($cat);
        }
    }
}
