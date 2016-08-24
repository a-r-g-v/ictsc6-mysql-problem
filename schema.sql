drop table if exists message;

create table message(id integer primary key AUTO_INCREMENT, name text, body text) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ;
insert into message(name ,body) values("社内伝言板", "社内伝言板です");

