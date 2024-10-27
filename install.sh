composer install

cp .env.example .env

touch database/database.sqlite

php artisan key:generate

php artisan migrate:fresh --seed
