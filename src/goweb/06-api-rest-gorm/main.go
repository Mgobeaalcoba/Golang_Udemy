package main

import (
	"gorm/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	// Ejecuto función GORM de migración:

	// models.MigrarUser() // Transforma el objeto user a una tabla users en el motor de base de datos. En este caso sin el registro de creación ya que por cada atributo se genera una columna y create_date no es un atributo de la struct user. MigrarUser() solo se ejecuta si no existe la tabla. Si ya existe no se ejecuta.

	// RUTAS:

	mux := mux.NewRouter() // Metodo de gorilla/mux

	// ENDPOINTS:

	// EndPoint "GET" para obtener info de usuarios:

	mux.HandleFunc("/api/user/", handlers.GetUsers).Methods("GET") // Ruta para obtener todos los usuarios.

	mux.HandleFunc("/api/user/{id:[0-9]+}", handlers.GetUser).Methods("GET") // Ruta para obtener solo un usuario.

	mux.HandleFunc("/api/user/", handlers.CreateUser).Methods("POST") // Ruta para registrar un nuevo usuario.

	mux.HandleFunc("/api/user/{id:[0-9]+}", handlers.UpdateUser).Methods("PUT")

	mux.HandleFunc("/api/user/{id:[0-9]+}", handlers.DeleteUser).Methods("DELETE")

	// SERVIDOR:

	log.Fatal(http.ListenAndServe("localhost:5007", mux)) // usado para handlerssinresponse

	/*

		// RUTAS:

		mux := mux.NewRouter() // Metodo de gorilla/mux

		// ENDPOINTS:

		// EndPoint "GET" para obtener info de usuarios:

		mux.HandleFunc("/api/user/", handlers.GetUsers).Methods("GET")           // Ruta para obtener todos los usuarios.
		mux.HandleFunc("/api/user/{id:[0-9]+}", handlers.GetUser).Methods("GET") // Ruta para obtener solo un usuario.

		// EndPoint "POST" para agregar info de usuarios:

		mux.HandleFunc("/api/user/", handlers.CreateUser).Methods("POST") // Ruta para registrar un nuevo usuario.

		// EndPoint "PUT" para editar info de usuarios:

		mux.HandleFunc("/api/user/{id:[0-9]+}", handlers.UpdateUser).Methods("PUT") //

		// EndPoint "DELETE" para eliminar info de usuarios:

		mux.HandleFunc("/api/user/{id:[0-9]+}", handlers.DeleteUser).Methods("DELETE") //

		// SERVIDOR:

		//log.Fatal(http.ListenAndServe("localhost:5007", mux)) // usado para handlerssinresponse
		log.Fatal(http.ListenAndServe("localhost:5007", mux))

	*/

}

// Para crear los handlers de cada uno de los endpoints voy a armarlos en una carpeta sobre la raiz llamada "handlers" con un file dentro de ella llamado "handlers.go"
