create or replace function feed(urid int, lim int default 10, offs int default 0)
returns int[]
as $$
select array(
    select id
    from Post
    where userId = urid or userId in (
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
create or replace function friends_json(urid int)
returns jsonb as $$
with t as (select id, name, avatarS
from Profile
where id in (
    select user2
    from Relationship
    where user1 = urid
        and type = 'friend'
))
select jsonb_agg(t) from t;
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
create or replace function search_name(u int, pattern text)
returns jsonb as
$$
with
t as (
    select id, cardinality(mutual_friends(u, id)) as mutual
    from Profile
    where lower(name) like format('%%%s%%', lower(pattern))
    and id not in (
        select id
        from Relationship
        where user2 = u and type = 'block'
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
        on r.user1 = u and r.user2 = id
    )
    union
    (
        select id, mutual, type
        from t left join Relationship r
        on r.user1 = id and r.user2 = u
    )
)   
select jsonb_agg(jsonb_build_object('id', id, 'mutual', mutual, 'type', type)) from rel;
$$ language sql;

explain analyse
select search_name(1, '%jo%');
drop function search_name(int, text);


--------------------------------------------------
create or replace function search_name(u int, pattern text)
returns jsonb as
$$
with
t as (
    select id, (
        select count(*)
        from Relationship R1
        join (
            select user2 from Relationship where user1 = u and type = 'friend'
        ) as R2 on R1.user2 = R2.user2
        where user1 = id and type = 'friend'
    ) as mutual
    from Profile
    where lower(name) like format('%%%s%%', lower(pattern))
    and id not in (
        select user1
        from Relationship
        where user2 = u and type = 'block'
    )
),
rel as (
    select * from
    (
        select id, mutual,
            case type
                when 'request' then 'follow'
                else type
            end
        from t left join Relationship r
        on r.user1 = u and r.user2 = id
        
        -- select id, mutual,
        --     case type
        --         when 'request' then 'follow'
        --         else type
        --     end
        -- from Relationship r right join t on r.user2 = t.id
        -- where r.user1 = u
    union
        select id, mutual, type
        from t left join Relationship r
        on r.user1 = id and r.user2 = u

        -- select id, mutual, type
        -- from Relationship r right join t on r.user1 = t.id
        -- where r.user2 = u
    ) as tb
    order by type, mutual desc
)   
select jsonb_agg(jsonb_build_object('id', id, 'mutual', mutual, 'type', type)) from rel;
$$ language sql;


-- https://stackoverflow.com/questions/24006291/postgresql-return-result-set-as-json-array/24006432
-- https://stackoverflow.com/questions/25678509/postgres-recursive-query-with-row-to-json


explain analyse
select search_name(1, '%jo%');


[

 "QUERY PLAN": "Result  (cost=0.00..0.26 rows=1 width=32) (actual time=21.492..21.493 rows=1 loops=1)"
,

 "QUERY PLAN": "Planning Time: 3.091 ms"
,

 "QUERY PLAN": "Execution Time: 21.707 ms"

]

[
  {
    "QUERY PLAN": "Result  (cost=0.00..0.26 rows=1 width=32) (actual time=15.295..15.296 rows=1 loops=1)"
  },
  {
    "QUERY PLAN": "Planning Time: 3.033 ms"
  },
  {
    "QUERY PLAN": "Execution Time: 15.407 ms"
  }
]
--------------------------------------------------
create or replace function n_mutual_friends(u1 int, u2 int)
returns bigint
as $$
select count(*)
from Relationship R1
join (
    select user2 from Relationship where user1 = u2 and type = 'friend'
) as R2 on R1.user2 = R2.user2
where user1 = u1 and type = 'friend'
$$ language sql;

select n_mutual_friends(1, 2);
drop function n_mutual_friends;