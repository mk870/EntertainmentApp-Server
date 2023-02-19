# MoviePlus  Server (Backend)
<img src="https://i.ibb.co/51PDVwk/gin.png" alt="gin" border="0" width="400" align="center"> 

## Project Summary 
* This  is a Restful Gin application with CRUD operations that allow users to create an account to the movieplus app, save or add actors and movies , delete those actors and movies and login and out using Go jwt authentication.
* It uses  Gin with Tomcat server as a framework.
* The app has 12 endpoints namely : /home, /signup, /logout, /verifification, /refresh-token, /login, /movie, /actor, /movies, /actors, /movie/id and /actor/id.
* Uses Go Jwt to secure these endpoints.
* Uses Gin GORM to persist data to a postgreSQL database.


### **Resources Used**
***
**Go Version**: 1.19.4

**Dependencies**: Jwt Token, GORM, Go Mail, GIN, postgreSQL-Driver, Go-Cors and Google uuid.  
![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=flat&logo=go&logoColor=white) 	![JWT](https://img.shields.io/badge/JWT-black?style=flat&logo=JSON%20web%20tokens) 	![Postgres](https://img.shields.io/badge/postgres-%23316192.svg?style=flat&logo=postgresql&logoColor=white)

**For Web Framework Requirements**: go.mod

**APIs**: None

### **EndPoints Building**
***
Built 5 Controllers, authController, actorController,userController, verificationController and movieController.
#### **User Account Creation Endpoints:** 
* **/signup (PostMethod)**: Takes in firstname, lastname, password and email for user signup. A Jwt token is created as an authentication tool, its stored on the database and also sent by go mail to user email for verification. The password is encrypted using BCryptPasswordEncoder.

* **/verification  (GetMethod)**: validates the email token against the one on the database, once verified the account is enabled. 
* **/login  (GetMethod)**: A Jwt token is created and returned if user login credentials are valid. 


#### **UserActors Endpoints:**  
* **/actor (PostMethod)**:  saves users' actors to the database with all the actor's properties like name, gender, birthday, tmdb_id and birth_place. 
* **/actors (GetMethod)**:  retrieves all the saved actors of a client from the database.
* **/actor/id (DeleteMethod)** : deletes a specific actor by id from the database.

#### **UserMovies Endpoints:**  
* **/movie (PostMethod)**:  saves users' movies to the database with all the movie's properties like title, runtime, release_date and tmdb_id. 
* **/movies (GetMethod)**:  retrieves all the saved movies of a client from the database.
* **/movie/id (DeleteMethod)** : deletes a specific movie by id from the database.  

### **Data Storage**
Used GORM (ORM) to persist and retrieve data from a postgreSQL database.  
Built 4 models: 
* User Model to store app users.
* VerificationToken Model to store signup verification tokens.
* Actor Model to store favourite actors of users. 
* Movie Model to store favourite movies of users.



### **Productionization**
***
In this step I deployed the postgreSQL database to AWS via 3rd party and deployed the Gin app to Railway Cloud.

**Live Implemantation:** [MoviePlus](https://movie-plus-frontend.vercel.app)