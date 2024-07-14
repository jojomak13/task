<?php

namespace App\Http\Controllers\Api;

use App\Http\Controllers\Controller;
use App\Http\Requests\Api\Products\CreateProductRequest;
use App\Http\Requests\Api\Products\UpdateProductRequest;
use App\Http\Resources\ProductResource;
use App\Http\Traits\HasResponse;
use App\Models\Product;
use Illuminate\Http\Request;
use Jenssegers\Optimus\Optimus;

class ProductController extends Controller
{
    use HasResponse;

    /**
     * Display a listing of the resource.
     */
    public function index(Request $request)
    {
        $products = Product::query()
            ->when($request->get('name'), function ($query, $value) {
                $query->where('name', 'like', '%' . $value . '%');
            })
            ->when($request->get('category_id'), function ($query, $value) {
                $query->where('category_id', app(Optimus::class)->decode($value));
            })
            ->when($request->get('category_name'), function ($query, $value) {
                $query->whereHas('category', function() use($value)  {
                    $this->where('name', 'like', '%' . $value . '%');
                });
            })
            ->orderBy($request->get('order_by', 'id') , request()->get('direction', 'asc'))
            ->paginate(15);

        return $this->success(data: ProductResource::collection($products));
    }

    /**
     * Store a newly created resource in storage.
     */
    public function store(CreateProductRequest $request)
    {
        $this->authorize('store');

        $product = Product::create($request->validated());

        return $this->success(message: __('Product created successfully'), data: new ProductResource($product));
    }

    /**
     * Display the specified resource.
     */
    public function show(Product $product)
    {
        return $this->success(data: new ProductResource($product));
    }

    /**
     * Update the specified resource in storage.
     */
    public function update(UpdateProductRequest $request, Product $product)
    {
        $this->authorize('update', $product);

        $product->update($request->validated());

        return $this->success(message: __('Product updated successfully'), data: new ProductResource($product));
    }

    /**
     * Remove the specified resource from storage.
     */
    public function destroy(Product $product)
    {
        $this->authorize('delete', $product);

        $product->delete();

        return $this->success(message: __('Product deleted successfully'));
    }
}
