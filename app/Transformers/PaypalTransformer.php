<?php

namespace App\Transformers;

use App\Types\TransactionStatusType;
use Flugg\Responder\Transformers\Transformer;

class PaypalTransformer extends Transformer
{
    /**
     * @param  object $record
     * @return array
     */
    public function transform($record)
    {
        return [
            'email'         => $record->parentEmail,
            'amount'        => $record->parentAmount,
            'currency'      => $record->Currency,
            'status'        => $this->getStatus($record->statusCode),
            'reference_id'  => $record->parentIdentification,
            'creation_date' => $record->registerationDate,
        ];
    }

    private function getStatus(string $status): TransactionStatusType
    {
        return [
            '1' => TransactionStatusType::AUTHORISED,
            '2' => TransactionStatusType::DECLINE,
            '3' => TransactionStatusType::REFUNDED,
        ][$status];
    }
}
