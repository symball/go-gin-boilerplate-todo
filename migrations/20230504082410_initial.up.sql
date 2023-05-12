SET statement_timeout = 0;

--bun:split

CREATE TABLE users (
   id serial PRIMARY KEY,
   username VARCHAR ( 255 ) NOT NULL,
   password VARCHAR ( 255 ),
   roles TEXT[],
   status TEXT
);

--bun:split

CREATE TABLE todos (
      id serial PRIMARY KEY,
      title VARCHAR ( 255 ) NOT NULL,
      summary VARCHAR ( 255 ),
      content TEXT,
      complete BOOLEAN DEFAULT FALSE
);
