create table Roles
(
    id        serial primary key not null,
    Role_name varchar(255)       not null unique
);

create table Users
(
    id            serial primary key not null,
    Name          varchar(255)       not null,
    Surname       varchar(255)       not null,
    Username      varchar(255)       not null unique,
    Email         varchar(255)       not null unique,
    Password_hash varchar(255)       not null,
    Role_id       int references Roles (id) on delete cascade
);

INSERT INTO Roles (Role_name) VALUES ('admin');