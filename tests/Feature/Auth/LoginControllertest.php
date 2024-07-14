<?php

namespace Tests\Feature\Auth;

use App\Models\User;
use Illuminate\Foundation\Testing\RefreshDatabase;
use Tests\TestCase;

class LoginControllertest extends TestCase
{
    use RefreshDatabase;

    /** @test */
    public function it_login_user_correctly(): void
    {
        $user = User::factory()->create();

        $this->postJson(route('auth.login'), [
            'email' => $user->email,
            'password' => 'password',
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
    public function it_fails_if_credintials_not_valid(): void
    {
        $user = User::factory()->create();

        $this->postJson(route('auth.login'), [
            'email' => $user->email,
            'password' => 'invalid-password',
            'device_name' => 'test',
        ])
            ->assertStatus(422)
            ->assertInvalid(['email']);
    }
}
