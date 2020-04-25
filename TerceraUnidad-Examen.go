package main

import (
	"bufio"                            // Leer líneas incluso si tienen espacios
	"database/sql"                     // Interactuar con bases de datos
	"fmt"                              // Imprimir mensajes y esas cosas
	_"mysql-master"// La librería que nos permite conectar a MySQL
	"os"                               // El búfer, para leer desde la terminal con os.Stdin
)

type Libro struct {
	Nombre, Descripcion, Autor, Editorial, FechaPublicacion string
	Id                                   int
}

func obtenerBaseDeDatos() (db *sql.DB, e error) {
	usuario := "root"
	pass := "12345"
	host := "tcp(127.0.0.1:3306)"
	nombreBaseDeDatos := "biblioteca"
	// Debe tener la forma usuario:contraseña@protocolo(host:puerto)/nombreBaseDeDatos
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@%s/%s", usuario, pass, host, nombreBaseDeDatos))
	if err != nil {
		return nil, err
	}
	return db, nil
}

func main() {
	menu := `¿Qué deseas hacer?
[1] -- Agregar
[2] -- Mostrar
[3] -- Actualizar
[4] -- Eliminar
[5] -- Salir
----->	`
	var eleccion int
	var l Libro
	for eleccion != 5 {
		fmt.Print(menu)
		fmt.Scanln(&eleccion)
		scanner := bufio.NewScanner(os.Stdin)
		switch eleccion {
		case 1:
			fmt.Println("Ingresa el nombre del libro:")
			if scanner.Scan() {
				l.Nombre = scanner.Text()
			}
			fmt.Println("Ingresa la descripcion del libro:")
			if scanner.Scan() {
				l.Descripcion = scanner.Text()
			}
			fmt.Println("Ingresa el autor:")
			if scanner.Scan() {
				l.Autor = scanner.Text()
			}
			fmt.Println("Ingresa la editorial:")
			if scanner.Scan() {
				l.Editorial = scanner.Text()
			}
			fmt.Println("Ingresa la fecha de publicación:")
			if scanner.Scan() {
				l.FechaPublicacion = scanner.Text()
			}
			err := insertar(l)
			if err != nil {
				fmt.Printf("Error insertando: %v", err)
			} else {
				fmt.Println("Insertado correctamente")
			}


		case 2:
			libros, err := obtenerLibros()
			if err != nil {
				fmt.Printf("Error obteniendo los libros: %v", err)
			} else {
				for _, libro := range libros {
					fmt.Println("====================")
					fmt.Printf("Id: %d\n", libro.Id)
					fmt.Printf("Nombre: %s\n", libro.Nombre)
					fmt.Printf("Descripcion: %s\n", libro.Descripcion)
					fmt.Printf("Autor: %s\n", libro.Autor)
					fmt.Printf("Editorial: %s\n", libro.Editorial)
					fmt.Printf("Fecha de ´publicación: %s\n", libro.FechaPublicacion)
				}
			}
	// -------------------------------------------------
			
		case 3:
			fmt.Println("Ingresa el id del libro:")
			fmt.Scanln(&l.Id)
			fmt.Println("Ingresa el nuevo nombre:")
			if scanner.Scan() {
				l.Nombre = scanner.Text()
			}
			fmt.Println("Ingresa la nueva descripción:")
			if scanner.Scan() {
				l.Descripcion = scanner.Text()
			}
			fmt.Println("Ingresa el nuevo autor:")
			if scanner.Scan() {
				l.Autor = scanner.Text()
			}
			fmt.Println("Ingresa la nueva editorial:")
			if scanner.Scan() {
				l.Editorial = scanner.Text()
			}
			fmt.Println("Ingresa la nueva fecha de publicación:")
			if scanner.Scan() {
				l.Editorial = scanner.Text()
			}
			err := actualizar(l)
			if err != nil {
				fmt.Printf("Error actualizando: %v", err)
			} else {
				fmt.Println("Actualizado correctamente")
			}

		case 4:
			// ------------
			fmt.Println("Ingresa el ID del contacto que deseas eliminar:")
			fmt.Scanln(&l.Id)
			err := eliminar(l)
			if err != nil {
				fmt.Printf("Error eliminando: %v", err)
			} else {
				fmt.Println("Eliminado correctamente")
			}
		}
	}
}

func eliminar(l Libro) error {
	db, err := obtenerBaseDeDatos()
	if err != nil {
		return err
	}
	defer db.Close()

	sentenciaPreparada, err := db.Prepare("DELETE FROM libros WHERE id = ?")
	if err != nil {
		return err
	}
	defer sentenciaPreparada.Close()

	_, err = sentenciaPreparada.Exec(l.Id)
	if err != nil {
		return err
	}
	return nil
}

func insertar(l Libro) (e error) {
	db, err := obtenerBaseDeDatos()
	if err != nil {
		return err
	}
	defer db.Close()

	
	sentenciaPreparada, err := db.Prepare("INSERT INTO libros (nombre, descripcion, autor, editorial, fecha_publicacion) VALUES(?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer sentenciaPreparada.Close()

	_, err = sentenciaPreparada.Exec(l.Nombre, l.Descripcion, l.Autor, l.Editorial, l.FechaPublicacion)
	if err != nil {
		return err
	}
	return nil
}

func obtenerLibros() ([]Libro, error) {
	libros := []Libro{}
	db, err := obtenerBaseDeDatos()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	filas, err := db.Query("SELECT id, nombre, descripcion, autor, editorial, fecha_publicacion  FROM libros")

	if err != nil {
		return nil, err
	}
	
	defer filas.Close()

	var l Libro

	for filas.Next() {
		err = filas.Scan(&l.Id, &l.Nombre, &l.Descripcion, &l.Autor, &l.Editorial,  &l.FechaPublicacion)
		
		if err != nil {
			return nil, err
		}
	
		libros = append(libros, l)
	}
	
	return libros, nil
}

func actualizar(l Libro) error {
	db, err := obtenerBaseDeDatos()
	if err != nil {
		return err
	}
	defer db.Close()

	sentenciaPreparada, err := db.Prepare("UPDATE libros SET nombre = ?, descripcion = ?, autor = ?, editorial = ?, fecha_publicacion = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer sentenciaPreparada.Close()

	_, err = sentenciaPreparada.Exec(l.Nombre, l.Descripcion, l.Autor, l.Editorial, l.FechaPublicacion, l.Id)
	return err 
}