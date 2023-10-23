SELECT 'up SQL query';
CREATE TABLE persons
(
    id         BIGSERIAL PRIMARY KEY,
    name       VARCHAR(255) NOT NULL,
    surname    VARCHAR(255),
    patronymic VARCHAR(255),
    age        INT,
    gender     VARCHAR(255)
);
CREATE TABLE countries
(
    id          BIGSERIAL REFERENCES persons (id) ON DELETE CASCADE,
    country_id  VARCHAR(255),
    probability FLOAT
);
SELECT 'down SQL query';
