create database dars_db

\c dars_db;

create table students(id serial primary key, name varchar, age int, phone_number varchar(13));

\d students

insert into students(name, age, phone_number) values('Ahmat', 18, '+998999635522');
insert into students(name, age, phone_number) values('Mehmet', 22, '+998999634455'), ('Abdulloh', 15, '+998886355272');
insert into students(name, age, phone_number) values('Bittabola', 20, '+998999634455'), ('Nmadur', 31, '+998858355272')('kimdur', 52, '+998955635522');
insert into students(name, age, phone_number) values('Bittabola', 20, '+998999634455'), ('Nmadur', 31, '+998858355272'),('kimdur', 52, '+998955635522');
insert into students(name, age, phone_number) values('Mehmet', 22, '+998999634455');
insert into students(name, age, phone_number) values('Avaz', 17, '+998999159455');

update students set age = 25 where age > 30;

select * from students;

delete from students where name = 'Nmadur';

select * from students;

select age, count(*) same_age from students group by age;

select * from students order by age;

select * from students order by age desc;

select * from students inner join students on students.id = students.age;

create table marks(id serial primary key, mark int, name_student varchar, id_students int);

insert into marks(mark, id_students) values(2, 1), (2, 2), (5, 3), (4, 4), (3, 8), (5, 7);

select * from students inner join marks on students.id = marks.id_students;

select * from students left join marks on students.id = marks.id_students;

select * from students right join marks on students.id = marks.id_students;

select * from students full outer join marks on students.id = marks.id_students;

select * from students right join marks on students.id = marks.id_students;

select * from students cross join marks;

--tugadi 
