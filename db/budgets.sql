create table budgets (
    id serial primary key,
    userid varchar not null unique,
    name varchar not null,
    amount int not null,
    year int not null,
    month int null,
    createdat timestamp not null default now(),
    updatedat timestamp not null default now()
);

insert into
    budgets (userid, name, amount, year, month)
values
    (
        'google-oauth2|109298383726697703588',
        'Groceries',
        500,
        2018,
        1
    );