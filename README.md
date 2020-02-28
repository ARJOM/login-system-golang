# Sistema de Login

Sistema de login web utilizando Golang, HTML e MySQL

Script de criação de usuário:
```SQL
create table usuarios(
	id integer auto_increment,
    username varchar(80) unique,
    password varchar(100) not null,
    primary key(id)
)
```