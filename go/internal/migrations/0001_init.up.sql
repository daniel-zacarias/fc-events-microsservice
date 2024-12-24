CREATE TABLE clients (
    id VARCHAR(255),
    name VARCHAR(255),
    email VARCHAR(255),
    created_at DATE
);

CREATE TABLE account (
    id VARCHAR(255),
    client_id VARCHAR(255),
    balance INT,
    created_at DATE
);

CREATE TABLE transactions (
    id VARCHAR(255),
    account_id_from VARCHAR(255),
    account_id_to VARCHAR(255),
    amount INT,
    created_at DATE
);
