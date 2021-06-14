select * from Reaction where userId = 1 and postId = 4078;
select * from Reaction where postId = 4078;

insert into Reaction
values (1, 4078, 'sad')
on conflict (userId, postId)
do update set type = 'like';

insert into Reaction
values (2, 4078, 'haha')
on conflict (userId, postId)
do update set type = 'sad';

delete from Reaction
where userId in (1,2) and postId = 4078;

select reaction from Post where id = 4078;


update Post
set reaction = array[2,1,0,0,1,0]
where id = 4078;



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

    if (o is null) then
        raise notice 'insert';
        r[n] := r[n] + 1;
    elsif (n is null) then
        raise notice 'delete';
        r[o] := r[o] - 1;
    else
        raise notice 'update';
        r[o] := r[o] - 1;
        r[n] := r[n] + 1;
    end if;

    update Post set reaction = r where id = postId;
    raise notice '%', r;
    return new;
end;
$$ language plpgsql;

drop function reaction_update;


create trigger reaction_type_update
after insert or update or delete
on Reaction
for each row
execute function reaction_update();

drop trigger reaction_type_update on Reaction;


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
$$
language plpgsql;

create trigger profile_insert
before insert on Profile
for each row
execute function check_name();

select * from profile
order by id desc limit 1;

explain analyse
select 'first _name' ~ '.*[0-9!"#$%&\()*+,-./:;<=>?@]+.*';

drop trigger profile_insert on profile;
