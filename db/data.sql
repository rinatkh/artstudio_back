create table public.users
(
    uuid       text primary key not null default gen_random_uuid(),
    firstname  text             not null,
    surname    text             not null,
    middlename text,
    sex        character(1)     not null,
    birth_date int              not null,
    role       text             not null,
    image      text             not null,
    created_at int              not null
);

alter table users
    owner to postgres;

create table public.auth
(
    uuid          text             NOT NULL REFERENCES users (uuid),
    email         text primary key not NULL,
    password_hash text             NOT NULL,
    password_salt text             not null
);

alter table auth
    owner to postgres;

create table public.cabinets
(
    id   bigserial primary key not null,
    name text
);

alter table cabinets
    owner to postgres;

create table public.subjects
(
    id          bigserial primary key not null,
    name        text                  NOT NULL,
    cabinet_id  bigint                NOT NULL REFERENCES cabinets (id),
    description text,
    teacher_id  text                  NOT NULL REFERENCES users (uuid),
    image       text                  not null
);

alter table subjects
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

create table public.DepartmentSubjects
(
    department_id bigint NOT NULL REFERENCES departments (id),
    subject_id    bigint NOT NULL REFERENCES subjects (id),
    PRIMARY KEY (department_id, subject_id)
);

alter table DepartmentSubjects
    owner to postgres;

create table public.schedules
(
    id           bigserial not null,
    subject_id   bigint    NOT NULL REFERENCES subjects (id),
    student_id   text      NOT NULL REFERENCES users (uuid),
    cabinet_id   bigint    NOT NULL REFERENCES cabinets (id),
    PRIMARY KEY (id, subject_id, student_id),
    when_time    int       not null,
    is_paid      boolean   not null,
    is_finished  boolean   not null,
    is_confirmed boolean   not null,
    description  text
);

alter table schedules
    owner to postgres;

create table public.statistic_students
(
    id                   bigserial primary key not null,
    student_id           text                  NOT NULL REFERENCES users (uuid),
    balance_lessons      int                   not null default 0,
    done_lessons         int                   not null default 0,
    need_payment_lessons int                   not null default 0
);

alter table statistic_students
    owner to postgres;

create table public.statistic_teachers
(
    id              bigserial primary key not null,
    teacher_id      text                  NOT NULL REFERENCES users (uuid),
    future_lessons  int                   not null default 0,
    need_dz_lessons int                   not null default 0,
    past_lessons    int                   not null default 0
);

alter table statistic_teachers
    owner to postgres;

create table public.time_teachers
(
    id          bigserial primary key not null,
    teacher_id  text                  NOT NULL REFERENCES users (uuid),
    start_time  int                   not null,
    finish_time int                   not null
);

alter table time_teachers
    owner to postgres;


create table public.cabinetTimes
(
    id          bigserial primary key not null,
    subject_id  bigint                NOT NULL REFERENCES subjects (id),
    cabinet_id  bigint                NOT NULL REFERENCES cabinets (id),
    start_time  int                   not null,
    finish_time int                   not null
);

alter table cabinetTimes
    owner to postgres;

create table public.clients
(
    id               bigserial primary key not null,
    student_id       text                  NOT NULL REFERENCES users (uuid),
    nicknameTelegram text,
    number           text,
    status           int                   not null,
    update_time      int                   not null
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

