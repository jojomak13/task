<?php

use Illuminate\Database\Migrations\Migration;
use Illuminate\Database\Schema\Blueprint;
use Illuminate\Support\Facades\Schema;

return new class extends Migration
{
    /**
     * Run the migrations.
     */
    public function up(): void
    {
        Schema::create('transactions', function (Blueprint $table) {
            $table->id();
            $table->string('email');
            $table->decimal('amount', 8, 3);
            $table->string('currency');
            $table->tinyInteger('status');
            $table->string('reference_id');
            $table->foreignid('provider_id')
                ->constrained('providers')
                ->nullOnUpdate()
                ->cascadeOnUpdate();
            $table->timestamp('creation_date');
            $table->timestamps();
        });
    }

    /**
     * Reverse the migrations.
     */
    public function down(): void
    {
        Schema::dropIfExists('transactions');
    }
};
