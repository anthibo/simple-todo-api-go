
# Simple Todo API with golang and gin framework

A simple crud functionalities of a todo API
 


## API Reference

#### Get all items

```http
  GET localhost:9090/todos/
```
#### Get item

```http
  GET /api/todos/${id}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `string` | **Required**. Id of item to fetch |

#### add item
```http
  POST /api/todos
```
body schema : 
```json
   {
        "id": "1",
        "item": "Clean Room",
        "completed": false
    }
```
#### toggle item
```http
  PATCH /api/todos/${id}
```
| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `string` | **Required**. Id of item to fetch |


## Upcoming features:

- add users to the system provided by authentication mechanism
- persist items to database
- implement a simple ui to show items