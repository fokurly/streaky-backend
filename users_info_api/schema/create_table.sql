CREATE TABLE IF NOT EXISTS user_register_info (
    ID bigserial unique,
    email varchar(100) unique,
    login varchar(100) unique,
    password varchar(100)
);

CREATE TABLE IF NOT EXISTS user_friend_list (
    UserID bigint,
    Friends_IDs bigint[],
    foreign key (UserID) references user_register_info (ID)
);

CREATE TABLE IF NOT EXISTS task_info (
    ID bigserial unique,
    UserID bigint,
    Deadline timestamp,
    Name varchar(100),
    Description text,
    Punish varchar(500),
    FrequencyPeriod varchar(50)[],
    foreign key (UserID) references user_register_info (ID)
);

insert into user_register_info (login, password) values ('first_log', 'first_pass');
insert into user_register_info (login, password) values ('second', 'seco');
--insert into user_register_info (login, password) values ('first_log', 'fwefwef');

select * from user_register_info;
--drop table user_register_info cascade;