create or replace function feed(urid int, lim int default 10, offs int default 0)
returns int[]
as $$
select array(
    select id
    from Post
    where userId in (
        select user2
        from Relationship
        where user1 = urid
            and type = 'friend'
    )
    order by created desc
    limit lim offset offs
)
$$ language sql;

select feed(1);
select feed(1, 10, 10);
drop function feed;


create or replace function friends(urid int)
returns table (id int, name text, avatarS text)
as $$
select id, name, avatarS
from Profile
where id in (
    select user2
    from Relationship
    where user1 = 1
        and type = 'friend'
)
$$ language sql;

create or replace function friends_json(urid int)
returns jsonb as $$
select jsonb_agg(t) from friends(urid) as t;
$$ language sql;

select friends_json(1);


-- https://stackoverflow.com/questions/24006291/postgresql-return-result-set-as-json-array/24006432
-- https://stackoverflow.com/questions/25678509/postgres-recursive-query-with-row-to-json