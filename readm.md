# Mapup Intersection API


In this API, we can get the intersection point from the requested linestring with the already existing lines

## /token

Using this we should create a token

## /intersection 

We should pass the linestrings as request body, along with `Token` as header in the format of
```
"coordinates": [
    [-96.79512, 32.77823],
    [-96.79469, 32.77832]
]
``` 

We will get the intersecting line index and intersecting points


Added the postman collection for your ref.

To run this code 
```
make build
./mapup
```
