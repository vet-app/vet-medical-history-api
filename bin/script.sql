\c vet;

CREATE TABLE "users"
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

CREATE TABLE "species"
(
    id      VARCHAR(36) NOT NULL,
    name    VARCHAR(60) NOT NULL,
    deleted Boolean     NOT NUll,

    CONSTRAINT id_pk_spc PRIMARY KEY (id)
);

CREATE TABLE "breeds"
(
    id        VARCHAR(36) NOT NULL,
    name      VARCHAR(60) NOT NULL,
    specie_id VARCHAR(36) NOT NULL,
    deleted   Boolean     NOT NUll,

    CONSTRAINT id_pk_brd PRIMARY KEY (id)
);

CREATE TABLE "pets"
(
    id          SERIAL NOT NULL,
    name        VARCHAR(60)  NOT NULL,
    born_date   DATE         NOT NULL,
    weight      VARCHAR(10),
    user_id     VARCHAR(36)  NOT NULL,
    breed_id    VARCHAR(36)  NOT NULL,
    picture_url VARCHAR      NOT NULL,
    mime_type   VARCHAR(255) NOT NULL,
    created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT id_pk_pet PRIMARY KEY (id)
);

INSERT INTO "users" (id, name, address, phone_number, email)
VALUES ('BGkus5m0kIeqDUFnCC7NCSPfNzC2',
        'John Smith',
        'Cra 163 No. 34 - 67',
        '408-237-2345',
        'john.smith@example.com');