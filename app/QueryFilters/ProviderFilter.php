<?php

namespace App\QueryFilters;

class ProviderFilter extends Filter
{
    protected function filterName(): string
    {
        return 'provider';
    }

    protected function applyFilter($builder)
    {
        return $builder->where('provider_id', request($this->filterName()));
    }
}
