INSERT INTO clients(id, name, email, created_at)
VALUES 
    ('e48dc30b-e087-460c-87b3-d825156d4659', 'John Doe', 'j@j.com', now()),
    ('900528df-7de8-457c-a350-451fa1523629', 'Johana Doe', 'j2@j.com', now());

INSERT INTO account(id, client_id, balance, created_at)
VALUES 
    ('46444f3b-c687-4db8-b08a-d8ad56fb2a3c', 'e48dc30b-e087-460c-87b3-d825156d4659', 1000, now()),
    ('e01a8e9c-f875-4e04-a737-490cfa6ebc4a', '900528df-7de8-457c-a350-451fa1523629', 1000, now());