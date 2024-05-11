create table users(
	id serial primary key,
	email text default '',
	password text default '',
	firstname text default '',
	lastname text default '',
	createdat text default '',
	roleid integer default 0
);

create table quizes(
	id serial Primary key,
	userid integer ,
	title text default '',
	qcount integer default 0,
	passed integer default 0,
	foreign key (userid) references users(id) on delete cascade
)
create table question(
	quizid integer ,
	id serial primary key,
	question text default '',
	answer text default '',
	orderid serial,
	foreign key (quizid) references quizes(id) on delete cascade
);
create table variants(
	questionid integer,
	variant text default '',
	orderid serial,
	foreign key (questionid) references question(id) on delete cascade
);
create table results(
	userid integer ,
	quizid integer,
	answer text default '',
	ball int default 0,
	Primary key(userid , quizid),
	foreign key (quizid) references quizes(id) on delete cascade,
	foreign key (userid) references users(id) on delete cascade
);

create table quizaccess(
	quizid int ,
	userid int,
	Primary key(userid , quizid),
	foreign key (userid) references users(id) on delete cascade,
	foreign key (quizid) references quizes(id) on delete cascade
);