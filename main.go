package main

import (
	"movieplusApi/controllers"
	"movieplusApi/db"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowCredentials = true
	config.AddAllowHeaders("Authorization", "token", "User-Agent", "Accept")
	router.Use(cors.New(config))
	db.Connect()
	controllers.GoogleSignUp(router)
	controllers.GoogleLogin(router)
	controllers.CreateUser(router)
	controllers.GetUsers(router)
	controllers.UpdateUser(router)
	controllers.GetUser(router)
	controllers.DeleteUser(router)
	controllers.Login(router)
	controllers.LoginOut(router)
	controllers.CreateActor(router)
	controllers.GetActors(router)
	controllers.DeleteActor(router)
	controllers.CreateMovie(router)
	controllers.GetMovies(router)
	controllers.DeleteMovie(router)
	controllers.GetVerificationToken(router)
	controllers.VerificationTokenValidation(router)
	controllers.GetNewAccessToken(router)
	controllers.CreateAlbum(router)
	controllers.DeleteAlbum(router)
	controllers.GetAlbums(router)
	controllers.GetArtists(router)
	controllers.CreateArtist(router)
	controllers.DeleteArtist(router)
	controllers.CreateTrack(router)
	controllers.DeleteTrack(router)
	controllers.GetTracks(router)
	controllers.CreateTvShow(router)
	controllers.DeleteTvShow(router)
	controllers.GetTvShows(router)
	controllers.GetNews(router)
	router.Run()
}
