<?php

use Illuminate\Http\Request;
use Illuminate\Support\Facades\Route;
use App\Http\Controllers\Api;

/*
|--------------------------------------------------------------------------
| API Routes
|--------------------------------------------------------------------------
|
| Here is where you can register API routes for your application. These
| routes are loaded by the RouteServiceProvider and all of them will
| be assigned to the "api" middleware group. Make something great!
|
*/

// Auth
Route::prefix('auth')->name('auth.')->group(function () {
    Route::post('login', [Api\Auth\LoginController::class, 'login'])->name('login');
    Route::post('register', Api\Auth\RegisterController::class)->name('register');
});


Route::middleware('auth:api')->group(function () {
    Route::delete('auth/logout', [Api\Auth\LoginController::class, 'logout']);

    Route::get('/me', function (Request $request) {
        return $request->user();
    })->name('profile');

    Route::apiResource('products', Api\ProductController::class)->except(['show', 'index']);
});

Route::get('categories', Api\CategoryController::class)->name('categories.index');

Route::get('products', [Api\ProductController::class, 'index'])->name('products.index');
Route::get('products/{product}', [Api\ProductController::class, 'show'])->name('products.shows');
