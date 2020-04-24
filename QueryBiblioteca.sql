CREATE DATABASE IF NOT EXISTS `Biblioteca`;
USE `Biblioteca`;

CREATE TABLE IF NOT EXISTS Libros (
	id bigint unsigned not null auto_increment,
 nombre varchar(255) not null,
  descripcion varchar(255) not null,
  autor varchar(255) not null,
  Editorial varchar(255) not null,
  fecha_publicacion date not null,
 
  PRIMARY KEY (id)
) ;
DELETE FROM Libros WHERE id=2;

