create table customers(
    id UUID primary key,
    name varchar(255) not null,
    cpf varchar(14) not null unique
);
