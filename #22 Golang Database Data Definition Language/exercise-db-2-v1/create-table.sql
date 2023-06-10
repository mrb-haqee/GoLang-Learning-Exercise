-- TODO: answer here
CREATE TABLE persons(
    id INTEGER PRIMARY KEY,
    NIK varchar(255) UNIQUE NOT NULL,
    fullname varchar(255) NOT NULL,
    gender varchar(50) NOT NULL,
    birth_date date NOT NULL,
    is_married BOOLEAN,
    height FLOAT,
    weight FLOAT,
    address TEXT
)


