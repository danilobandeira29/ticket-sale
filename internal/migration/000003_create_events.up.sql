create table events(
    id UUID primary key,
    name varchar(255) not null,
    description text,
    date timestamp with time zone not null,
    is_published boolean default false,
    total_spots int default 0,
    total_spots_reserved int default 0,
    partner_id UUID not null references partners(id) on delete cascade
);

create table event_sections(
    id UUID primary key,
    name varchar(255) not null,
    description text,
    is_published boolean default false,
    total_spots int default 0,
    total_spots_reserved int default 0,
    price numeric default 0,
    event_id UUID not null references events(id) on delete cascade
);

create table event_spots(
    id UUID primary key,
    event_section_id UUID not null references event_sections(id),
    location text not null,
    is_published boolean default false,
    is_reserved boolean default false
);
