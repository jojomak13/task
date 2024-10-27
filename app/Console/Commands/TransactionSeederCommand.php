<?php

namespace App\Console\Commands;

use Exception;
use App\Models\Provider;
use Illuminate\Support\Str;
use Illuminate\Console\Command;
use App\Services\JsonSeederService;
use Illuminate\Support\Facades\Storage;

use function Laravel\Prompts\error;
use function Laravel\Prompts\select;
use function Laravel\Prompts\text;

class TransactionSeederCommand extends Command
{
    /**
     * The name and signature of the console command.
     *
     * @var string
     */
    protected $signature = 'seed:transaction';

    /**
     * The console command description.
     *
     * @var string
     */
    protected $description = 'Seed transactions from json file';

    /**
     * Execute the console command.
     */
    public function handle()
    {
        $handle = Storage::disk('root')->readStream('/Users/joseph/Sites/task/data.json');

        $filePath = text(
            label: 'File Path',
            placeholder: 'E.g. /data/transactions.json',
            hint: 'It should be an absolute path',
            required: true,
            validate: fn(string $path) => match (true) {
                !Str::endsWith($path, '.json') => 'Only json files are acceptable',
                !Storage::disk('root')->fileExists($path) => 'This is not a valid file path or not exists',
                default => null
            }
        );

        $provider = select(
            label: 'Provider',
            options: Provider::query()->pluck('name', 'id'),
            required: true
        );

        try {
            JsonSeederService::from($filePath)
                ->setProvider(Provider::find($provider))
                ->seed();
        } catch (Exception $e) {
            error($e->getMessage());
        }
    }
}
