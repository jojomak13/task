## Routes List

### Users

|Name|Method|Path|
|---|---|---|
|Users|**GET**|`/api/v1/users`|

## Seeder Command
All you need is to run `php artisan seed:transaction` and you will see the magic 🪄

## Running The App [Hard way]
1. run `git clone https://github.com/jojomak13/task`
2. run `cd task`
3. run `composer install`
4. run `cp .env.example .env`
5. run `php artisan key:generate`
6. run `php artisan migrate:fresh --seed`
7. run `php artisan serve`


## Running The App 🚀 🚨 [Easy Way]
1. run `git clone https://github.com/jojomak13/task`
2. run `cd task`
3. run `sh ./install.sh`
4. run `php artisan serve`
