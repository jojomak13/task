<?php

namespace App\Domains\Orders\DTOs;

use Illuminate\Http\Request;

readonly class OrderData
{
    public function __construct(
        public int $userId,
        public array $items, // Could be an array of ItemDTOs
        public string $paymentMethod,
        public ?string $promoCode = null,
    ) {}

    public static function fromRequest(Request $request): self
    {
        return new self(
            userId: $request->user()->id,
            items: $request->validated('items'),
            paymentMethod: $request->validated('payment_method'),
            promoCode: $request->validated('promo_code'),
        );
    }
}
