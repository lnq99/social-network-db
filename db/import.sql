copy Profile from '/home/ql/DB/SocialNetwork/db/csv/profile.csv' delimiter ',' csv header;
copy Post from '/home/ql/DB/SocialNetwork/db/csv/post.csv' delimiter ',' csv header;
copy Comment from '/home/ql/DB/SocialNetwork/db/csv/comment.csv' delimiter ',' csv header;
copy Reaction from '/home/ql/DB/SocialNetwork/db/csv/reaction.csv' delimiter ',' csv header;
copy Relationship from '/home/ql/DB/SocialNetwork/db/csv/relationship.csv' delimiter ',' csv header;
copy Notification from '/home/ql/DB/SocialNetwork/db/csv/notification.csv' delimiter ',' csv header;
copy Album from '/home/ql/DB/SocialNetwork/db/csv/album.csv' delimiter ',' csv header;
copy Photo from '/home/ql/DB/SocialNetwork/db/csv/photo.csv' delimiter ',' csv header;


create temp table tmp_1
as select * from Profile with no data;
copy tmp_1 from '/home/ql/DB/SocialNetwork/db/csv/profile.csv' delimiter ',' csv header;

select email, count(*) as cnt
from tmp_1
group by email
order by cnt desc
limit 10;

drop table tmp_1;



create table tmpPost (
	id			int,
	userId		int		not null	references Profile(id),
	created		date,
	tags		text,
	content		text,
	atchType	atch_t	default 	null,
    atchId		int,
	reaction	int[6],
	nextReact	int		default 	0,
	nextCmt		int		default 	0,
	primary key	(id, userId)
	-- foreign key (account_id) references account(account_id)
);

copy tmpPost from '/home/ql/DB/SocialNetwork/db/csv/post.csv' delimiter ',' csv header;

insert into Post (
	id	,	
	userId	,
	created	,
	tags	,
	content	,
	atchType,
	atchId	,
	reaction,
	nextReact,
	nextCmt	
)
select * from tmpPost;

drop table tmpPost;


create temp table tmp_1
as select * from Profile with no data;
copy tmp_1 from '/home/ql/DB/SocialNetwork/db/csv/profile.csv' delimiter ',' csv header;

select email, count(*) as cnt
from tmp_1
group by email
order by cnt desc
limit 10;

drop table tmp_1;



begin;
create temp table tmp_2 on commit drop
as select * from Relationship with no data;

copy tmp_2 from '/home/ql/DB/SocialNetwork/db/csv/relationship.csv' delimiter ',' csv header;

insert into Relationship
select distinct on (user1, user2) * from tmp_2;
commit;
