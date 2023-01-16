create table users (
  id serial primary key,
  userid varchar not null unique,
  email varchar not null unique,
  phone varchar not null,
  locale varchar not null
);

insert into
  users (userid, email, phone, locale)
values
  (
    'google-oauth2|109298383726697703588',
    'test@gmail.com',
    '+1234567890',
    'en'
  );