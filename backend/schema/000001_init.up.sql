create type gender_t as enum('M','F');
create type notif_t as enum('like','cmt','request','accept');
create type relation_t as enum('friend','block','request', 'follow');
create type atch_t as enum('photo','video','poll','none');
create type react_t as enum('like','love','haha','wow','sad','angry');


create table Profile (
	id			serial 		primary key,
	name		varchar(32)	not null,
	gender		gender_t,
	birthdate	date,
	email		text		not null	unique,
	phone		decimal(13)	default 0,
	salt		char(8)		not null,
	hash		char(28)	not null,
	created		timestamp			default now(),
	intro		text	default		'',
	avatarS		text	default		'',
	avatarL		text	default		'',
	postCount	int		default 	0,
	photoCount	int		default 	0
);

create table Post (
	id			serial	primary key,
	userId		int		not null	references Profile(id)	 on delete cascade,
	created		timestamp			default now(),
	tags		text	default 	'',
	content		text,
	atchType	atch_t	default 	'none',
	atchId		int		default 	0,
	atchUrl		text	default 	'',
	reaction	int[6],
	cmtCount	int		default 	0
);

create table Comment (
	id			bigserial			primary key,
	userId		int		not null	references Profile(id)	 on delete cascade,
	postId		int		not null	references Post(id)		 on delete cascade,
	parentId	int,
	content		text,
	created		timestamp			default now()
);

create table Reaction (
	userId 		int		not null	references 	Profile(id)	 on delete cascade,
	postId 		int		not null	references 	Post(id)	 on delete cascade,
	type 		react_t	default 	'like',
	primary key	(userId, postId)
);

create table Relationship (
	user1		int		not null	references Profile(id)	 on delete cascade,
	user2		int		not null	references Profile(id)	 on delete cascade,
	created		timestamp			default now(),
	type		relation_t,
	other		text	default 	'',
	primary key	(user1, user2)
);

create table Notification (
	id			bigserial			primary key,
	userId		int		not null	references Profile(id)	 on delete cascade,
	type 		notif_t,
	created		timestamp			default now(),
	fromUserId 	int		not null	references Profile(id)	 on delete cascade,
	postId 		int		default 	0,
	cmtId 		int		default 	0
);

create table Album (
	id			serial	primary key,
	userId 		int		not null	references Profile(id)	 on delete cascade,
	descr 		text	default 	'',
	created 	timestamp			default now()
);

create table Photo (
	id			serial	primary key,
	userId 		int		not null	references Profile(id)	 on delete cascade,
	albumId 	int		not null	references Album(id)	 on delete cascade,
	url 		text,
	created 	timestamp			default now()
);
