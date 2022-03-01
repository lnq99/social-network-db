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
    select * from (
        select id, mutual,
            case type
                when 'request' then 'follow'
                else type
            end
        from t left join Relationship r
        on r.user1 = u and r.user2 = id
    union
        select id, mutual, type
        from t left join Relationship r
        on r.user1 = id and r.user2 = u
    ) as tb
    order by type, mutual desc
)   
select jsonb_agg(jsonb_build_object('id', id, 'mutual', mutual, 'type', type)) from rel;
$$ language sql;

--------------------------------------------------
create or replace function reaction_update()
returns trigger as
$$
declare
    o int; n int; postId int; r int[];
begin
    if (old is null) then
        postId := new.postId;
    else
        postId := old.postId;
    end if;

    r := (select reaction from Post where id = postId);
    raise notice '%', r;
    o := (select array_position(array['like','love','haha','wow','sad','angry'], old.type::text));
    n := (select array_position(array['like','love','haha','wow','sad','angry'], new.type::text));

    if (o is null) then     -- raise notice 'insert';
        r[n] := r[n] + 1;
    elsif (n is null) then  -- raise notice 'delete';
        r[o] := r[o] - 1;
    else                    -- raise notice 'update';
        r[o] := r[o] - 1;
        r[n] := r[n] + 1;
    end if;

    update Post set reaction = r where id = postId;
    raise notice '%', r;
    return new;
end;
$$ language plpgsql;

create trigger reaction_type_update
after insert or update or delete
on Reaction for each row
execute function reaction_update();

--------------------------------------------------
create or replace function check_name()
returns trigger as
$$
begin
    if (new.name ~ '.*[0-9!"#$%&\*+,-./:;<=>?@]+.*')
    then
        raise exception 'Invalid name!';
    else
        return new;
    end if;
end;
$$ language plpgsql;

create trigger profile_insert
before insert on Profile for each row
execute function check_name();