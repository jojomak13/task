<?php

namespace Tests\Feature\Auth;

use Illuminate\Foundation\Testing\RefreshDatabase;
use Tests\TestCase;

class RegisterControllertest extends TestCase
{
    use RefreshDatabase;

    /** @test */
    public function it_register_new_user(): void
    {
        $this->postJson(route('auth.register'), [
            'name' => 'test',
            'email' => 'a@a.com',
            'password' => 'password',
            'password_confirmation' => 'password',
            'device_name' => 'test',
        ])
            ->assertStatus(200)
            ->assertJsonStructure([
                'status',
                'message',
                'data' => [
                    'id',
                    'name',
                    'email',
                    'token'
                ]
            ]);
    }

    /** @test */
    public function it_fails_if_name_not_exists(): void
    {
        $this->postJson(route('auth.register'), [
            'email' => 'a@a.com',
            'password' => 'password',
            'password_confirmation' => 'password',
            'device_name' => 'test',
        ])
            ->assertStatus(422)
            ->assertInvalid(['name']);
    }
    
    /** @test */
    public function it_fails_if_email_not_exists(): void
    {
        $this->postJson(route('auth.register'), [
            'name' => 'test',
            'password' => 'password',
            'password_confirmation' => 'password',
            'device_name' => 'test',
        ])
            ->assertStatus(422)
            ->assertInvalid(['email']);
    }

    /** @test */
    public function it_fails_if_email_not_valid_email(): void
    {
        $this->postJson(route('auth.register'), [
            'name' => 'test',
            'email' => 'invalid-email',
            'password' => 'password',
            'password_confirmation' => 'password',
            'device_name' => 'test',
        ])
            ->assertStatus(422)
            ->assertInvalid(['email']);
    }

    /** @test */
    public function it_fails_if_password_not_exists(): void
    {
        $this->postJson(route('auth.register'), [
            'name' => 'test',
            'email' => 'a@a.com',
            'password_confirmation' => 'password',
            'device_name' => 'test',
        ])
            ->assertStatus(422)
            ->assertInvalid(['password']);
    }

    /** @test */
    public function it_fails_if_password_confirmation_not_match(): void
    {
        $this->postJson(route('auth.register'), [
            'name' => 'test',
            'email' => 'a@a.com',
            'password' => 'password',
            'password_confirmation' => '123',
            'device_name' => 'test',
        ])
            ->assertStatus(422)
            ->assertInvalid(['password']);
    }

    /** @test */
    public function it_fails_if_device_name_not_exists(): void
    {
        $this->postJson(route('auth.register'), [
            'name' => 'test',
            'email' => 'a@a.com',
            'password' => 'password',
            'password_confirmation' => 'password',
        ])
            ->assertStatus(422)
            ->assertInvalid(['device_name']);
    }
}
