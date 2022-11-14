create table user_wallets
(
    user_id int not null,
    balance int not null default 0 check ( balance >= 0 ),

    primary key (user_id)
);

create table services
(
    id   int generated always as identity,
    name varchar(255) not null,
    url  varchar(255) default null,

    primary key (id)
);

create table transactions
(
    id         int generated always as identity,
    user_id    int not null,
    service_id int,
    order_id   int not null,
    amount     int not null,

    primary key (id),
    foreign key (service_id) references services (id)
);


create table reservations
(
    id         int generated always as identity,
    user_id    int not null,
    service_id int,
    order_id   int not null,
    amount     int not null,

    primary key (id),
    foreign key (service_id) references services (id)
);