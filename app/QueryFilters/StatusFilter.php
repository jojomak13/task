<?php

namespace App\QueryFilters;

class StatusFilter extends Filter
{
    protected function filterName(): string
    {
        return 'statusCode';
    }

    protected function applyFilter($builder)
    {
        return $builder->where('status', request($this->filterName()));
    }
}
