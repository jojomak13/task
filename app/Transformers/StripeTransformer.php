<?php

namespace App\Transformers;

use App\Types\TransactionStatusType;
use Carbon\Carbon;
use Flugg\Responder\Transformers\Transformer;

class StripeTransformer extends Transformer
{
    /**
     * @param  object $record
     * @return array
     */
    public function transform($record)
    {
        return [
            'email'         => $record->email,
            'amount'        => $record->balance,
            'currency'      => $record->currency,
            'status'        => $this->getStatus($record->status),
            'reference_id'  => $record->id,
            'creation_date' => Carbon::createFromFormat('d/m/Y', $record->created_at),
        ];
    }

    private function getStatus(string $status): TransactionStatusType
    {
        return [
            '100' => TransactionStatusType::AUTHORISED,
            '200' => TransactionStatusType::DECLINE,
            '300' => TransactionStatusType::REFUNDED,
        ][$status];
    }
}
