create table book
(
    id          int primary key,
    name        varchar,
    page        int,
    author_name varchar
);


create table author
(
    id   int primary key,
    name varchar         
);

 insert into book(id,name,page,author_name) values(7,'dodi6',120006,'KHdodi6'),(6,'dodi6',120006,'KHdodi6'),(5,'dodi3',120003,'KHdodi3'),(4,'dodi4',120004,'KHdodi4');

 insert into author(id,name) values(4,'dodi6'),(5,'dodi6'),(6,'dodi3'),(7,'dodi4');


SELECT book.id,book.name                             
FROM book  
LEFT JOIN author
ON book.id = author.id;
