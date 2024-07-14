<?php

namespace App\Models;

use App\Models\Traits\HasOptimus;
use Illuminate\Database\Eloquent\Factories\HasFactory;
use Illuminate\Database\Eloquent\Model;

class Category extends Model
{
    use HasFactory, HasOptimus;

    protected $guarded = ['id'];
}
