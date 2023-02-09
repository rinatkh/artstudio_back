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
    uuid          text             NOT NULL,
    email         text primary key not NULL,
    password_hash text             NOT NULL,
    password_salt text             not null,
    foreign key (uuid) REFERENCES users (uuid)
        on delete cascade
        on update no action
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
    cabinet_id  bigint                NOT NULL,
    description text,
    teacher_id  text                  NOT NULL,
    image       text                  not null,
    foreign key (cabinet_id) REFERENCES cabinets (id)
        on delete cascade
        on update no action,
    foreign key (teacher_id) REFERENCES users (uuid)
        on delete cascade
        on update no action
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
    department_id bigint NOT NULL,
    subject_id    bigint NOT NULL,
    PRIMARY KEY (department_id, subject_id),
    foreign key (department_id) REFERENCES departments (id)
        on delete cascade
        on update no action,
    foreign key (subject_id) REFERENCES subjects (id)
        on delete cascade
        on update no action
);

alter table DepartmentSubjects
    owner to postgres;

create table public.schedules
(
    id           bigserial not null,
    subject_id   bigint    NOT NULL,
    student_id   text      NOT NULL,
    cabinet_id   bigint    NOT NULL,
    PRIMARY KEY (id, subject_id, student_id),
    when_time    int       not null,
    is_paid      boolean   not null,
    is_finished  boolean   not null,
    is_confirmed boolean   not null,
    description  text,
    foreign key (cabinet_id) REFERENCES cabinets (id)
        on delete cascade
        on update no action,
    foreign key (student_id) REFERENCES users (uuid)
        on delete cascade
        on update no action,
    foreign key (subject_id) REFERENCES subjects (id)
        on delete cascade
        on update no action
);

alter table schedules
    owner to postgres;

create table public.statistic_students
(
    id                   bigserial primary key not null,
    student_id           text                  NOT NULL,
    balance_lessons      int                   not null default 0,
    done_lessons         int                   not null default 0,
    need_payment_lessons int                   not null default 0,
    foreign key (student_id) REFERENCES users (uuid)
        on delete cascade
        on update no action
);

alter table statistic_students
    owner to postgres;

create table public.statistic_teachers
(
    id              bigserial primary key not null,
    teacher_id      text                  NOT NULL,
    future_lessons  int                   not null default 0,
    need_dz_lessons int                   not null default 0,
    past_lessons    int                   not null default 0,
    foreign key (teacher_id) REFERENCES users (uuid)
        on delete cascade
        on update no action
);

alter table statistic_teachers
    owner to postgres;

-- CREATE OR REPLACE FUNCTION create_statistic_student() RETURNS TRIGGER AS
-- $$
-- BEGIN
--     IF NOT EXISTS(Select student_id
--                   from statistic_students
--                   where student_id = (SELECT uuid FROM Users WHERE NEW.role = 'STUDENT' AND uuid = NEW.uuid)) THEN
--         INSERT INTO statistic_students (student_id) VALUES (NEW.uuid);
--     END IF;
--     IF NOT EXISTS(Select teacher_id
--                   from statistic_teachers
--                   where teacher_id = (SELECT uuid FROM Users WHERE NEW.role = 'TEACHER' AND uuid = NEW.uuid)) THEN
--         INSERT INTO statistic_teachers (teacher_id) VALUES (NEW.uuid);
--     END IF;
-- END;
-- $$ LANGUAGE plpgsql;
--
-- CREATE TRIGGER create_statistic_student AFTER INSERT ON public.users FOR EACH ROW EXECUTE PROCEDURE create_statistic_student();

create table public.time_teachers
(
    id          bigserial primary key not null,
    teacher_id  text                  NOT NULL,
    start_time  int                   not null,
    finish_time int                   not null,
    foreign key (teacher_id) REFERENCES users (uuid)
        on delete cascade
        on update no action
);

alter table time_teachers
    owner to postgres;


create table public.cabinetTimes
(
    id          bigserial primary key not null,
    subject_id  bigint                NOT NULL,
    cabinet_id  bigint                NOT NULL,
    start_time  int                   not null,
    finish_time int                   not null,
    foreign key (cabinet_id) REFERENCES cabinets (id)
        on delete cascade
        on update no action,
    foreign key (subject_id) REFERENCES subjects (id)
        on delete cascade
        on update no action
);

alter table cabinetTimes
    owner to postgres;

create table public.clients
(
    id               bigserial primary key not null,
    student_id       text                  NOT NULL,
    nicknameTelegram text,
    number           text,
    status           int                   not null,
    update_time      int                   not null,
    foreign key (student_id) REFERENCES users (uuid)
        on delete cascade
        on update no action
);

alter table clients
    owner to postgres;

create table public.CardInfo
(
    id         bigserial primary key not null,
    student_id text                  NOT NULL,
    number     text,
    foreign key (student_id) REFERENCES users (uuid)
        on delete cascade
        on update no action
);

alter table CardInfo
    owner to postgres;

