
# Simple Todo API with golang and gin framework

A simple crud functionalities of a todo API
 


## API Reference

#### Get all items

```http
  GET localhost:9090/todos/
```
#### Get item

```http
  GET localhost:9090/todos/${id}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `number` | **Required**. Id of item to fetch |

#### add item
```http
  POST localhost:9090/todos
```
body schema : 
```json
   {
        "item": "Clean Room",
    }
```
#### toggle item
```http
  PATCH localhost:9090/todos/${id}
```
| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `number` | **Required**. Id of item to fetch |


## Upcoming features:
 - add users to the system provided by authentication mechanism
 - add unit testing
 - benchmark and load test the application
