CREATE TYPE type_account AS ENUM ('balance', 'reserved');

create table accounts (
	id_account bigserial unique, 
	id_user bigint,
	amount decimal check (amount >= 0),
	type_acc type_account,
	primary key (id_account)
);

create table services (
	id_service bigserial primary key,
	price decimal check (price >= 0)
);

create table orders (
	id_order bigserial primary key,
	id_service bigserial not null,
	account bigserial not null,
	foreign key (id_service) references services (id_service),
	foreign key (account) references accounts (id_account)
);

create type type_transaction as enum ('deposit', 'pay', 'transit');

create table transactions (
	id_transaction bigserial primary key,
	price decimal,
	ta_date timestamp,
	ta_type type_transaction,
	from_acc bigserial references accounts (id_account),
	to_acc bigserial references accounts (id_account)
	
);