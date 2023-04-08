CREATE TABLE IF NOT EXISTS user_register_info
(
    ID        bigserial unique,
    email     varchar(100) unique,
    full_name varchar(100),
    login     varchar(100) unique,
    password  varchar(100)
);

---select * from user_register_info;
---SELECT id FROM user_register_info WHERE login='fokuryl' and password='d8578edf8458ce06fbc5bb76a58c5ca4';

CREATE TABLE IF NOT EXISTS user_friend_list
(
    UserID                  bigint,
    Friends_IDs             bigint[],
    Unconfirmed_Friends_Ids bigint[],
    foreign key (UserID) references user_register_info (ID)
);

CREATE TABLE IF NOT EXISTS user_notification
(
    UserID                  bigint,
    NotificationFrom             bigint,
    Message text,
    foreign key (UserID) references user_register_info (ID),
    foreign key (NotificationFrom) references user_register_info (ID)
);

---drop table user_notification;

CREATE TABLE IF NOT EXISTS task_info
(
    ID              bigserial unique,
    UserID          bigint,
    FirstObserver   bigint,
    SecondObserver  bigint,
    Name            varchar(100),
    Description     text,
    Punish          varchar(500),
    FrequencyPeriod varchar(50)[],
    Status          varchar(30),
    StartDate       varchar(50),
    EndDate         varchar(50),
    foreign key (UserID) references user_register_info (ID)
);

-- select * from task_info;
-- SELECT COUNT(*) FROM task_info;
--
-- drop table task_info;
--
-- select * from task_info;
--
-- insert into task_info(userid) values (1);
-- insert into task_info(userid) values (3);
-- insert into task_info(userid) values (5);
-- insert into task_info(userid) values (2);
--
-- insert into user_tasks(user_id) values (2);
-- insert into user_tasks(user_id) values (3);
-- insert into user_tasks(user_id) values (1);
CREATE TABLE IF NOT EXISTS user_tasks
(
    user_id            bigint,
    task_ids           bigint[],
    observer_tasks_ids bigint[],
    foreign key (user_id) references user_register_info (ID)
);

-- drop table user_tasks;

--
-- select * from user_tasks;
--
-- insert into user_register_info (login, password) values ('first_log', 'first_pass');
-- insert into user_register_info (login, password) values ('second', 'seco');
--insert into user_register_info (login, password) values ('first_log', 'fwefwef');

-- select * from user_register_info;
-- select * from user_friend_list;
-- drop table user_register_info cascade;
--
-- drop table user_friend_list;
--
-- select * from user_friend_list;
-- insert into user_friend_list(userid) values (1);
-- insert into user_friend_list(userid) values (2);
-- insert into user_friend_list(userid, friends_ids) values (1, '{4,5,6}');
--
-- update user_friend_list set friends_ids='{1}' where userid=1;