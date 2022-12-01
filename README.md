# Shipping Service

## Run The App 🚀
1. add in your `env` variable `MONGO_URI` variable with the connection string eg. 
```
export MONGO_URI=mongodb://localhost:27017/shipping
``` 
2. run `npm install`
3. run `npm run dev`

## Create new shipping request
the only endpoint that exists is `localhost:8080/api/shipping`
the request body like this one
```
{
    "service": "ups",
    "shippingType": "UPSExpress",
    "width": 45.5,
    "height": 20.50,
    "length": 2,
    "weight": 50
}
```
