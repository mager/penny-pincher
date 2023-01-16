create table users (
  id serial primary key,
  userID varchar not null unique,
  email varchar not null unique,
  phone varchar not null,
  locale varchar not null
);

insert into
  users (userID, email, phone, locale)
values
  ('test', 'test@gmail.com', '+1234567890', 'en');