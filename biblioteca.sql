create database if not exists biblioteca;
use biblioteca;

create table if not exists libros(
    id int(12) not null auto_increment primary key,
    titulo varchar(100) not null,
    descripcion varchar(450) not null,
    autor varchar(200) not null,
    editorial varchar(200) not null,
    fechapublicacion varchar(100)
);