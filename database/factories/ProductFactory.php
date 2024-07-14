<?php

namespace Database\Factories;

use App\Models\Category;
use Illuminate\Database\Eloquent\Factories\Factory;

/**
 * @extends \Illuminate\Database\Eloquent\Factories\Factory<\App\Models\Product>
 */
class ProductFactory extends Factory
{
    /**
     * Define the model's default state.
     *
     * @return array<string, mixed>
     */
    public function definition(): array
    {
        return [
            'name' => fake()->words(asText: true),
            'description' => fake()->words(10, asText: true),
            'price' => fake()->numberBetween(100, 1000),
            'rating' => fake()->numberBetween(0, 5),
            'category_id' => Category::query()->inRandomOrder()->first()->id,
        ];
    }
}
