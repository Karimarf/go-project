create table room
(
    id          int primary key auto_increment,
    name        text      not null,
    capacity    int       not null,
    description text      not null,
    created_at  timestamp not null default now(),
    updated_at  timestamp not null default now()
);

create table reservation
(
    id          int primary key auto_increment,
    room_id     int not null,
    start_date  date not null,
    end_date    date not null,
    created_at  timestamp not null default now(),
    updated_at  timestamp not null default now(),
    foreign key (room_id) references room (id)
);