# Sprint 2
## Work Completed
- The rewrite of the backend to use a database has been successfully completed. The backend now uses sqlite to create a local database and store post data there instead of memory. This means that data will be saved even when the backend server is shut down and needs to be booted up again.
- An additional table in the backend database was added to store replies, which have a many-to-one relationship with threads.
- The backend and frontend were already integrated, but work has been done to make sure the new database works properly in junction with the front end.
- For the frontend, routing was added, which means that there is now support for multiple pages on the website. A new page was added for each specific thread, so now you can click on certain threads and go to a page that only shows info for that specific thread. A basic login page was also made, but is currently not fully implemented.
## Unit tests and Cypress tests for frontend
- The Cypress test is a basic test to make sure that the routing works (it clicks on a hyperlink and checks if the webpage title is correct
- For the unit test, it checks the function "addThread()" on the main page to see if filling out the form and running the function adds the information to the webpage.
## Unit tests for backend
- On the backend, tests were made for these functions:
    - getThreads
    - postThreads
    - postReply
    - allThreads
    - threadByID
    - addReply
    - getThreadsByID (currently doesn't work)
- These tests are to ensure that the GET and POST functionality of the backend works in junction with the database.
- Currently these tests work on the actual database, but we plan in implementing a mock database to test these on for the next sprint.
## Backend API Documentation
- [Show threads](#1-show-threads) : `GET /backend/threads`
- [Add threads](#2-add-threads): `POST /backend/threads`
- [Show thread by id](#3-show-thread-by-id) : `GET /backend/threads/:id`
- [Add replies](#4-add-replies) : `POST /backend/threads/:id`


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
