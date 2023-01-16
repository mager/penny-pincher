create table budgets (
    id serial primary key,
    userID varchar not null unique,
    name varchar not null,
    amount int not null,
    year int not null,
    month int null,
    created_at timestamp not null default now(),
    updated_at timestamp not null default now()
);

insert into
    budgets (userID, name, amount, year, month)
values
    (
        'test',
        'Groceries',
        10000,
        2018,
        1
    );