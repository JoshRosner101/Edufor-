# Sprint 3
## Work Completed
- A mock database was implemented on the backend for unit testing, which allows us to test our functions without clogging the main database
- User authentication was added to the backend. The backend now supports registering users, logging in, logging out, and checking what user is logged in. This was achieved by using JSON Web Tokens and cookies to securely keep track of who is logged in, and encrypting user passwords on the backend using bcrypt. This functionality still needs to be integrated with the frontend.
- [insert work here]
## Unit tests and Cypress tests for frontend
- [insert tests here]
## Unit tests for backend
- On the backend, tests were made for these functions:
    - [insert tests here]

## Updated Backend API Documentation
- Getting Threads Data
  - [Show threads](#1-show-threads) : `GET /backend/threads`
  - [Add threads](#2-add-threads): `POST /backend/threads`
  - [Show thread by id](#3-show-thread-by-id) : `GET /backend/threads/:id`
  - [Add replies](#4-add-replies) : `POST /backend/threads/:id`
- User Authentication
  - [Register](#5-register) : `POST /backend/users/register`
  - [Login](#6-login) : `POST /backend/users/login`
  - [Current user](#7-current-user) : `GET /backend/users/user`
  - [Logout](#8-logout) : `POST /backend/users/logout`


## 1. Show threads
Gets all of the threads currently stored in the database.

**URL** : `/backend/threads`

**Method** : `GET`

### Successful Response:
**Code** : `200 OK`

**Content example** : For a database with only two threads stored, each of which having no replies.
```json
[
    {
        "id": 1,
        "username": "poster27",
        "title": "I need some help with a project",
        "body": "I've been working on a piece of software and could use a helping hand.",
        "time": "2/7/2023, 1:43:27 PM",
        "replies": null
    },
    {
        "id": 2,
        "username": "TheRealGogle",
        "title": "How do I write a database in golang? ",
        "body": "I want to try learning golang but I don't know where to start. I could use some help",
        "time": "2/8/2023, 3:14:53 AM",
        "replies": null
    }
]
```

## 2. Add threads
Adds a new thread with a unique ID number to the database.

**URL** : `/backend/threads`

**Method** : `POST`

**Data example** :

```json
{
    "username": "someRandomName",
    "title": "Helpful Angular Tips ",
    "body": "I've been working on software for a long time and I have a few helpful pointers as to how you could use angular as the main interface for your next software project.",
    "time": "2/9/2023, 1:37:11 PM",
}
```

### Successful Response:
**Code** : `201 Created`

**Content example** : Creating a new thread with the id 3.
```json
{
    "id": 2,
    "username": "TheRealGogle",
    "title": "How do I write a database in golang? ",
    "body": "I want to try learning golang but I don't know where to start. I could use some help",
    "time": "2/8/2023, 3:14:53 AM",
    "replies": null
}
```

## 3. Show thread by id
Gets a thread based on its id.

**URL** : `/backend/threads/:id`

**Method**: `GET`

**URL Parameters** : `id=[integer]`

### Successful Response:
**Code** : `200 OK`

**Content example** : For a thread with the id 4, that has one reply

```json
{
    "id": 4,
    "username": "poster23",
    "title": "Golang Problems",
    "body": "I need help!",
    "time": "2/9/2023, 4:03:39 PM",
    "replies": [
        {
            "replyid": 1,
            "username": "anotherNewUser",
            "body": "I could probably help you out!",
            "time": "2/15/2023, 11:37:28 AM",
            "replypost": 4
        }
    ]
}
```
### Error Response
**Code** : `400 Bad Request`

**Content example** : A thread with an id 5 that does not exist in the database.

```json
{
    "error": "threadByID 5: no such thread"
}
```


## 4. Add replies
Adds a new reply based on the id of the post.

**URL** : `/backend/threads/:id`

**Method**: `POST`

**URL Parameters** : `id=[integer]`

**Data example** :

```json
{
    "username": "anotherRandomUser",
    "body": "I could probably help!",
    "time": "2/13/2023, 2:53:21 PM",
    "replypost": 1
}
```

### Successful Response:
**Code** : `201 Created`

**Content example** : Creating a new reply for a post with the id 1.
```json
{
    "replyid": 5,
    "username": "anotherRandomUser",
    "body": "I could probably help!",
    "time": "2/13/2023, 2:53:21 PM",
    "replypost": 1
}
```

## 5. Register
Adds a new user to the database, automatically encrypting their password in the progress.

**URL** : `/backend/users/register`

**Method**: `POST`

**Data example** :

```json
{
    "name": "randomUser4",
    "password": "hello27"
}
```

### Successful Response:
**Code** : `201 Created`

**Content example** : Registering a new user with name "randomUser4" and password "hello27"
```json
{
    "userid": 9,
    "name": "randomUser4"
}
```

### Error Response
**Code** : `400 Bad Request`

**Content example** : Registering an account with a username that already exists

```json
{
    "error": "username already exists"
}
```

## 6. Login
Returns a cookie with a jwt token that has the user's information upon success (if given the correct username and password). The cookie expires after 24 hours.

**URL** : `/backend/users/login`

**Method**: `POST`

**Data example** :

```json
{
    "name": "randomUser2",
    "password": "hello27"
}
```

### Successful Response:
**Code** : `200 OK`

**Content example** : Logging into an account with username "randomUser2" and password "hello27"
```json
{
    "message": "success"
}
```

### Error Response
**Code** : `404 Not Found`

**Content example** : Logging into an account with a username that doesn't exist

```json
{
    "error": "userByName randomUser32: no such user"
}
```

**Code** : `400 Bad Request`

**Content example** : Logging into an account with an incorrect password

```json
{
    "error": "incorrect password"
}
```

## 7. Current User
Check the jwt token to see what user is logged in

**URL** : `/backend/users/user`

**Method** : `GET`

### Successful Response:
**Code** : `200 OK`

**Content example** : Trying to check user credentials for the currently logged in user
```json
{
    "userid": 9,
    "name": "randomUser4"
}
```

### Error Response
**Code** : `401 Unauthorized`

**Content example** : Trying to check user credentials when no user is logged in

```json
{
    "message": "unauthenticated"
}
```

## 8. Logout
Makes the cookie holding user information expire instantly, effectively logging out the user.

**URL** : `/backend/users/logout`

**Method**: `POST`

### Successful Response:
**Code** : `200 OK`

**Content example** : Logging out of a logged in account
```json
{
    "message": "success"
}
```
