<?php

namespace Tests\Feature;

use App\Models\Category;
use App\Models\Product;
use App\Models\User;
use Illuminate\Foundation\Testing\RefreshDatabase;
use Laravel\Passport\Passport;
use Tests\TestCase;

class ProductControllerTest extends TestCase
{
    use RefreshDatabase;

    /** @test */
    public function it_return_with_products_with_valid_struture(): void
    {
        Category::factory(2)->create();
        Product::factory(10)->create();

        $response = $this->getJson(route('products.index'))
            ->assertStatus(200)
            ->assertJsonStructure([
                'status',
                'message',
                'data'
            ])
            ->json('data');

        $this->assertIsArray($response);
        $this->assertArrayHasKey('id', $response[0]);
        $this->assertArrayHasKey('name', $response[0]);
        $this->assertArrayHasKey('description', $response[0]);
        $this->assertArrayHasKey('price', $response[0]);
        $this->assertArrayHasKey('rating', $response[0]);
        $this->assertArrayHasKey('category', $response[0]);
        $this->assertArrayHasKey('id', $response[0]['category']);
        $this->assertArrayHasKey('name', $response[0]['category']);
    }


    /** @test */
    public function it_return_with_product_correctly(): void
    {
        Category::factory()->create();
        $product = Product::factory()->create();

        $this->getJson(route('products.show', $product->getRouteKey()))
            ->assertStatus(200)
            ->assertJsonStructure([
                'status',
                'message',
                'data'
            ]);
    }

    /** @test */
    public function it_return_with_not_found_if_product_id_not_encoded(): void
    {
        Category::factory()->create();
        $product = Product::factory()->create();

        $this->getJson(route('products.show', $product->getKey()))
            ->assertStatus(404);
    }

    /** @test */
    public function it_prevent_normal_user_from_creating_product(): void
    {
        Passport::actingAs(User::factory()->create([
            'role' => User::USER_TYPE
        ]));

        $category = Category::factory()->create();

        $this->postJson(route('products.store'), [
            'name' => 'product name',
            'description' => 'product description',
            'price' => 100,
            'category_id' => $category->getRouteKey(),
        ])
            ->assertStatus(403);
    }

    /** @test */
    public function it_can_create_product_if_is_admin(): void
    {
        Passport::actingAs(User::factory()->create([
            'role' => User::ADMIN_TYPE
        ]));

        $category = Category::factory()->create();

        $this->postJson(route('products.store'), [
            'name' => 'product name',
            'description' => 'product description',
            'price' => 100,
            'category_id' => $category->getRouteKey(),
        ])
            ->assertStatus(200);
    }

    /** @test */
    public function it_prevent_normal_user_from_updating_product(): void
    {
        Passport::actingAs(User::factory()->create([
            'role' => User::USER_TYPE
        ]));

        $category = Category::factory()->create();
        $product = Product::factory()->create();

        $this->patchJson(route('products.update', $product->getRouteKey()), [
            'name' => 'product name',
            'description' => 'product description',
            'price' => 100,
            'category_id' => $category->getRouteKey(),
        ])
            ->assertStatus(403);
    }

    /** @test */
    public function it_can_update_product_if_is_admin(): void
    {
        Passport::actingAs(User::factory()->create([
            'role' => User::ADMIN_TYPE
        ]));

        $category = Category::factory()->create();
        $product = Product::factory()->create();

        $this->patchJson(route('products.update', $product->getRouteKey()), [
            'name' => 'product name',
            'description' => 'product description',
            'price' => 100,
            'category_id' => $category->getRouteKey(),
        ])
            ->assertStatus(200);
    }

    /** @test */
    public function it_prevent_normal_user_from_deleting_product(): void
    {
        Passport::actingAs(User::factory()->create([
            'role' => User::USER_TYPE
        ]));

        Category::factory()->create();
        $product = Product::factory()->create();

        $this->deleteJson(route('products.destroy', $product->getRouteKey()))
            ->assertStatus(403);
    }

    /** @test */
    public function it_allow_admin_to_delete_product(): void
    {
        Passport::actingAs(User::factory()->create([
            'role' => User::ADMIN_TYPE
        ]));

        Category::factory()->create();
        $product = Product::factory()->create();

        $this->deleteJson(route('products.destroy', $product->getRouteKey()))
            ->assertStatus(200);
    }
}
