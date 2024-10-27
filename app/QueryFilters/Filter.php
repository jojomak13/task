<?php

namespace App\QueryFilters;

use Closure;

abstract class Filter
{
    protected abstract function applyFilter($builder);

    protected abstract function filterName(): string;

    public function handle($request, Closure $next)
    {
        $builder = $next($request);

        if(!request()->has($this->filterName())) {
            return $builder;
        }

        return $this->applyFilter($builder);
    }

}
