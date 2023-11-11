# onboarding-adaptation-service
Service for hackaton i.moscov &amp; Proscom 


Серивс работает на стеке технологий: React, NodeJS, Golang, Docker.
Для решения были использованы также UE5, Blender, Figma, Adobe Premier Pro, Photoshop.
Проект построен по микросервесной архитектуре. 


На сайте имеет три страницы (Главная, Онбординг и Адаптация), каждая отвечает за свои функции.
Также имеются функции по типу создания различных объектов, просмотров их и выполнения.


Имя - пароль для входа в тестовые аккаунты:
accountant@test.com - accountant
sysadmin@test.com - sysadmin
employer@test.com - employer


Примеры вызова функций бэкэнда:
{
    "action": "reg",
    "reg": {
        "email": "employer@test.com",
        "first_name": "employer",
        "last_name": "test",
        "password": "employer"
    }
}

{
    "action": "get_all_user"
}

{
    "action": "get_user_by_email",
    "email": {
        "email": "sysadmin@test.com"
    }
}

{
    "action": "auth_user",
    "auth": {
        "email": "sysadmin@test.com",
        "password": "sysadmin"
    }
}

{
    "action": "get_all_knowledge",
    "id": {
        "id": 1
    }
}

{
    "action": "get_percent_knowledge",
    "id": {
        "id": 1
    }
}

{
    "action": "add_knowledge",
    "known": {
        "title": "22te4",
        "description": "2tes3t"
    }
}

{
    "action": "add_users_knowledge",
    "users_known": {
        "user_id": 2,
        "knowledge_id": 2
    }
}

{
    "action": "get_all_instructions"
}

{
    "action": "add_instruction",
    "instruction": {
        "title": "111testtit1le",
        "description": "222testdescr2iption"
    }
}

{
    "action": "add_users_instruction",
    "users_instructions": {
        "user_id": 1,
        "instruction_id": 3
    }
}

{
    "action": "solve_instruction",
    "users_instructions": {
        "user_id": 3,
        "instruction_id": 2
    }
}

{
    "action": "get_users_instructions",
    "id": {
        "id": 2
    }
}

{
    "action": "get_instruction",
    "id": {
        "id": 2
    }
}

{
    "action": "get_percent_instructions",
    "id": {
        "id": 3
    }
}


SQL скрипт для создания бд с пользователями:
-- public.instructions definition

-- Drop table

-- DROP TABLE public.instructions;

CREATE TABLE public.instructions (
	id serial4 NOT NULL,
	title varchar NULL,
	description text NULL,
	path varchar NULL,
	created_at timestamp NULL,
	updated_at timestamp NULL,
	CONSTRAINT instructions_pkey PRIMARY KEY (id)
);

-- public.knowledges definition

-- Drop table

-- DROP TABLE public.knowledges;

CREATE TABLE public.knowledges (
	id serial4 NOT NULL,
	title varchar NOT NULL,
	description text NULL,
	created_at timestamp NULL,
	updated_at timestamp NULL,
	CONSTRAINT knowledges_pkey PRIMARY KEY (id)
);

-- public.users definition

-- Drop table

-- DROP TABLE public.users;

CREATE TABLE public.users (
	id serial4 NOT NULL,
	email varchar NOT NULL,
	first_name varchar NULL,
	last_name varchar NULL,
	"password" varchar NOT NULL,
	profession varchar NULL,
	created_at timestamp NULL,
	updated_at timestamp NULL,
	CONSTRAINT users_pkey PRIMARY KEY (id)
);

-- public.users_instructions definition

-- Drop table

-- DROP TABLE public.users_instructions;

CREATE TABLE public.users_instructions (
	user_id int4 NOT NULL,
	instruction_id int4 NOT NULL,
	solved_at timestamp NULL,
	CONSTRAINT users_instructions_pk PRIMARY KEY (user_id, instruction_id),
	CONSTRAINT users_instructions_instruction_id_fkey FOREIGN KEY (instruction_id) REFERENCES public.instructions(id),
	CONSTRAINT users_instructions_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id)
);

-- public.users_knowledges definition

-- Drop table

-- DROP TABLE public.users_knowledges;

CREATE TABLE public.users_knowledges (
	user_id int4 NOT NULL,
	knowledge_id int4 NOT NULL,
	solved_at timestamp NULL,
	CONSTRAINT users_knowledges_pk PRIMARY KEY (user_id, knowledge_id),
	CONSTRAINT users_knowledges_knowledge_id_fkey FOREIGN KEY (knowledge_id) REFERENCES public.knowledges(id),
	CONSTRAINT users_knowledges_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id)
);

INSERT INTO public.users
(email,first_name,last_name,"password",profession,created_at,updated_at) VALUES
('accountant@test.com','accountant','test','$2a$12$KG4QeMLI7SjVcRVcn1tz9.YuZ.qjlO35oMWgO2Eqkv9O5/Uz6.CUG','accountant','2023-11-06 15:10:33.928984','2023-11-06 15:10:33.928985'),
('sysadmin@test.com','sysadmin','test','$2a$12$EI6VAOqythxyy2lkhOvrLuWjXkpvrdJqWzyxC9oopSnThjnFczd6q','sysadmin','2023-11-06 15:13:33.691082','2023-11-06 15:13:33.691082'),
('employer@test.com','employer','test','$2a$12$3aRSQ9tALryMnC4fG3rGkuJ1urKUZt09QTX5SbcyaABMgd9R1kWqW','employer','2023-11-06 15:15:40.920456','2023-11-06 15:15:40.920456');
