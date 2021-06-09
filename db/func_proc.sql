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


--------------------------------------------------
create or replace function friends(urid int)
returns table (id int, name text, avatarS text)
as $$
select id, name, avatarS
from Profile
where id in (
    select user2
    from Relationship
    where user1 = urid
        and type = 'friend'
)
$$ language sql;

create or replace function friends_json(urid int)
returns jsonb as $$
select jsonb_agg(t) from friends(urid) as t;
$$ language sql;

select friends_json(1);
drop function friends_json;


--------------------------------------------------
create or replace function mutual_friends(u1 int, u2 int)
returns int[]
as $$
select array(
    select R1.user2 friend
    from Relationship R1
    join (
        select user2 from Relationship where user1 = u2 and type = 'friend'
    ) as R2 on R1.user2 = R2.user2
    where user1 = u1 and type = 'friend'
)
$$ language sql;

select mutual_friends(1, 2);
drop function mutual_friends;


explain analyse
select R1.user2 friend
from Relationship R1
join (
    select user2 from Relationship where user1 = 2 and type = 'friend'
) as R2 on R1.user2 = R2.user2
where user1 = 1 and type = 'friend';

explain analyse
select user2 from Relationship where user1 = 1 and type = 'friend'
intersect
select user2 from Relationship where user1 = 2 and type = 'friend';


--------------------------------------------------
create or replace function search_name(id int, pattern text)
returns jsonb as
$$
with
t as (
    select id, cardinality(mutual_friends(1, id)) as mutual
    from Profile
    where lower(name) like format('%%%s%%', pattern)
    and id not in (
        select id
        from Relationship
        where user2 = 1 and type = 'block'
    )
),
rel as (
    (
        select id, mutual,
            case type
                when 'request' then 'follow'
                else type
            end
        from t left join Relationship r
        on r.user1 = 1 and r.user2 = id
    )
    union
    (
        select id, mutual, type
        from t left join Relationship r
        on r.user1 = id and r.user2 = 1
    )
)   
select jsonb_agg(jsonb_build_object('id', id, 'mutual', mutual, 'type', type)) from rel;
$$
language sql;

select search_name(10, '%ok%');
drop function search_name;


-- https://stackoverflow.com/questions/24006291/postgresql-return-result-set-as-json-array/24006432
-- https://stackoverflow.com/questions/25678509/postgres-recursive-query-with-row-to-json