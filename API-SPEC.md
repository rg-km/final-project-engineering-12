API Specification
----------------------------------------------------------------

All API must use this authentication
Request:
- Header:
    - X-Api-Key: "your secret api key"

SUMMARY:
- Users:                            Line 20
- User_course:                      Line 154
- User_details:                     Line 253
- Courses:                          Line 392
- Modules:                          Line 527
- Module_submission:                Line 654
- Module_articles:                  Line 769
- Answers:                          Line 872
- Question:                         Line 989

// Users
------------------------------
Create Users
------------------------------
Request:
- Method: POST
- Endpoint: /api/users
- Header:
    - Content-Type: application/json
    - Accept: application/json
- Body:
    {
        "name" : "string",
        "username" : "string", // unique
        "email" : "string", // unique
        "password" : "string",
        "role" : "integer", // enum(1,2)
        "email_verified_at" : "timestamp", // timestamp
    }
Response:
    {   
        "code" : "number",
        "status" : "string",
        "data" : {
            "id" : "integer", // primary key
            "name" : "string",
            "username" : "string", // unique
            "email" : "string", // unique
            "password" : "string",
            "role" : "integer", // enum(1,2)
            "email_verified_at" : "timestamp", // timestamp
            "created_at" : "timestamp", // timestamp
            "updated_at" : "timestamp" // timestamp
        }
    }
------------------------------
Get Users
------------------------------
Request:
- Method: GET
- Endpoint: /api/users/{id}
- Header:
    - Accept: application/json
Response:
    {
        "code" : "number",
        "status" : "string",
        "data" : {
            "id" : "integer", // primary key
            "name" : "string",
            "username" : "string", // unique
            "email" : "string", // unique
            "password" : "string",
            "role" : "integer", // enum(1,2)
            "email_verified_at" : "timestamp", // timestamp
            "created_at" : "timestamp", // timestamp
            "updated_at" : "timestamp" // timestamp
        }
    }
------------------------------
Update Users
------------------------------
Request:
- Method: PUT
- Endpoint: /api/users/{id}
- Header:
    - Content-Type: application/json
    - Accept: application/json
- Body:
    {
        "name" : "string",
        "username" : "string", // unique
        "email" : "string", // unique
        "password" : "string",
        "role" : "integer", // enum(1,2)
    }
Response:
    {
        "code" : "number",
        "status" : "string",
        "data" : {
            "id" : "integer", // primary key
            "name" : "string",
            "username" : "string", // unique
            "email" : "string", // unique
            "password" : "string",
            "role" : "integer", // enum(1,2)
            "email_verified_at" : "timestamp", // timestamp
            "created_at" : "timestamp", // timestamp
            "updated_at" : "timestamp" // timestamp
        }
    }
------------------------------
List Users
------------------------------
Request:
- Method: GET
- Endpoint: /api/users
- Header:
    - Accept: application/json
- Query Param:
    - size : number
    - page : number
Response:
    {
        "code" : "number",
        "status" : "string",
        "data" : [
            {
                "id" : "integer", // primary key
                "name" : "string",
                "username" : "string", // unique
                "email" : "string", // unique
                "password" : "string",
                "role" : "integer", // enum(1,2)
                "email_verified_at" : "timestamp", // timestamp
                "created_at" : "timestamp", // timestamp
                "updated_at" : "timestamp" // timestamp
            },
        ]
    }
------------------------------
Delete Users
------------------------------
Request:
- Method: DELETE 
- Endpoint: /api/users/{id}
- Header:
    - Accept: application/json
Response:
    {
        "code" : "number",
        "status" : "string"
    }
// User_course
------------------------------
Create User_course
------------------------------
Request:
- Method: POST
- Endpoint: /api/user_course
- Header:
    - Content-Type: application/json
    - Accept: application/json
- Body:
    {
        "user_id" : "integer", // foreign key1
        "course_id" : "integer" // foreign key2
    }
Response:
    {
        "code" : "number",
        "status" : "string",
        "data" : {
            "user_id" : "integer", // foreign key1
            "course_id" : "integer" // foreign key2
        }
    }
------------------------------
Get User_course
------------------------------
Request:
- Method: GET
- Endpoint: /api/user_course/{user_id}/{course_id}
- Header:
    - Accept: application/json
Response:
    {
        "code" : "number",
        "status" : "string",
        "data" : {
            "user_id" : "integer", // foreign key1
            "course_id" : "integer" // foreign key2
        }
    }
------------------------------
Update User_course
------------------------------
Request:
- Method: PUT
- Endpoint: /api/user_course/{user_id}/{course_id}
- Header:
    - Content-Type: application/json
    - Accept: application/json
- Body:
    {
        "user_id" : "integer", // foreign key1
        "course_id" : "integer" // foreign key2
    }
Response:
    {
        "code" : "number",
        "status" : "string",
        "data" : {
            "user_id" : "integer", // foreign key1
            "course_id" : "integer" // foreign key2
        }
    }
------------------------------
List User_course
------------------------------
Request:
- Method: GET
- Endpoint: /api/user_course
- Header:
    - Accept: application/json
- Query Param:
    - size : number
    - page : number
Response:
    {
        "code" : "number",
        "status" : "string",
        "data" : [
            {
                "user_id" : "integer", // foreign key1
                "course_id" : "integer" // foreign key2
            },
        ]
    }
------------------------------
Delete User_course
------------------------------
Request:
- Method: DELETE
- Endpoint: /api/user_course/{user_id}/{course_id}
- Header:
    - Accept: application/json
Response:
    {
        "code" : "number",
        "status" : "string"
    }
// User_details
------------------------------
Create User_details
------------------------------
Request:
- Method: POST
- Endpoint: /api/user_details
- Header:
    - Content-Type: application/json
    - Accept: application/json
- Body:
    {
        "user_id" : "integer", // foreign key1
        "phone" : "string",
        "gender" : "integer", // enum(1,2)
        "type_of_disability" : "integer", // enum(1,2,3)
        "address" : "string",
        "birthdate" : "date", // date
        "image" : "string,
        "description" : "string"
    }
Response:
    {
        "code" : "number",
        "status" : "string",
        "data" : {
            "id" : "integer", // primary key
            "user_id" : "integer", // foreign key1
            "phone" : "string",
            "gender" : "integer", // enum(1,2)
            "type_of_disability" : "integer", // enum(1,2,3)
            "address" : "string",
            "birthdate" : "date", // date
            "image" : "string,
            "description" : "string"
        }
    }
------------------------------
Get User_details
------------------------------
Request:
- Method: GET
- Endpoint: /api/user_details/{id}
- Header:
    - Accept: application/json
Response:
    {
        "code" : "number",
        "status" : "string",
        "data" : {
            "id" : "integer", // primary key
            "user_id" : "integer", // foreign key1
            "phone" : "string",
            "gender" : "integer", // enum(1,2)
            "type_of_disability" : "integer", // enum(1,2,3)
            "address" : "string",
            "birthdate" : "date", // date
            "image" : "string,
            "description" : "string"
        }
    }
------------------------------
Update User_details
------------------------------
Request:
- Method: PUT
- Endpoint: /api/user_details/{id}
- Header:
    - Content-Type: application/json
    - Accept: application/json
- Body:
    {
        "user_id" : "integer", // foreign key1
        "phone" : "string",
        "gender" : "integer", // enum(1,2)
        "type_of_disability" : "integer", // enum(1,2,3)
        "address" : "string",
        "birthdate" : "date", // date
        "image" : "string,
        "description" : "string"
    }
Response:
    {
        "code" : "number",
        "status" : "string",
        "data" : {
            "id" : "integer", // primary key
            "user_id" : "integer", // foreign key1
            "phone" : "string",
            "gender" : "integer", // enum(1,2)
            "type_of_disability" : "integer", // enum(1,2,3)
            "address" : "string",
            "birthdate" : "date", // date
            "image" : "string,
            "description" : "string"
        }
    }
------------------------------
List User_details
------------------------------
Request:
- Method: GET
- Endpoint: /api/user_details
- Header:
    - Accept: application/json
- Query Param:
    - size : number
    - page : number
Response:
    {
        "code" : "number",
        "status" : "string",
        "data" : [
            {
                "id" : "integer", // primary key
                "user_id" : "integer", // foreign key1
                "phone" : "string",
                "gender" : "integer", // enum(1,2)
                "type_of_disability" : "integer", // enum(1,2,3)
                "address" : "string",
                "birthdate" : "date", // date
                "image" : "string,
                "description" : "string"
            },
        ]
    }
------------------------------
Delete User_details
------------------------------
Request:
- Method: DELETE
- Endpoint: /api/user_details/{id}
- Header:
    - Accept: application/json
Response:
    {
        "code" : "number",
        "status" : "string"
    }
// Courses
------------------------------
Create Courses
------------------------------
Request:
- Method: POST
- Endpoint: /api/courses
- Header:
    - Content-Type: application/json
    - Accept: application/json
- Body:
    {
        "name" : "string",
        "code_course" : "string", // unique
        "class" : "string",
        "tools" : "string", // longtext
        "about" : "string", // longtext
        "description" : "string" // longtext
    }   
Response:
    {
        "code" : "number",
        "status" : "string",
        "data" : {
            "id" : "integer", // primary key
            "name" : "string",
            "code_course" : "string", // unique
            "class" : "string",
            "tools" : "string", // longtext
            "about" : "string", // longtext
            "description" : "string", // longtext
            "created_at" : "timestamp", // timestamp
            "updated_at" : "timestamp" // timestamp 
        }
    }
------------------------------
Get Courses
------------------------------
Request:
- Method: GET
- Endpoint: /api/courses/{id}
- Header:
    - Accept: application/json  
Response:
    {
        "code" : "number",
        "status" : "string",
        "data" : {
            "id" : "integer", // primary key
            "name" : "string",
            "code_course" : "string", // unique
            "class" : "string",
            "tools" : "string", // longtext
            "about" : "string", // longtext
            "description" : "string", // longtext
            "created_at" : "timestamp", // timestamp
            "updated_at" : "timestamp" // timestamp 
        }
    }
------------------------------
Update Courses
------------------------------
Request:
- Method: PUT
- Endpoint: /api/courses/{id}
- Header:
    - Content-Type: application/json
    - Accept: application/json
- Body:
    {
        "name" : "string",
        "code_course" : "string", // unique
        "class" : "string",
        "tools" : "string", // longtext
        "about" : "string", // longtext
        "description" : "string" // longtext
    }
Response:
    {
        "code" : "number",
        "status" : "string",
        "data" : {
            "id" : "integer", // primary key
            "name" : "string",
            "code_course" : "string", // unique
            "class" : "string",
            "tools" : "string", // longtext
            "about" : "string", // longtext
            "description" : "string", // longtext
            "created_at" : "timestamp", // timestamp
            "updated_at" : "timestamp" // timestamp 
        }
    }
------------------------------
List Courses
------------------------------
Request:
- Method: GET
- Endpoint: /api/courses
- Header:
    - Accept: application/json
- Query Param:
    - size : number
    - page : number
Response:
    {
        "code" : "number",
        "status" : "string",
        "data" : [
            {
                "id" : "integer", // primary key
                "name" : "string",
                "code_course" : "string", // unique
                "class" : "string",
                "tools" : "string", // longtext
                "about" : "string", // longtext
                "description" : "string", // longtext
                "created_at" : "timestamp", // timestamp
                "updated_at" : "timestamp" // timestamp 
            },
        ]
    }
------------------------------
Delete Courses
------------------------------
Request:
- Method: DELETE
- Endpoint: /api/courses/{id}
- Header:
    - Accept: application/json
Response:
    {
        "code" : "number",
        "status" : "string"
    }
// Modules
------------------------------
Create Modules
------------------------------
Request:
- Method: POST
- Endpoint: /api/modules
- Header:
    - Content-Type: application/json
    - Accept: application/json
- Body:
    {
        "course_id" : "integer", // foreign key1
        "name" : "string",
        "is_locked" : "boolean",
        "estimate" : "integer",
        "deadline" : "timestamp", // timestamp
        "grade" : "integer"
    }
Response:
    {
        "code" : "number",
        "status" : "string",
        "data" : {
            "id" : "integer", // primary key
            "course_id" : "integer", // foreign key1
            "name" : "string",
            "is_locked" : "boolean",
            "estimate" : "integer",
            "deadline" : "timestamp", // timestamp
            "grade" : "integer"
        }
    }
------------------------------
Get Modules
------------------------------
Request:
- Method: GET
- Endpoint: /api/modules/{id}
- Header:
    - Accept: application/json
Response:
    {
        "code" : "number",
        "status" : "string",
        "data" : {
            "id" : "integer", // primary key
            "course_id" : "integer", // foreign key1
            "name" : "string",
            "is_locked" : "boolean",
            "estimate" : "integer",
            "deadline" : "timestamp", // timestamp
            "grade" : "integer"
        }
    }
------------------------------
Update Modules
------------------------------
Request:
- Method: PUT
- Endpoint: /api/modules/{id}
- Header:
    - Content-Type: application/json
    - Accept: application/json
- Body:
    {
        "course_id" : "integer", // foreign key1
        "name" : "string",
        "is_locked" : "boolean",
        "estimate" : "integer",
        "deadline" : "timestamp", // timestamp
        "grade" : "integer"
    }
Response:
    {
        "code" : "number",
        "status" : "string",
        "data" : {
            "id" : "integer", // primary key
            "course_id" : "integer", // foreign key1
            "name" : "string",
            "is_locked" : "boolean",
            "estimate" : "integer",
            "deadline" : "timestamp", // timestamp
            "grade" : "integer"
        }
    }
------------------------------
List Modules
------------------------------
Request:
- Method: GET
- Endpoint: /api/modules
- Header:
    - Accept: application/json
- Query Param:
    - size : number
    - page : number
Response:
    {
        "code" : "number",
        "status" : "string",
        "data" : [
            {
                "id" : "integer", // primary key
                "course_id" : "integer", // foreign key1
                "name" : "string",
                "is_locked" : "boolean",
                "estimate" : "integer",
                "deadline" : "timestamp", // timestamp
                "grade" : "integer"
            },
        ]
    }
------------------------------
Delete Modules
------------------------------
Request:
- Method: DELETE
- Endpoint: /api/modules/{id}
- Header:
    - Accept: application/json
Response:
    {
        "code" : "number",
        "status" : "string"
    }
// Module_submissions
------------------------------
Create Module_submissions
------------------------------
Request:
- Method: POST
- Endpoint: /api/module_submissions
- Header:
    - Content-Type: application/json
    - Accept: application/json
- Body:
    {
        "module_id" : "integer", // foreign key1
        "file" : "string",
        "type" : "string",
        "max_size" : "integer"
    }
Response:
    {
        "code" : "number",
        "status" : "string",
        "data" : {
            "id" : "integer", // primary key
            "module_id" : "integer", // foreign key1
            "file" : "string",
            "type" : "string",
            "max_size" : "integer"
        }
    }
------------------------------
Get Module_submissions
------------------------------
Request:
- Method: GET
- Endpoint: /api/module_submissions/{id}
- Header:
    - Accept: application/json
Response:
    {
        "code" : "number",
        "status" : "string",
        "data" : {
            "id" : "integer", // primary key
            "module_id" : "integer", // foreign key1
            "file" : "string",
            "type" : "string",
            "max_size" : "integer"
        }
    }
------------------------------
Update Module_submissions
------------------------------
Request:
- Method: PUT
- Endpoint: /api/module_submissions/{id}
- Header:
    - Content-Type: application/json
    - Accept: application/json
- Body:
    {
        "module_id" : "integer", // foreign key1
        "file" : "string",
        "type" : "string",
        "max_size" : "integer"
    }
Response:
    {
        "code" : "number",
        "status" : "string",
        "data" : {
            "id" : "integer", // primary key
            "module_id" : "integer", // foreign key1
            "file" : "string",
            "type" : "string",
            "max_size" : "integer"
        }
    }
------------------------------
List Module_submissions
------------------------------
Request:
- Method: GET
- Endpoint: /api/module_submissions
- Header:
    - Accept: application/json
- Query Param:
    - size : number
    - page : number
Response:
    {
        "code" : "number",
        "status" : "string",
        "data" : [
            {
                "id" : "integer", // primary key
                "module_id" : "integer", // foreign key1
                "file" : "string",
                "type" : "string",
                "max_size" : "integer"
            },
        ]
    }
------------------------------
Delete Module_submissions
------------------------------
Request:
- Method: DELETE
- Endpoint: /api/module_submissions/{id}
- Header:
    - Accept: application/json
Response:
    {
        "code" : "number",
        "status" : "string"
    }
// Module_articles
------------------------------
Create Module_articles
------------------------------
Request:
- Method: POST
- Endpoint: /api/module_articles
- Header:
    - Content-Type: application/json
    - Accept: application/json
- Body:
    {
        "module_id" : "integer", // foreign key1
        "content" : "string" // longtext
    }
Response:
    {
        "code" : "number",
        "status" : "string",
        "data" : {
            "id" : "integer", // primary key
            "module_id" : "integer", // foreign key1
            "content" : "string" // longtext
        }
    }
------------------------------
Get Module_articles
------------------------------
Request:
- Method: GET
- Endpoint: /api/module_articles/{id}
- Header:
    - Accept: application/json
Response:
    {
        "code" : "number",
        "status" : "string",
        "data" : {
            "id" : "integer", // primary key
            "module_id" : "integer", // foreign key1
            "content" : "string" // longtext
        }
    }
------------------------------
Update Module_articles
------------------------------
Request:
- Method: PUT
- Endpoint: /api/module_articles/{id}
- Header:
    - Content-Type: application/json
    - Accept: application/json
- Body:
    {
        "module_id" : "integer", // foreign key1
        "content" : "string" // longtext
    }
Response:
    {
        "code" : "number",
        "status" : "string",
        "data" : {
            "id" : "integer", // primary key
            "module_id" : "integer", // foreign key1
            "content" : "string" // longtext
        }
    }
------------------------------
List Module_articles
------------------------------
Request:
- Method: GET
- Endpoint: /api/module_articles
- Header:
    - Accept: application/json
- Query Param:
    - size : number
    - page : number
Response:
    {
        "code" : "number",
        "status" : "string",
        "data" : [
            {
                "id" : "integer", // primary key
                "module_id" : "integer", // foreign key1
                "content" : "string" // longtext
            },
        ]
    }
------------------------------
Delete Module_articles
------------------------------
Request:
- Method: DELETE
- Endpoint: /api/module_articles/{id}
- Header:
    - Accept: application/json
Response:
    {
        "code" : "number",
        "status" : "string"
    }
// Answers
------------------------------
Create Answers
------------------------------
Request:
- Method: POST
- Endpoint: /api/answers
- Header:
    - Content-Type: application/json
    - Accept: application/json
- Body:
    {
        "question_id" : "integer", // foreign key1
        "user_id" : "integer", // foreign key2
        "description" : "string"
    }
Response:
    {
        "code" : "number",
        "status" : "string",
        "data" : {
            "id" : "integer", // primary key
            "question_id" : "integer", // foreign key1
            "user_id" : "integer", // foreign key2
            "description" : "string",
            "created_at" : "timestamp", // timestamp
            "updated_at" : "timestamp" // timestamp
        }
    }
------------------------------
Get Answers
------------------------------
Request:
- Method: GET
- Endpoint: /api/answers/{id}
- Header:
    - Accept: application/json
Response:
    {
        "code" : "number",
        "status" : "string",
        "data" : {
            "id" : "integer", // primary key
            "question_id" : "integer", // foreign key1
            "user_id" : "integer", // foreign key2
            "description" : "string",
            "created_at" : "timestamp", // timestamp
            "updated_at" : "timestamp" // timestamp
        }
    }
------------------------------
Update Answers
------------------------------
Request:
- Method: PUT
- Endpoint: /api/answers/{id}
- Header:
    - Content-Type: application/json
    - Accept: application/json
- Body:
    {
        "question_id" : "integer", // foreign key1
        "user_id" : "integer", // foreign key2
        "description" : "string"
    }
Response:
    {
        "code" : "number",
        "status" : "string",
        "data" : {
            "id" : "integer", // primary key
            "question_id" : "integer", // foreign key1
            "user_id" : "integer", // foreign key2
            "description" : "string",
            "created_at" : "timestamp", // timestamp
            "updated_at" : "timestamp" // timestamp
        }
    }
------------------------------
List Answers
------------------------------
Request:
- Method: GET
- Endpoint: /api/answers
- Header:
    - Accept: application/json
- Query Param:
    - size : number
    - page : number
Response:
    {
        "code" : "number",
        "status" : "string",
        "data" : [
            {
                "id" : "integer", // primary key
                "question_id" : "integer", // foreign key1
                "user_id" : "integer", // foreign key2
                "description" : "string",
                "created_at" : "timestamp", // timestamp
                "updated_at" : "timestamp" // timestamp
            },
        ]
    }
------------------------------
Delete Answers
------------------------------
Request:
- Method: DELETE
- Endpoint: /api/answers/{id}
- Header:
    - Accept: application/json
Response:
    {
        "code" : "number",
        "status" : "string"
    }
// Questions
------------------------------
Create Questions
------------------------------
Request:
- Method: POST
- Endpoint: /api/questions
- Header:
    - Content-Type: application/json
    - Accept: application/json
- Body:
    {
        "module_id" : "integer", // foreign key1
        "user_id" : "integer", // foreign key2
        "title" : "string",
        "tags" : "string",
        "description" : "string"
    }
Response:
    {
        "code" : "number",
        "status" : "string",
        "data" : {
            "id" : "integer", // primary key
            "module_id" : "integer", // foreign key1
            "user_id" : "integer", // foreign key2
            "title" : "string",
            "tags" : "string",
            "description" : "string",
            "created_at" : "timestamp", // timestamp
            "updated_at" : "timestamp" // timestamp
        }
    }
------------------------------
Get Questions
------------------------------
Request:
- Method: GET
- Endpoint: /api/questions/{id}
- Header:
    - Accept: application/json
Response:
    {
        "code" : "number",
        "status" : "string",
        "data" : {
            "id" : "integer", // primary key
            "module_id" : "integer", // foreign key1
            "user_id" : "integer", // foreign key2
            "title" : "string",
            "tags" : "string",
            "description" : "string",
            "created_at" : "timestamp", // timestamp
            "updated_at" : "timestamp" // timestamp
        }
    }
------------------------------
Update Questions
------------------------------
Request:
- Method: PUT
- Endpoint: /api/questions/{id}
- Header:
    - Content-Type: application/json
    - Accept: application/json
- Body:
    {
        "module_id" : "integer", // foreign key1
        "user_id" : "integer", // foreign key2
        "title" : "string",
        "tags" : "string",
        "description" : "string"
    }
Response:
    {
        "code" : "number",
        "status" : "string",
        "data" : {
            "id" : "integer", // primary key
            "module_id" : "integer", // foreign key1
            "user_id" : "integer", // foreign key2
            "title" : "string",
            "tags" : "string",
            "description" : "string",
            "created_at" : "timestamp", // timestamp
            "updated_at" : "timestamp" // timestamp
        }
    }
------------------------------
List Questions
------------------------------
Request:
- Method: GET
- Endpoint: /api/questions
- Header:
    - Accept: application/json
- Query Param:
    - size : number
    - page : number
Response:
    {
        "code" : "number",
        "status" : "string",
        "data" : [
            {
                "id" : "integer", // primary key
                "module_id" : "integer", // foreign key1
                "user_id" : "integer", // foreign key2
                "title" : "string",
                "tags" : "string",
                "description" : "string",
                "created_at" : "timestamp", // timestamp
                "updated_at" : "timestamp" // timestamp
            },
        ]
    }
------------------------------
Delete Questions
------------------------------
Request:
- Method: DELETE
- Endpoint: /api/questions/{id}
- Header:
    - Accept: application/json
Response:
    {
        "code" : "number",
        "status" : "string"
    }