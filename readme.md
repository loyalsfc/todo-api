# Todo API Documentation

## Introduction
This API allows users to manage their todo lists. It provides endpoints for retrieving todo lists, fetching a single todo item by ID, adding new todo items, updating existing todo items, and deleting todo items.

## Base URL

## Endpoints
https://todo-api-r0vy.onrender.com/v1

### Get Todo Lists
Retrieves a list of all todo items.

**Request**
GET /todos

**Response**
```json
    [
        {
            "id": 1,
            "title": "Complete project",
            "description": "Finish building the project by next week",
            "is_completed": true,
            "created_at": "2024-03-19 03:24:47.333714"
        },
        {
            "id": 2,
            "title": "Go for a run",
            "description": "Run in the park for 30 minutes",
            "is_completed": true,
            "created_at": "2024-03-19 03:24:47.333714"
        },
        ...
    ]
```

### Get Single Todo
Retrieves a single todo item by ID.

**Request**
GET /todos/{todoId}

**Response**
```json
    {
        "id": 1,
        "title": "Complete project",
        "description": "Finish building the project by next week",
        "is_completed": true,
        "created_at": "2024-03-19 03:24:47.333714"
    }
```

### Add Todo
Adds a new todo item.

**Request**
POST /todo

```json
    {
        "title": "Complete project",
        "description": "Finish building the project by next week",
        "is_completed": true,
    }
```

**Response**
```json
    {
        "id": 1,
        "title": "Complete project",
        "description": "Finish building the project by next week",
        "is_completed": true,
        "created_at": "2024-03-19 03:24:47.333714"
    }
```

### Update Todo
Updates an existing todo item by ID.

**Request**
PUT /todo/{todoId}

```json
    {
        "title": "Complete project",
        "description": "Finish building the project by next week",
        "is_completed": true,
    }
```

**Response**
```json
    {
        "id": 1,
        "title": "Complete project",
        "description": "Finish building the project by next week",
        "is_completed": true,
        "created_at": "2024-03-19 03:24:47.333714"
    }
```

### Delete Todo
Deletes a todo item by ID.

**Request**
DELETE /todo/{todoId}

**Response**
```json
    {
        "message": "Todo Deleted successfully"
    }
```