create database SocialNetwork;
drop database if exists "Social Network";

create type gender_t as enum('M','F');
create type notif_t as enum('like','cmt','request','accept');
create type relation_t as enum('friend','block','request');
create type atch_t as enum('photo','video','poll');
create type react_t as enum('like','love','haha','wow','sad','angry');

drop type gender_t;
drop type notif_t;
drop type relation_t;
drop type atch_t;
drop type react_t;

create table Profile (
	id			serial primary key,
	name		varchar(32)	not null,
	gender		gender_t,
	birthday	date,
	email		text		not null	unique,
	phone		decimal(13)	unique,
	salt		char(8)		not null,
	hash		char(40)	not null,
	created		date,
	intro		text,
	avatarS		text,
	avatarL		text,
	nextPost	int		default 	0,
	nextNotif	int		default 	0,
	nextPhoto	int		default 	0
);

create table Post (
	id			int,
	userId		int		not null	references Profile(id),
	created		date,
	modified	date	default 	null,
	tags		text,
	content		text,
	atchType	atch_t	default 	null,
	atchId		int,
	atchUrl		text	default 	null,
	reaction	int[6],
	nextReact	int		default 	0,
	nextCmt		int		default 	0,
	primary key	(id, userId)
	-- foreign key (account_id) references account(account_id)
);

create table Comment (
	id			int,
	userId		int,
	postId		int,
	parentId	int,
	content		text,
	created		date,
	replyCount	int,
	primary key	(id, userId, postId),
	foreign key (userId, postId) references Post(userId, id)
);

create table Reaction (
	userId 		int		not null	references Profile(id),
	authorId 	int,
	postId 		int,
	type 		react_t	default 	'like',
	primary key	(userId, authorId, postId),
	foreign key (authorId, postId) references Post(userId, id)
);

create table Relationship (
	user1		int		not null	references Profile(id),
	user2		int		not null	references Profile(id),
	created		date,
	type		relation_t,
	other		text	default 	null,
	primary key	(user1, user2)
);

create table Notification (
	id			int,
	userId		int		not null	references Profile(id),
	type 		notif_t,
	created		date,
	fromUserId 	int		references Profile(id),
	postId 		int		default 	null,
	cmtId 		int		default 	null,
	primary key	(id, userId)
);

create table Album (
	id			int,
	userId 		int		not null	references Profile(id),
	descr 		text	default 	null,
	created 	date,
	primary key	(id, userId)
);

create table Photo (
	id			int,
	userId 		int,
	albumId 	int,
	url 		text,
	created 	date,
	primary key	(id, userId),
	foreign key (userId, albumId) references Album(userId, id)
);

select * from Profile;
select * from Post;
select * from Comment;
select * from Reaction;
select * from Relationship;
select * from Notification;
select * from Album;
select * from Photo;

drop table if exists Profile;
drop table if exists Post;
drop table if exists Comment;
drop table if exists Reaction;
drop table if exists Relationship;
drop table if exists Notification;
drop table if exists Album;
drop table if exists Photo;
