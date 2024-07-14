<?php

namespace App\Http\Traits;

use Illuminate\Http\JsonResponse;
use Illuminate\Contracts\Container\BindingResolutionException;
use Illuminate\Http\Response;

trait HasResponse
{
    /**
     * @param mixed $data 
     * @param mixed $message 
     * @return JsonResponse 
     * @throws BindingResolutionException 
     */
    public function success($data = null, $message = null)
    {
        return response()->json([
            'status' => true,
            'message' => $message,
            'data' => $data,
        ]);
    }

    /**
     * @param mixed $errors 
     * @param mixed $message 
     * @return JsonResponse 
     * @throws BindingResolutionException 
     */
    public function fail($errors = null, $message = null)
    {
        return response()->json([
            'status' => false,
            'message' => $message,
            'errors' => $errors,
        ], Response::HTTP_BAD_REQUEST);
    }
}
