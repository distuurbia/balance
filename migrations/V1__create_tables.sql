-- Create new profiles table
create table balances (
	balanceid uuid,
	profileid uuid,
	tsTime timestamp,
	amount double precision,
	primary key (balanceid)
);
