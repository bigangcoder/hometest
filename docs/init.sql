-- drop table user_account;
create table user_account(
    id int auto_increment primary key,
    account varchar(255),
    balance int,
    currency varchar(255),
    created_at datetime,
    updated_at timestamp default now(),
    disabled int default 0
) engine = innodb default charset = utf8;
insert into user_account(account, balance, currency, created_at)
values("bob123", 10000, "USD", now());
insert into user_account(account, balance, currency, created_at)
values("alice456", 1, "USD", now());
-- drop table user_payment 
create table user_payment(
    id int auto_increment primary key,
    account varchar(255),
    amount int,
    to_account varchar(255),
    from_account varchar(255),
    direction varchar(255),
    created_at datetime,
    updated_at timestamp default now(),
    disabled int default 0
) engine = innodb default charset = utf8;
insert into user_payment(
        account,
        amount,
        to_account,
        from_account,
        direction,
        created_at
    )
values(
        "bob123",
        10000,
        "alice456",
        "",
        "outgoing",
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
        "alice456",
        10000,
        "",
        "bob123",
        "incoming",
        now()
    );