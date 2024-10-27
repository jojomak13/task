<?php

namespace App\Http\Controllers\Api;

use App\Http\Controllers\Controller;
use App\Http\Resources\TransactionResource;
use App\Models\Transaction;
use App\QueryFilters;
use Illuminate\Http\Request;
use Illuminate\Pipeline\Pipeline;

class TransactionController extends Controller
{
    public function __invoke(Request $request)
    {
        $pipe = app(Pipeline::class)
            ->send(Transaction::query())
            ->through([
                QueryFilters\ProviderFilter::class,
                QueryFilters\StatusFilter::class,
                QueryFilters\MinBalanceFilter::class,
                QueryFilters\MaxBalanceFilter::class,
            ])
            ->thenReturn();

        return TransactionResource::collection($pipe->simplePaginate());
    }
}
