-- create  table students(id serial primary key, name varchar,surname varchar,phoneNumber varchar,email varchar,address varchar,password varchar);

create unique index product_id_3 on dodi (surname, name, email);

create unique index product_id_2 on dodi ( name, surname, email);

create unique index product_id_1 on dodi (email, surname, name);

create unique index product_2 on dodi (address, phoneNumber);

create unique index product_2 on dodi (password, name);


explain (analyse )
select * from dodi where surname = 'TCwdSEX@ubwUhXM.net' and name  = 'Forrest' and email = 'TCwdSEX@ubwUhXM.net';

explain (analyse )
select * from dodi where name  = 'Forrest' and surname = 'TCwdSEX@ubwUhXM.net' and email = 'TCwdSEX@ubwUhXM.net';

explain (analyse )
select * from dodi where email = 'TCwdSEX@ubwUhXM.net' and surname = 'TCwdSEX@ubwUhXM.net' and name  = 'Forrest';

explain (analyse )
select * from dodi where address = 'nil' phoneNumber = 'phone_number';

explain (analyse )
select * from dodi where name = 'Lavada' password = 'XMHNywFswkcNtbLGVsbuprPdcheCvTiTqmlcMNbRNIFyhQHFkb';

