-- drop table user_account;
CREATE TABLE IF NOT EXISTS user_account(
    id SERIAL primary key,
    account varchar(255) unique,
    balance int,
    currency varchar(255),
    created_at timestamp,
    updated_at timestamp default now(),
    disabled int default 0
);
INSERT INTO user_account(account, balance, currency, created_at)
VALUES('bob123', 10000, 'USD', now());
insert into user_account(account, balance, currency, created_at)
values('alice456', 1, 'USD', now());
-- drop table user_payment 
CREATE TABLE IF NOT EXISTS user_payment(
    id SERIAL primary key,
    account varchar(255),
    amount int,
    to_account varchar(255),
    from_account varchar(255),
    direction varchar(255),
    created_at timestamp,
    updated_at timestamp default now(),
    disabled int default 0
);
insert into user_payment(
        account,
        amount,
        to_account,
        from_account,
        direction,
        created_at
    )
values(
        'bob123',
        10000,
        'alice456',
        '',
        'outgoing',
        now()
    );
insert into user_payment(
        account,
        amount,
        to_account,
        from_account,
        direction,
        created_at
    )
values(
        'alice456',
        10000,
        '',
        'bob123',
        'incoming',
        now()
    );