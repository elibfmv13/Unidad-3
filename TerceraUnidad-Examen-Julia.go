package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

/*
http://localhost:8069/libros ---> POST
http://localhost:8069/libros ---> GET
http://localhost:8069/libros/1 ---> GET
http://localhost:8069/libros/1 ---> DELETE
*/

var DB *sql.DB // variable global db

type Libros struct {
	Id        int    `json:"id"`
	Nombre      string `json:"nombre"`
	Descripcion       string `json:"descripcion"`
	Autor 		string `json:"autor"`
	Editorial 		string `json:"editorial"`
	Fecha string `json:"fecha_publicacion"`

}

func main() {
	r := gin.Default()

	DB, err := sql.Open("mysql", "root:12345@/biblioteca")
	if err != nil {
		panic(err.Error())
	}
	defer DB.Close()

	err = DB.Ping()
	if err != nil {
		panic(err.Error())
	}

	r.POST("/libros", func(c *gin.Context) {
		libro := Libros{}                 // crear estructura donde se guardara el json de usuario
		err := c.ShouldBindJSON(&libro) // func que decodifica el body del request en ela estructura y valida que sea un json valido
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		stmt, err := DB.Query("INSERT INTO libros (`nombre`, `descripcion`, `autor`, `Editorial`, `fecha_publicacion`) VALUES (?,?,?,?,?)", libro.Nombre, libro.Descripcion, libro.Autor, libro.Editorial, libro.Fecha)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		defer stmt.Close()

		c.JSON(200, libro)
	})

	r.GET("/libros", func(c *gin.Context) {
		rows, err := DB.Query("SELECT * FROM libros")
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		defer rows.Close()

		var libro1 []Libros // array donde se guardaran los datos traidos por el query
		for rows.Next() {
			var libro Libros // se crea una var temporal para asignar el valor de la iteracion
			rows.Scan(&libro.Id, &libro.Nombre, &libro.Descripcion, &libro.Autor, &libro.Editorial, &libro.Fecha)
			libro1 = append(libro1, libro) // la variable temporal se 'mete' al array de tados
		}

		c.JSON(200, libro1)
	})

	r.GET("/libros/:id", func(c *gin.Context) {
		id := c.Param("id")

		var libro Libros // crear estructura donde se guardara el json de usuario
		err := DB.QueryRow("SELECT * FROM Libros WHERE id=?", id).Scan(&libro.Id, &libro.Nombre, &libro.Descripcion, &libro.Autor, &libro.Editorial, &libro.Fecha)
		if err == sql.ErrNoRows {
			c.JSON(404, gin.H{"msg": "libro no encontrado"})
			return
		}
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, libro)
	})

	//-------------------------------------------------------------
	

	r.DELETE("/libros/:id", func(c *gin.Context) {
		id := c.Param("id")

		var libro Libros // crear estructura donde se guardara el json de usuario
		err := DB.QueryRow("DELETE  FROM Libros WHERE id=?", id).Scan(&libro.Id, &libro.Nombre, &libro.Descripcion, &libro.Autor, &libro.Editorial, &libro.Fecha)
		if err == sql.ErrNoRows {
			c.JSON(200, gin.H{"msg": "libro  eliminado"})
			return
		}
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, libro)
	})

	r.Run(":8069") 
}