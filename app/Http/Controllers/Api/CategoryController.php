<?php

namespace App\Http\Controllers\Api;

use App\Http\Controllers\Controller;
use App\Http\Resources\CategoryResource;
use App\Http\Traits\HasResponse;
use App\Models\Category;
use Illuminate\Http\Request;

class CategoryController extends Controller
{
    use HasResponse;

    /**
     * Handle the incoming request.
     */
    public function __invoke(Request $request)
    {
        $categories = Category::latest()->paginate(15);

        return $this->success(data: CategoryResource::collection($categories));
    }
}
