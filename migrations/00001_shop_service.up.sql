CREATE TABLE merch_catalog
(
    id    SERIAL PRIMARY KEY,
    name  VARCHAR(50) NOT NULL UNIQUE,
    price INT         NOT NULL
);

INSERT INTO merch_catalog (name, price)
VALUES ('t-shirt', 80),
       ('cup', 20),
       ('book', 50),
       ('pen', 10),
       ('powerbank', 200),
       ('hoody', 300),
       ('umbrella', 200),
       ('socks', 10),
       ('wallet', 50),
       ('pink-hoody', 500)
ON CONFLICT (name) DO NOTHING;

CREATE TABLE users
(
    id         SERIAL PRIMARY KEY,
    username   VARCHAR(255) NOT NULL UNIQUE,
    password   VARCHAR(255) NOT NULL,
    created_at TIMESTAMP    NOT NULL DEFAULT now()
);

CREATE TABLE wallet
(
    id             SERIAL PRIMARY KEY,
    user_id        INT NOT NULL UNIQUE,
    coins_quantity INT NOT NULL DEFAULT 1000,

    FOREIGN KEY (user_id) REFERENCES users (id),
    CONSTRAINT wallet check (coins_quantity >= 0)
);

CREATE TABLE user_item
(
    id         SERIAL PRIMARY KEY,
    user_id    INT NOT NULL,
    merch_id   INT NOT NULL,
    ordered_at TIMESTAMP DEFAULT now(),
    quantity   INT,

    FOREIGN KEY (user_id) REFERENCES users (id),
    FOREIGN KEY (merch_id) REFERENCES merch_catalog (id),

    CONSTRAINT user_item check (quantity > 0)
);

CREATE TABLE coin_exchange
(
    id                SERIAL PRIMARY KEY,
    recipient_user_id INT REFERENCES users (id),
    sender_user_id    INT REFERENCES users (id),
    amount            INT       NOT NULL CHECK (amount > 0),
    exchange_at       TIMESTAMP NOT NULL DEFAULT now()
);
