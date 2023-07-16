# Movie CRUD API
This is a CRUD API for managing movies. It provides endpoints to create, read, update, and delete movies.  
I used a tutorial video to learn and create this project: https://youtu.be/jFfo23yIWac?t=1234  


## API Endpoints
The following endpoints are available:  
GET /movies: Get a list of all movies.  
GET /movies/{id}: Get details of a specific movie by ID.  
POST /movies: Create a new movie.  
PUT /movies/{id}: Update an existing movie by ID.  
DELETE /movies/{id}: Delete a movie by ID.  

### Request and Response Examples
#### Get all movies  
Request: GET http://localhost:8000/movies  
Response:  
[{  
    "id": "1",  
    "isbn": "438227",  
    "title": "Movie One",  
    "director": { "firstname": "John", "lastname": "Doe" }  
  },  
  {  
    "id": "2",  
    "isbn": "45455",  
    "title": "Movie Two",  
    "director": { "firstname": "Steve", "lastname": "Smith" }  
  }]  

#### Get a movie by ID  
Request: GET http://localhost:8000/movies/1  
Response:  
{  
  "id": "1",  
  "isbn": "438227",  
  "title": "Movie One",  
  "director": { "firstname": "John", "lastname": "Doe" }  
}  

#### Create a movie  
Request: POST http://localhost:8000/movies  
{  
  "isbn": "9781234567890",  
  "title": "New Movie",  
  "director": { "firstname": "Jane", "lastname": "Smith" }  
}  

Response:  
{  
  "id": "3",  
  "isbn": "9781234567890",  
  "title": "New Movie",  
  "director": { "firstname": "Jane", "lastname": "Smith" }  
}  

#### Update a movie  
Request: PUT http://localhost:8000/movies/3  
{  
  "title": "Updated Movie"  
}  

Response:  
{  
  "id": "3",  
  "isbn": "9781234567890",  
  "title": "Updated Movie",  
  "director": { "firstname": "Jane", "lastname": "Smith" }  
}  

#### Delete a movie  
Request: DELETE http://localhost:8000/movies/3  
Response: No content (HTTP status code: 204)
