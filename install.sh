composer install

cp .env.example .env

php artisan key:generate

php artisan migrate:fresh --seed

php artisan passport:install