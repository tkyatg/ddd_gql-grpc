create type gender as enum ('女性', '男性', 'その他');
create table if not exists users.users (
  user_uuid uuid not null default gen_random_uuid(),
  name varchar not null,
  email varchar not null,
  password varchar not null,
  telephone_number varchar not null,
  type gender not null,
  created_at timestamp not null DEFAULT current_timestamp,
  updated_at timestamp not null DEFAULT current_timestamp,
  constraint users_pkey primary key (user_uuid)
);
