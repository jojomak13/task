<?php

namespace App\Domains\Orders\Actions;

use App\Domains\Orders\Contracts\InventoryStockChecker;
use App\Domains\Orders\DTOs\OrderData;
use InvalidArgumentException;

class CreateOrderAction
{
    public function __construct(
        protected InventoryStockChecker $inventory
    ) {}

    public function handle(OrderData $orderData)
    {
        foreach ($orderData->items as $item) {
            $isAvailable = $this->inventory->isAvailable($item->id, $item->quantity);

            if (!$isAvailable) {
                throw new InvalidArgumentException("Product {$item->id} insufficient quantity");
            }

            // create order and save into database
        }
    }
}
