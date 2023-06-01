# MoviePlus  Server (Backend)
<img src="https://i.ibb.co/51PDVwk/gin.png" alt="gin" border="0" width="400" align="center"> 

## Project Summary 
* This  is a Restful Gin application with CRUD operations that allow users to create an account to the tube-max app, save or add actors, tvshows, albums, artists, songs and movies , delete those actors and movies and login and out using Go jwt authentication.
* It uses  Gin with Tomcat server as a framework.
* The app has 19 endpoints namely : /home, /signup, /logout, /verifification, /refresh-token, /login, /movie, /actor, /movies, /actors, /movie/id, tshows/id, track/id, album/id, tracks, albums, artists, artist/id, and /actor/id.
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
Built 9 Controllers, authController, actorController, userController, albumController, artistController, trackController, tvShowController, verificationController and movieController.
#### **User Account Creation Endpoints:** 
* **/signup (PostMethod)**: Takes in firstname, lastname, password and email for user signup. A Jwt token is created as an authentication tool, its stored on the database and also sent by go mail to user email for verification. The password is encrypted using BCryptPasswordEncoder.

* **/verification  (GetMethod)**: validates the email token against the one on the database, once verified the account is enabled. 
* **/login  (GetMethod)**: A Jwt token is created and returned if user login credentials are valid. 


#### **Actors Endpoints:**  
* **/actor (PostMethod)**
* **/actors (GetMethod)**
* **/actor/id (DeleteMethod)**

#### **Movies Endpoints:**  
* **/movie (PostMethod)** 
* **/movies (GetMethod)**
* **/movie/id (DeleteMethod)**  

#### **TvShows Endpoints:**  
* **/tvShow (PostMethod)** 
* **/tvShows (GetMethod)**
* **/tvShow/id (DeleteMethod)** 

#### **Tracks Endpoints:**  
* **/track (PostMethod)** 
* **/tracks (GetMethod)**
* **/track/id (DeleteMethod)** 

#### **Artists Endpoints:**  
* **/artist (PostMethod)**. 
* **/artists (GetMethod)**
* **/artist/id (DeleteMethod)**

#### **Albums Endpoints:**  
* **/album (PostMethod)**
* **/albums (GetMethod)**
* **/album/id (DeleteMethod)**


### **Data Storage**
Used GORM (ORM) to persist and retrieve data from a postgreSQL database.  
Built 8 models: 
* User Model to store app users.
* VerificationToken Model to store signup verification tokens.
* Actor Model. 
* Movie Model.
* Artist Model.
* TvShow Model.
* Track Model.
* Album Model.



### **Productionization**
***
In this step I deployed the postgreSQL database to Aws via 3rd party and deployed the Gin app to Render Cloud.

**Live Implemantation:** [Tube-Max](https://tube-max.vercel.app/)