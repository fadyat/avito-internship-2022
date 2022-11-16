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

create type transaction_type as enum ('replenishment', 'withdrawal');

create table transactions
(
    id         int generated always as identity,
    user_id    int              not null,
    amount     int              not null,
    type       transaction_type not null,
    created_at timestamp default current_timestamp,

    primary key (id),
    foreign key (user_id) references user_wallets (user_id)
);

create type reservation_status as enum ('pending', 'released', 'cancelled');

create table reservations
(
    id         int generated always as identity,
    user_id    int                not null,
    service_id int,
    order_id   int                not null,
    amount     int                not null,
    status     reservation_status not null,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp,

    primary key (id),
    foreign key (service_id) references services (id),
    foreign key (user_id) references user_wallets (user_id)
);