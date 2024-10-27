<?php

namespace App\Http\Resources;

use Illuminate\Http\Request;
use Illuminate\Http\Resources\Json\JsonResource;

class TransactionResource extends JsonResource
{
    /**
     * Transform the resource into an array.
     *
     * @return array<string, mixed>
     */
    public function toArray(Request $request): array
    {
        return [
            'id'            => $this->id,
            'email'         => $this->email,
            'amount'        => $this->amount,
            'currency'      => $this->currency,
            'status'        => $this->status->name,
            'created_at'    => $this->creation_date->toDateString(),
        ];
    }
}
