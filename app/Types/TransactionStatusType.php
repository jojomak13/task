<?php

namespace App\Types;

enum TransactionStatusType: int
{
    case AUTHORISED     = 1;
    case DECLINE        = 2;
    case REFUNDED       = 3;
}
