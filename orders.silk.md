# Orders resources

## GET /orders

Get all of orders items
* Content-Type: `application/json`
* Accept: `application/json`

===

### Example response

* Status: `200`
* Content-Type: `application/json; charset=utf-8`

```
[{"id":"57e3f8e85e812900039d4647","menu_id":"58059ff853490a001277d862","menu_name":"ต้มยำกุ้ง","amount":1,"price":99,"seat":"05","status":"Waiting"}]
```

## POST /orders

* Content-Type: `application/json`
* Accept: `application/json`

```json
{
  "menu_id": "58059ff853490a001277d862",
  "menu_name": "ต้มยำกุ้ง",
  "amount": 1,
  "price": 99.5,
  "seat": "05",
  "status": "Waiting"
}
```
===

### Example response
* Status: `201`

## PUT /orders/57e3f8e85e812900039d4647
* Content-Type: `application/json`
* Accept: `application/json`

```json
{
    "status": "Served"
}
```

===

### Example response
* Status: `200`

## DELETE /orders/57e3f8e85e812900039d4647
* Content-Type: `application/json`
* Accept: `application/json`

===

### Example response
* Status: `200`


#PATCH /users/:id(.:format) users#update
