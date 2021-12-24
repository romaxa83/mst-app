CREATE TABLE users
(
    id                       bigserial not null unique primary key,
    name                     varchar(255) not null,
    email                    varchar(255) not null unique,
    phone                    varchar(255) unique,
    password                 varchar(255) not null,
    status                   int default '0' not null,
    refresh_token            varchar(255) unique,
    refresh_token_expires_at timestamp(0),
    deleted_at               timestamp(0),
    created_at               timestamp(0),
    updated_at               timestamp(0)
);

CREATE TABLE todo_lists
(
    id          serial       not null unique,
    title       varchar(255) not null,
    description varchar(255),
    deleted_at  timestamp(0),
    created_at  timestamp(0),
    updated_at  timestamp(0)
);

CREATE TABLE users_lists
(
    user_id int references users (id) on delete cascade      not null,
    list_id int references todo_lists (id) on delete cascade not null,
    primary key (user_id, list_id)
);

CREATE TABLE todo_items
(
    id          serial       not null unique,
    title       varchar(255) not null,
    description varchar(255),
    done        boolean      not null default false,
    deleted_at  timestamp(0),
    created_at  timestamp(0),
    updated_at  timestamp(0)
);


CREATE TABLE todo_lists_items
(
    item_id int references todo_items (id) on delete cascade not null,
    list_id int references todo_lists (id) on delete cascade not null,
    primary key (item_id, list_id)
);

CREATE TABLE media
(
    id                bigserial not null unique primary key,
    type              varchar(255) not null,
    content_type      varchar(255) not null unique,
    name              varchar(255) unique,
    size              varchar(255) not null,
    status            int default '0' not null,
    url               varchar(255) unique,
    upload_started_at timestamp(0)
);