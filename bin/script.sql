\c vet;

CREATE TABLE "user"
(
    id           VARCHAR(36)  NOT NULL,
    name         VARCHAR(60)  NOT NULL,
    address      VARCHAR(100) NOT NULL,
    phone_number VARCHAR(30)  NOT NULL,
    email        VARCHAR(100) NOT NULL,
    created_at   TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at   TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT id_pk_usr PRIMARY KEY (id)
);

INSERT INTO "user" (id, name, address, phone_number, email)
VALUES ('BGkus5m0kIeqDUFnCC7NCSPfNzC2',
        'John Smith',
        'Cra 163 No. 34 - 67',
        '408-237-2345',
        'john.smith@example.com');