create table users (
  id serial primary key,
  auth0 varchar not null unique,
  email varchar not null unique,
  phone varchar not null,
  locale varchar not null
);

insert into users (auth0, email, phone, locale) values
  ('foo', 'test@gmail.com', '+1234567890', 'en')
;