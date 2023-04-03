CREATE TABLE IF NOT EXISTS user_register_info (
    ID bigserial unique,
    email varchar(100) unique,
    full_name varchar(100),
    login varchar(100) unique,
    password varchar(100)
);

CREATE TABLE IF NOT EXISTS user_friend_list (
    UserID bigint,
    Friends_IDs bigint[],
    Unconfirmed_Friends_Ids bigint[],
    foreign key (UserID) references user_register_info (ID)
);

CREATE TABLE IF NOT EXISTS task_info (
    ID bigserial unique,
    UserID bigint,
    FirstObserver bigint,
    SecondObserver bigint,
    Deadline timestamp,
    Name varchar(100),
    Description text,
    Punish varchar(500),
    FrequencyPeriod varchar(50)[],
    foreign key (UserID) references user_register_info (ID)
);
--
-- insert into user_register_info (login, password) values ('first_log', 'first_pass');
-- insert into user_register_info (login, password) values ('second', 'seco');
--insert into user_register_info (login, password) values ('first_log', 'fwefwef');

select * from user_register_info;
drop table user_register_info cascade;

drop table user_friend_list;

select * from user_friend_list;
insert into user_friend_list(userid) values (1);
insert into user_friend_list(userid) values (2);
insert into user_friend_list(userid, friends_ids) values (1, '{4,5,6}');

update user_friend_list set friends_ids='{1}' where userid=1;