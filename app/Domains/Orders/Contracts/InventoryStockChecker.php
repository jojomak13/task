<?php

namespace App\Domains\Orders\Contracts;

interface InventoryStockChecker
{
    /**
     * Check if a specific quantity of a product is available.
     */
    public function isAvailable(int $productId, int $quantity): bool;
}
