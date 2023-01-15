create table users (
  id serial primary key,
  auth0_id varchar not null,
  email varchar not null,
  phone varchar not null,
  locale varchar not null,
);

insert into users (auth0_id, email, phone, locale) values
  ('foo', 'test@gmail.com', '+1234567890', 'en')
;