# Go Demo Microservice - Users API

Go microservice exploration for a blog post - Basic Users API

Structure:
```
app/
  application.go : Webserver starting/stopping logistics
  mappings.go : Maps URLs to route handlers
  
 
models/ : Stores all models for interacting with API
    user/
      user.go

routers/ : Parse requests, handle errors, and call service layer (business logic)
    user/
      user.go : Route handlers for operations that have to do with users (calls user service fts)

services/ : Actual functionality (business logic) ex. CRUD operations
    user/
      user.go : Functions that handle user operations ex. creating user
     
utils/ : General utilities
    errors/
      error.go : Error structs/utilities ex. RestError, BadRequestError, ...
```  

    
