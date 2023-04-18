
# Sprint 4
## Work Completed
- The UI has had a major overhaul, which makes the website look more professional than it previously was.
- A login and register page was added to allow for user authentication, meaning that users can now register, login and logout from the website.
- All functions from the backend API were implemented in the frontend. This means that the frontend has been successfully completed.
## Unit tests and Cypress tests for frontend
- We previously had a Cypress test to check if routing works and a unit test that checks if the "addThread()" function works properly.
- We now have Cypress tests that do the following:
    - A test for the login page to make sure all the regex implemented on it works properly. It checks if it can detect and incorrect username and incorrect password, then properly logs in.
    - A test that goes through most of the website's functionality. This test logs into the website, adds a thread, adds a reply to that thread, edits the thread, deletes the thread and logs out. This test is meant to show the functionality of the website when logged in, with everything working properly.
- End to end testing is very important for this website because logins are tracked by using cookies, so doing these tests in cypress allows for proper functioning.
## Unit tests for backend
- On the backend, we have tests for these functions:
    - getThreads
    - postThreads
    - postReply
    - allThreads
    - threadByID
    - addReply
    - register
    - login
    - userByName
    - userByID
    - putThread
    - deleteThreadByID
- No additional tests were added because the backend was mostly completed prior to this sprint. A majority of the changes made for this sprint were on the frontend to allow the frontend to access all backend functionality.

## Updated Backend API Documentation
- Getting Threads Data
  - [Show threads](#1-show-threads) : `GET /backend/threads`
  - [Add threads](#2-add-threads): `POST /backend/threads`
  - [Show thread by id](#3-show-thread-by-id) : `GET /backend/threads/:id`
  - [Add replies](#4-add-replies) : `POST /backend/threads/:id`
  - [Delete thread](#5-Delete-thread) : `DELETE /backend/threads/:id`
  - [Update thread](#6-Update-thread) : `PUT /backend/threads/:id`
- User Authentication
  - [Register](#7-register) : `POST /backend/users/register`
  - [Login](#8-login) : `POST /backend/users/login`
  - [Current user](#9-current-user) : `GET /backend/users/user`
  - [Logout](#10-logout) : `POST /backend/users/logout`


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
    "id": 3,
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

## 5. Delete thread
Deletes the thread with the id that matches the input along with all of its replies

**URL** : `/backend/threads/:id`

**Method**: `DELETE`

**URL Parameters** : `id=[integer]`

### Successful Response:
**Code** : `200 OK`

**Content example** : Deleting the thread with the id 1.

Initial:
```json
[
    {
        "id": 1,
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
                "replypost": 1
            }
        ]
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

Result:
```json
[
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

## 6. Update thread
Updates data in the existing thread with the id matching the input

**URL** : `/backend/threads/:id`

**Method**: `PUT`

**URL Parameters** : `id=[integer]`

### Successful Response:
**Code** : `200 OK`

**Content example** : Updating the body of the thread with id 3

Initial:
```json
{
    "id": 3,
    "username": "TheRealGogle",
    "title": "How do I write a database in golang? ",
    "body": "I want to try learning golang but I don't know where to start. I could use some help",
    "time": "2/8/2023, 3:14:53 AM",
    "replies": null
}
```
Result:
```json
{
    "id": 3,
    "username": "TheRealGogle",
    "title": "How do I write a database in golang? ",
    "body": "Thanks for all your help! I can now write databases with ease.",
    "time": "2/8/2023, 3:14:53 AM",
    "replies": null
}
```

## 7. Register
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

## 8. Login
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

## 9. Current User
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

## 10. Logout
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
