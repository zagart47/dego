SELECT 'up SQL query';
CREATE TABLE IF NOT EXISTS persons
(
    id         BIGSERIAL PRIMARY KEY ,
    name       VARCHAR(255) NOT NULL,
    surname    VARCHAR(255),
    patronymic VARCHAR(255),
    age        SMALLINT,
    gender     VARCHAR(255)

);

CREATE TABLE IF NOT EXISTS countries
(
    person_id   BIGINT REFERENCES persons (id) ON DELETE CASCADE,
    country_id  VARCHAR(255) NOT NULL,
    probability FLOAT,
    PRIMARY KEY (person_id, country_id)
);
SELECT 'down SQL query';
