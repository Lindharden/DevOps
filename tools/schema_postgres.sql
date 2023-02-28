drop table if exists user_acc;


create table user_acc (user_id serial primary key, username VARCHAR not null, email VARCHAR not null, pw_hash VARCHAR not null);


drop table if exists follower;


create table follower (who_id integer, whom_id integer);


drop table if exists message;


create table message (message_id serial primary key, author_id integer not null, text VARCHAR not null, pub_date integer, flagged integer);

