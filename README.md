# Project: Online Journaling Application
A platform that allows users to create, store, and manage personal journal entries. Users can write about their thoughts, experiences, and reflections in a digital format. Implementing such an application using microservices architecture involves breaking down the various components and functionalities of the app into separate, independent services.

Microservices may be relevant and necessary for such an application because of the following concepts:

1. **Scalability**: the application can experience varying levels of user activity. Microservices allow scaling specific components independently. For instance, the user registration/authentication service can be scaled separately from the journaling service to handle spikes in user activities.
2. **Independent Development**: the developer can update certain components of the application independently, such as add features to user authentication, or enhance journaling functionality without affecting other parts of the application.
3. **Technology Diversity**: Different components of the application may require different technologies or databases. Microservices offer the flexibility to choose the best technology stack for each service.
4. **Fault Isolation**: if one service of the application fails or encounters issues, it doesn't necessarily bring down the entire system. Users can still access other parts of the app, such as reading existing journal entries, while the problematic service is fixed.
5. **Easier Maintenance**: Maintenance becomes more manageable when you can focus on individual services. You can update, patch, or troubleshoot one service without disrupting the entire application.
    

Real-world examples of well-known projects that employ microservices include:

1. **Facebook**: Facebook utilizes a microservices architecture to manage various aspects of its platform. Different services handle user authentication, news feed generation, messaging, and more. This enables Facebook to scale and maintain its massive user base efficiently.
2. **Netflix**: Netflix is another prime example. It employs microservices to manage its extensive content library, user recommendations, and streaming services. This architecture enables Netflix to adapt quickly to user preferences and handle a global user base.
3. **Uber**: Uber relies on microservices to power its ride-sharing platform. Different services handle driver matching, ride tracking, payments, and more. This architecture has allowed Uber to expand rapidly and adapt to diverse market needs.
    

## **Application Services:**

1. User Authentication (JWT Middleware)
2. User Creation & Maintenance
3. Journal Entries Creation & Maintenance
4. User Accounts View
5. Journal Entries View
    

<img src="https://content.pstmn.io/d09cc3bf-c855-4463-94dc-14576beae6a1/bWljcm9zZXJ2aWNlcy1hcmNoaXRlY3R1cmUucG5n" alt="Microservices%20Architecture%20Diagram">

**Technology Stack:** Golang, Python, Gin, Gorm, gRPC, REST API

**Deployment and Scaling:** Kubernetes

## Endpoints

### Register User
POST request with the purpose of user registration. It sends a JSON payload containing a username and password for a user account creation.
#### Method: POST
>```
>http://localhost:8000/users/register
>```
#### Headers

|Content-Type|Value|
|---|---|
|Content-Type|application/json|
|Accept|application/json|

#### Body (**raw**)

```json
{
    "username":"test0",
    "password":"test0"
}
```

#### Response (**raw**)

```json
{
    "message":"User successfully created"
}
```

#### Authentication: no auth

### Login User
POST request designed for user authentication. It sends a JSON payload with a username and password to the server for login purposes.
#### Method: POST
>```
>http://localhost:8000/users/login
>```
#### Headers

|Content-Type|Value|
|---|---|
|Content-Type|application/json|
|Accept|application/json|

#### Body (**raw**)

```json
{
    "username":"test0",
    "password":"test0"
}
```

#### Response (**raw**)

```json
{
    "access-token":"{{accessToken}}"
}
```

#### Authentication: no auth

### Delete User
DELETE request intended to delete a user account, including authentication using a bearer token in the "Authorization" header, and the request body provides the user's password as a security measure to confirm the user's identity.
#### Method: DELETE
>```
>http://localhost:8000/users/delete/test0
>```
#### Headers

|Content-Type|Value|
|---|---|
|Content-Type|application/json|
|Accept|application/json|

#### Body (**raw**)

```json
{
    "password":"test0"
}
```

#### Response (**raw**)

```json
{
    "message":"User successfully deleted"
}
```

#### Authentication: bearer token

### Create Journal Entry
POST request to create a new journal entry for a specific user. It sends a JSON payload containing the username, title, and description of the new entry, along with authorization using a bearer token.
#### Method: POST
>```
>http://localhost:8000/entries/create
>```
#### Headers

|Content-Type|Value|
|---|---|
|Content-Type|application/json|
|Accept|application/json|

#### Body (**raw**)

```json
{
    "username":"test0",
    "title":"aaa",
    "description":"aaa"
}
```

#### Response (**raw**)

```json
{
    "message":"Journal entry successfully added"
}
```

#### Authentication: bearer token

### Get User's Journal Entries
GET request to retrieve a user's journal entries.
#### Method: GET
>```
>http://localhost:8000/entries/test0
>```
#### Headers

|Content-Type|Value|
|---|---|
|Accept|application/json|

#### Response (**raw**)

```json
{
    "entries": {
        "1": {
            "title":"aaa",
            "description":"aaa"
        }
    }
}
```

#### Authentication: bearer token

### Get User's Journal Entry
GET request to retrieve a user's specific journal entry.
#### Method: GET
>```
>http://localhost:8000/entries/test0/1
>```
#### Headers

|Content-Type|Value|
|---|---|
|Accept|application/json|

#### Response (**raw**)

```json
{
    "title":"aaa",
    "description":"aaa"
}
```

#### Authentication: bearer token

### Delete Journal Entry
DELETE request intended to delete a user's journal entry, including authentication using a bearer token in the "Authorization" header, and the request body provides the user's username to confirm the user.
#### Method: DELETE
>```
>http://localhost:8000/entries/delete
>```
#### Headers

|Content-Type|Value|
|---|---|
|Content-Type|application/json|
|Accept|application/json|

#### Body (**raw**)

```json
{
    "username":"test0",
    "title": "aaa"
}
```

#### Response (**raw**)

```json
{
    "message":"Journal entry successfully deleted"
}
```

#### Authentication: bearer token
