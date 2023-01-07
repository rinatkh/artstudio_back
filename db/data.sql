create table public.users
(
    uuid       text primary key         not null default gen_random_uuid(),
    firstname  text                     not null,
    surname    text                     not null,
    middlename text,
    sex        character(1)             not null,
    birth_date timestamp with time zone not null,
    fio        text,
    role       text                     not null,
    image      text                     not null,
    created_at timestamp with time zone not null default now()
);

CREATE
    OR REPLACE FUNCTION fill_fio() RETURNS TRIGGER AS
$$
declare
    t text;
BEGIN
    t = (SELECT CASE NEW.middlename IS NOT NULL
                    WHEN true THEN concat(' ', NEW.middlename)
                    WHEN false THEN ''
                    END
         from Users);
    UPDATE Users
    SET fio = concat(NEW.firstname, ' ', NEW.surname, t)
    WHERE uuid = NEW.uuid;
    RETURN NEW;
END;
$$
    LANGUAGE plpgsql;

CREATE TRIGGER fill_fio
    AFTER INSERT
    ON public.users
    FOR EACH ROW EXECUTE PROCEDURE fill_fio();

alter table users
    owner to postgres;

create table public.departments
(
    id          bigserial primary key not null,
    name        text                  NOT NULL,
    description text,
    image       text                  not null
);

alter table departments
    owner to postgres;

create table public.subjects
(
    id          bigserial primary key not null,
    name        text                  NOT NULL,
    description text,
    teacher_id  text                  NOT NULL REFERENCES users (uuid),
    image       text                  not null
);

alter table subjects
    owner to postgres;

create table public.DepartmentSubjects
(
    department_id bigint NOT NULL REFERENCES departments (id),
    subject_id    bigint NOT NULL REFERENCES subjects (id),
    PRIMARY KEY (department_id, subject_id)
);

create table public.schedules
(
    subject_id bigint                   NOT NULL REFERENCES subjects (id),
    student_id text                     NOT NULL REFERENCES users (uuid),
    PRIMARY KEY (subject_id, student_id),
    when_time  timestamp with time zone not null,
    isPaid     boolean                  not null
);

alter table schedules
    owner to postgres;

create table public.clients
(
    id               bigserial primary key    not null,
    student_id       text                     NOT NULL REFERENCES users (uuid),
    nicknameTelegram text,
    number           text,
    status           int                      not null,
    update_time      timestamp with time zone not null
);

alter table clients
    owner to postgres;

create table public.CardInfo
(
    id         bigserial primary key not null,
    student_id text                  NOT NULL REFERENCES users (uuid),
    number     text
);

alter table CardInfo
    owner to postgres;

