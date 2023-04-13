CREATE TABLE IF NOT EXISTS user_register_info
(
    ID        bigserial unique,
    email     varchar(100) unique,
    full_name varchar(100),
    login     varchar(100) unique,
    password  varchar(100)
);

CREATE TABLE IF NOT EXISTS user_friend_list
(
    UserID                  bigint,
    Friends_IDs             bigint[],
    Unconfirmed_Friends_Ids bigint[],
    foreign key (UserID) references user_register_info (ID)
);

CREATE TABLE IF NOT EXISTS user_notification
(
    UserID           bigint,
    NotificationFrom bigint,
    Message          text,
    foreign key (UserID) references user_register_info (ID),
    foreign key (NotificationFrom) references user_register_info (ID)
);

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

CREATE TABLE IF NOT EXISTS days
(
    taskID         bigint,
    secondObserverID bigint,
    firstObserverID bigint,
    day            varchar(50),
    status varchar(50),
    foreign key (taskID) references task_info (ID)
);


CREATE TABLE IF NOT EXISTS user_tasks
(
    user_id            bigint,
    task_ids           bigint[],
    observer_tasks_ids bigint[],
    foreign key (user_id) references user_register_info (ID)
);
