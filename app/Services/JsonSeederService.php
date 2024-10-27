<?php

namespace App\Services;

use Exception;
use App\Models\Provider;
use JsonCollectionParser\Parser;

class JsonSeederService
{
    private Parser $parser;
    private Provider $provider;
    private int $chunkSize = 100;

    private function __construct(private string $filePath)
    {
        $this->parser = new Parser();
    }

    public static function from(string $filePath): self
    {
        return new self($filePath);
    }

    public function setProvider(Provider $provider): self
    {
        $this->provider = $provider;

        return $this;
    }

    public function setChunkSize(int $size): self
    {
        $this->chunkSize = $size;

        return $this;
    }

    public function seed()
    {
        if (!isset($this->provider)) {
            throw new Exception('Provider not configured');
        }

        $this->parser->chunkAsObjects(
            $this->filePath,
            fn(array $data) => $this->handle($data),
            $this->chunkSize
        );
    }

    private function handle(array $items)
    {
        $transactions = transformation($items, $this->provider->transformer)->transform();

        $this->provider->transactions()->createMany($transactions);
    }
}
