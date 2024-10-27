<?php

namespace App\Models;

use App\Types\TransactionStatusType;
use Illuminate\Database\Eloquent\Model;

class Transaction extends Model
{
    protected $guarded = ['id'];

    protected $casts = [
        'creation_date' => 'datetime',
        'status'        => TransactionStatusType::class,
    ];
}
