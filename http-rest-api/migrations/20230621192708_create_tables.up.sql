create table if not exists users (
	id_user bigserial primary key,
	username varchar,
	surname varchar,
	email varchar unique,
	pass varchar
);

create table if not exists accounts (
	id_account bigserial, 
	id_user bigserial references users (id_user) not null,
	amount decimal check (amount >= 0),
	primary key (id_account)
);


create table if not exists transactions (
	id_transaction bigserial primary key,
	price decimal not null,
	ta_date timestamp not null,
	from_acc bigserial references accounts (id_account),
	to_acc bigserial references accounts (id_account)
);