create table users (
  id serial primary key,
  userid varchar not null unique,
  email varchar not null unique,
  phone varchar not null,
  country varchar not null
);

insert into
  users (userid, email, phone, country)
values
  (
    'dbd813db-1c7c-4c08-b521-e100ab44242b',
    'test@gmail.com',
    '1234567890',
    'US'
  );