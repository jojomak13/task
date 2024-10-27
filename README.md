## Routes List

### Auth

|Name|Method|Path|Request Body|
|---|---|---|---|
|Login|**POST**|`/auth/login`|`[email, password, device_name]`|
|Register|**POST**|`/auth/register`|`[name, email, password, device_name]`|
|Logout|**DELETE**|`/auth/logout`|`[]`|
|Profile|**GET**|`/auth/me`|`[]`|


### Categories

|Name|Method|Path|Request Body|
|---|---|---|---|
|List Categoties|**GET**|`/categories`|`[]`|



### Products

|Name|Method|Path|Request Body|
|---|---|---|---|
|List Products|**GET**|`/products`|`[category_id, category_name, name, order_by, direction]`|
|Show Product|**GET**|`/products/{id}`|`[]`|
|Delete Product|**DELETE**|`/products/{id}`|`[]`|
|Create Product|**POST**|`/products`|`[name, description, price, category_id]`|
|Update Product|**POST**|`/products/{id}`|`[name, description, price, category_id]`|


## Running The App [Hard way]
1. run `git clone https://github.com/jojomak13/task`
2. run `cd task`
3. run `composer install`
4. run `cp .env.example .env`
5. run `php artisan key:generate`
6. run `php artisan migrate:fresh --seed`
7. run `php artisan passport:install`
8. run `php artisan serve`


## Running The App 🚀 🚨 [Easy Way]
1. run `git clone https://github.com/jojomak13/task`
2. run `cd task`
3. run `sh ./install.sh`
4. run `php artisan serve`


## Postman Collection
you could find it on the root directory with name `Task.postman_collection.json`

## Testing
run `php artisan test`
