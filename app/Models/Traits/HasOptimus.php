<?php

namespace App\Models\Traits;

use Jenssegers\Optimus\Optimus;

trait HasOptimus
{
    public function getKey()
    {
        return $this->id;
    }

    public function getRouteKey()
    {
        return app(Optimus::class)->encode($this->id);
    }

    public function resolveRouteBindingQuery($query, $value, $field = null)
    {
        if (ctype_digit($value) || is_int($value)) {
            $value = app(Optimus::class)->decode($value);
        }
        
        return $query->where($field ?? $this->getRouteKeyName(), $value);
    }
}
