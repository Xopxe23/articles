CREATE TABLE users (
    id serial not null unique,
    name varchar(255) not null,
    surname varchar(255) not null,
    email varchar(255) not null unique,
    password varchar(255) not null
);

CREATE TABLE refresh_tokens (
    id serial not null unique,
    user_id int references users (id) on delete cascade not null,
    token varchar(255) not null unique,
    expires_at timestamp not null
);