create table users (
  id serial primary key,
  userid varchar not null unique,
  name varchar not null,
  email varchar not null unique,
  phone varchar not null,
  country varchar not null
);

insert into
  users (userid, name, email, phone, country)
values
  (
    'dbd813db-1c7c-4c08-b521-e100ab44242b',
    'Test User',
    'test@gmail.com',
    '1234567890',
    'US'
  );