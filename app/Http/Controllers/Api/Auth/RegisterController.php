<?php

namespace App\Http\Controllers\Api\Auth;

use App\Models\User;
use App\Http\Controllers\Controller;
use App\Http\Requests\Api\Auth\RegisterRequest;
use App\Http\Resources\UserResource;
use App\Http\Traits\HasResponse;

class RegisterController extends Controller
{
    use HasResponse;

    /**
     * Handle the incoming request.
     */
    public function __invoke(RegisterRequest $request)
    {
        $data = $request->validated();

        unset($data['password']);

        $user = User::create([
            'password' => BCrypt($request->input('password')),
            'role' => User::USER_TYPE,
        ] + $data);

        return $this->success(data: new UserResource($user));
    }
}