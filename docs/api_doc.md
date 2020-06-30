# API Documentation

## Endpoints without Auth
|URL|Endpoint|Type|
|--|--|--|
|/api/v1/login|Login  |POST|
|/api/v1/register|Register|POST|
|/api/v1/check/userAvailable/{username}|Check if Username is available|GET|
|/api/v1/check/emailAvailable/{email}|Check if Email is available|GET|

## /api/v1/login 
**Method: POST**  
This endpoint is used to authenticate and generate a token for an already registered user.

#### Request
##### Header:
|Key|Value|
|-|-|
|Content-Type|application/json|

##### Body:

 ```
 {
	 "username":"string"
	 "password":"SHA-256 string"
 }
 ```

#### Response
This Endpoint can return following StatusCodes:
| Code | Name |Reason|
|--|--|--|
|  200|OK  |Successful Request|
|400|Bad Request| Malformed Request|
|500|Internal Server Error| Serverside Error



#### Status 200 Success
##### Header: 
|Key|Value|
|-|-|
|Content-Type|application/json|

##### Body:

 ```
 {
	 "token":"Encoded JWT-Token String"
 }
 ```
#### Status 400/500 Failure
##### Header: 
|Key|Value|
|-|-|
|Content-Type|application/json|

##### Body:

 ```
 {
	 "err_message":"string"
 }
 ```

## /api/v1/register
**Method: POST**  
This endpoint is used to register a new user and generate authentification

#### Request
##### Header:
|Key|Value|
|-|-|
|Content-Type|application/json|

##### Body:

 ```
 {
	 "username_check":boolean	//callback of /api/v1/check/userAvailable/{id}
	 "email_check": boolean 		//callback of /api/v1/check/emailAvailable/{id}
	 "username":"string"
	 "email":"string"
	 "fist_name":"string"
	 "last_name":"string"
	 "password":"SHA-256 string"
 }
 ```

#### Response
This Endpoint can return following StatusCodes:
| Code | Name |Reason|
|--|--|--|
|  200|OK  |Successful Request|
|400|Bad Request| Malformed Request|
|409|Conflict|Username/Email already Exists or not Valid
|500|Internal Server Error| Serverside Error



#### Status 200 Success
##### Header: 
|Key|Value|
|-|-|
|Content-Type|application/json|

##### Body:

 ```
 {
	 "token":"JWT-Token String"
 }
 ```

#### Status 409 Conflict
##### Header: 
|Key|Value|
|-|-|
|Content-Type|application/json|

##### Body:

 ```
 {
	 "err_message":"string"
 	"email_check": boolean
 	"username_check": boolean
 }
 ```
 
#### Status 400/500 Failure
##### Header: 
|Key|Value|
|-|-|
|Content-Type|application/json|

##### Body:

 ```
 {
	 "err_message":"string"
 }
 ```

## /api/v1/check/(userAvailable | emailAvailable)/{id}
**Method: GET**  
This Endpoint is used to check if the Email/Username the Client is trying to register is available

#### Response:

| Code | Name |Reason|
|--|--|--|
|  200|OK  |Successful Request|
|400|Bad Request| Malformed Request|
|500|Internal Server Error| Serverside Error

#### Status 200 Success
##### Header:
|Key|Value|
|-|-|
|Content-Type|application/json|

##### Body: 
```
{
	"available": boolean
}
```

#### Status 400/500 Failure
##### Header: 
|Key|Value|
|-|-|
|Content-Type|application/json|

##### Body:

 ```
 {
	 "err_message":"string"
 }
 ```


## Endpoints with Auth
|URL|Endpoint|Type|
|--|--|--|
|/api/v1/list/create|Create List  |POST|
|/api/v1/list/all| Get All Lists|GET|
|/api/v1/list/get/{id}|Get Specific List| GET|
|/api/v1/list/delete/{id}|Delete List|DELETE|
|/api/v1/list/entry/add|Add List Entry  |POST|
|/api/v1/list/entry/edit|Edit List Entry |PUT|
|/api/v1/list/entry/delete|Delete List Entry |DELETE|
|/api/v1/user/change/password|Update User Password|PATCH|
|/api/v1/user/change/email|Update User Email | PATCH|
|/api/v1/user/delete|Delete User Data|DELETE|

## Authentification
All of the following endpoints need to contain a valid JsonWebToken in their Request-Header. 
This Token can be optained by calling the Login or Register endpoints first.
|Key|Value|
|-|-|
|Authorization|Bearer JWT-Token|

The value needs to be **"Bearer " + JsonWebToken**.


### Invalid/Missing JWT-Token
In case of an missing or invalid Token **401 Unauthorized** will be returned.


## /api/v1/list/create
**Method: POST**  
This endpoint is used to create a new AdenauerKreuz List and receive its ID in return.

#### Request
##### Header:
|Key|Value|
|-|-|
|Content-Type|application/json|
|Authorization|Bearer JWT-Token|


##### Body:

 ```  
 {
	"name": "string"
 }
 ```

#### Response
| Code | Name |Reason|
|--|--|--|
|  200|OK  |Successful Request|
|400|Bad Request| Malformed Request|
|401|Unauthorized|Invalid/Missing JWT|
|500|Internal Server Error| Serverside Error



#### Status 200 Success
##### Header: 
|Key|Value|
|-|-|
|Content-Type|application/json|


##### Body:

 ```  
 {
	 "listID": int
 }
 ```
 
#### Status 400/500 Failure
##### Header: 
|Key|Value|
|-|-|
|Content-Type|application/json|

##### Body:

 ```  
 {
	 "err_message": "string"
 }
 ```
 
 ## /api/v1/list/all
 
**Method: GET**  
This endpoint is used to get all Lists of a User

#### Request
##### Header:
|Key|Value|
|-|-|
|Authorization|Bearer JWT-Token|


#### Response
| Code | Name |Reason|
|--|--|--|
|  200|OK  |Successful Request|
|400|Bad Request| Malformed Request|
|401|Unauthorized|Invalid/Missing JWT|
|500|Internal Server Error| Serverside Error


#### Status 200 Success
##### Header: 
|Key|Value|
|-|-|
|Content-Type|application/json|


##### Body:

 ```  
{
	"count":  int,
	"lists":  [
        {
        "listID":  int,
        "name":  "string",
        "elements":  [
                {
                    "elementID":  int,
                    "score":  int,
                    "Content":  "string",
                    "connotation":  boolean
                },
                {
                    ...
                }
            ]
        },
        {
            ...
        }
	]
}
 ```
 
#### Status 400/500 Failure
##### Header: 
|Key|Value|
|-|-|
|Content-Type|application/json|

##### Body:

 ```  
 {
	 "err_message": "string"
 }
 ```
## /api/v1/list/get/{id}
**Method: GET**  
This endpoint is used to get a specific List by its ID

#### Request
##### Header:
|Key|Value|
|-|-|
|Authorization|Bearer JWT-Token|


#### Response
| Code | Name |Reason|
|--|--|--|
|  200|OK  |Successful Request|
|400|Bad Request| Malformed Request|
|401|Unauthorized|Invalid/Missing JWT|
|500|Internal Server Error| Serverside Error


#### Status 200 Success
##### Header: 
|Key|Value|
|-|-|
|Content-Type|application/json|


##### Body:

 ```  
 {
	"listID": int,
	"name": "string",
	"elements": [
        {
            "elementID": int,
            "score": int,
            "Content": "string",
            "connotation": boolean
        },
        {
            ...
        }
	]
}
 ```
 
#### Status 400/500 Failure
##### Header: 
|Key|Value|
|-|-|
|Content-Type|application/json|

##### Body:

 ```  
 {
	 "err_message": "string"
 }
 ```

## /api/v1/list/delete/{id}
**Method: DELETE**  
This Endpoint is used to delete a List and its Elements  

#### Request
##### Header:
|Key|Value|
|-|-|
|Authorization|Bearer JWT-Token|

#### Response
| Code | Name |Reason|
|--|--|--|
|  200|OK  |Successful Request|
|400|Bad Request| Malformed Request|
|401|Unauthorized|Invalid/Missing JWT|
|500|Internal Server Error| Serverside Error

#### Status 400/500 Failure
##### Header: 
|Key|Value|
|-|-|
|Content-Type|application/json|

##### Body:

 ```  
 {
	 err_message:"string"
 }
 ```

#### Request
##### Header:
|Key|Value|
|-|-|
|Authorization|Bearer JWT-Token|

## /api/v1/list/entry/add
**Method: POST**  
This endpoint is used to add a Entry to an existing List


#### Request
##### Header:
|Key|Value|
|-|-|
|Content-Type|application/json|
|Authorization|Bearer JWT-Token|


##### Body:

 ```  
{
	"listID": int,
	"entry": {
        "elementID": (int) null,
        "score": int,
        "content": "string",
        "connotation": boolean
    }
}
 ```

#### Response
| Code | Name |Reason|
|--|--|--|
|  200|OK  |Successful Request|
|400|Bad Request| Malformed Request|
|401|Unauthorized|Invalid/Missing JWT|
|500|Internal Server Error| Serverside Error



#### Status 200 Success
##### Header: 
|Key|Value|
|-|-|
|Content-Type|application/json|

##### Body:
```  
{
	"elementID": int,
	"score": int,
	"content": "string",
	"connotation": boolean
}
```


#### Status 400/500 Failure
##### Header: 
|Key|Value|
|-|-|
|Content-Type|application/json|

##### Body:

 ```  
 {
	 err_message:"string"
 }
 ```

## /api/v1/list/entry/edit
**Method: PUT**  
This endpoint is used to update an existing List Entry


### Request
##### Header:
|Key|Value|
|-|-|
|Content-Type|application/json|
|Authorization|Bearer JWT-Token|


##### Body:

 ```  
{
	"listID": int,
	"entry": {
        "elementID": int,
        "score": int,
        "content": "string",
        "connotation": boolean
    }
}
 ```

#### Response
| Code | Name |Reason|
|--|--|--|
|  200|OK  |Successful Request|
|400|Bad Request| Malformed Request|
|401|Unauthorized|Invalid/Missing JWT|
|500|Internal Server Error| Serverside Error



#### Status 200 Success
##### Header: 
|Key|Value|
|-|-|
|Content-Type|application/json|

##### Body:
```  
{
	"elementID": int,
	"score": int,
	"content": "string",
	"connotation": boolean
}
```


#### Status 400/500 Failure
##### Header: 
|Key|Value|
|-|-|
|Content-Type|application/json|

##### Body:

 ```  
 {
	 err_message:"string"
 }
 ```
## /api/v1/list/entry/delete
**Method: DELETE**  
This endpoint is used to delete an existing List Entry


### Request
##### Header:
|Key|Value|
|-|-|
|Content-Type|application/json|
|Authorization|Bearer JWT-Token|


##### Body:

 ```  
{
	"listID": int,
	"elementID": int
}
 ```

#### Response
| Code | Name |Reason|
|--|--|--|
|  200|OK  |Successful Request|
|400|Bad Request| Malformed Request|
|401|Unauthorized|Invalid/Missing JWT|
|500|Internal Server Error| Serverside Error

#### Status 400/500 Failure
##### Header: 
|Key|Value|
|-|-|
|Content-Type|application/json|

##### Body:

 ```  
 {
	 err_message:"string"
 }
 ```
## /api/v1/user/change/password
**Method: PATCH**  
This endpoint is used to delete an existing List Entry


### Request
##### Header:
|Key|Value|
|-|-|
|Content-Type|application/json|
|Authorization|Bearer JWT-Token|


##### Body:

 ```  
{
	"old_password":"SHA-256 string"
    "new_password":"SHA-256 string"
}
 ```

#### Response
| Code | Name |Reason|
|--|--|--|
|  200|OK  |Successful Request|
|400|Bad Request| Malformed Request|
|401|Unauthorized|Invalid/Missing JWT|
|500|Internal Server Error| Serverside Error

#### Status 400/500 Failure
##### Header: 
|Key|Value|
|-|-|
|Content-Type|application/json|

##### Body:

 ```  
 {
	 err_message:"string"
 }
 ```
## /api/v1/user/change/email
**Method: PATCH**  
This endpoint is used to delete an existing List Entry


### Request
##### Header:
|Key|Value|
|-|-|
|Content-Type|application/json|
|Authorization|Bearer JWT-Token|


##### Body:

 ```  
{
	"old_email":"string"
    "new_email":"string"
}
 ```

#### Response
| Code | Name |Reason|
|--|--|--|
|  200|OK  |Successful Request|
|400|Bad Request| Malformed Request|
|401|Unauthorized|Invalid/Missing JWT|
|409|Conflict|Email already Exists or not Valid
|500|Internal Server Error| Serverside Error

#### Status 400/500 Failure
##### Header: 
|Key|Value|
|-|-|
|Content-Type|application/json|

##### Body:

 ```  
 {
	 err_message:"string"
 }
 ```
## /api/v1/user/delete
**Method: DELETE**  
This endpoint is used to delete an existing List Entry


### Request
##### Header:
|Key|Value|
|-|-|
|Authorization|Bearer JWT-Token|


#### Response
| Code | Name |Reason|
|--|--|--|
|  200|OK  |Successful Request|
|400|Bad Request| Malformed Request|
|401|Unauthorized|Invalid/Missing JWT|
|500|Internal Server Error| Serverside Error

#### Status 400/500 Failure
##### Header: 
|Key|Value|
|-|-|
|Content-Type|application/json|

##### Body:

 ```  
 {
	 err_message:"string"
 }
 ```
