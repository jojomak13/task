<?php

namespace App\QueryFilters;

class MinBalanceFilter extends Filter
{
    protected function filterName(): string
    {
        return 'balanceMin';
    }

    protected function applyFilter($builder)
    {
        return $builder->where('amount', '>=', request($this->filterName()));
    }
}
