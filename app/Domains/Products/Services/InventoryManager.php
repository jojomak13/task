<?php

namespace App\Domains\Products\Services;

use App\Domains\Orders\Contracts\InventoryStockChecker;

class InventoryManager implements InventoryStockChecker
{
    public function isAvailable(int $productId, int $quantity): bool
    {
        throw new \Exception('Not implemented');
    }
}
