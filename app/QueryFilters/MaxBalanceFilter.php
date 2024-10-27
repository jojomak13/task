<?php

namespace App\QueryFilters;

class MaxBalanceFilter extends Filter
{
    protected function filterName(): string
    {
        return 'balanceMax';
    }

    protected function applyFilter($builder)
    {
        return $builder->where('amount', '<=', request($this->filterName()));
    }
}
