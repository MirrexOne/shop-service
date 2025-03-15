CREATE TYPE merch_type AS ENUM (
    't-shirt',
    'cup',
    'book',
    'pen',
    'powerbank',
    'hoody',
    'umbrella',
    'socks',
    'wallet',
    'pink-hoody'
);

CREATE TABLE merch
(
    id    SERIAL PRIMARY KEY,
    name  MERCH_TYPE,
    price INT
);
