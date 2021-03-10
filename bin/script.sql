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
    id          SERIAL       NOT NULL,
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

CREATE TABLE "events"
(
    id          VARCHAR(36) NOT NULL,
    title       VARCHAR(60) NOT NULL,
    description VARCHAR,
    start_date  DATE        NOT NULL,
    next_date   DATE,
    pet_id      INT         NOT NULL,
    vetstore_id VARCHAR(36),
    record_id   VARCHAR(36) NOT NULL,
    created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT id_pk_eve PRIMARY KEY (id)
);

CREATE TABLE "record_types"
(
    id      VARCHAR(36) NOT NULL,
    name    VARCHAR(60) NOT NULL,
    tag     VARCHAR(60),
    deleted Boolean     NOT NUll,

    CONSTRAINT id_pk_rct PRIMARY KEY (id)
);

CREATE TABLE "records"
(
    id             VARCHAR(36) NOT NULL,
    title          VARCHAR(60) NOT NULL,
    specie_id      VARCHAR(36) NOT NULL,
    record_type_id VARCHAR(36) NOT NULL,
    created_at     TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at     TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT id_pk_rec PRIMARY KEY (id)
);

ALTER TABLE "breeds"
    ADD CONSTRAINT brd_fk_spc FOREIGN KEY (specie_id) REFERENCES "species" (id);

ALTER TABLE "pets"
    ADD CONSTRAINT pet_fk_usr FOREIGN KEY (user_id) REFERENCES "users" (id),
    ADD CONSTRAINT pet_fk_brd FOREIGN KEY (breed_id) REFERENCES "breeds" (id);

ALTER TABLE "events"
    ADD CONSTRAINT eve_fk_pet FOREIGN KEY (pet_id) REFERENCES "pets" (id),
    ADD CONSTRAINT eve_fk_rec FOREIGN KEY (record_id) REFERENCES "records" (id);

ALTER TABLE "records"
    ADD CONSTRAINT rec_fk_spe FOREIGN KEY (specie_id) REFERENCES "species" (id),
    ADD CONSTRAINT rec_fk_rct FOREIGN KEY (record_type_id) REFERENCES "record_types" (id);

INSERT INTO "users" (id, name, address, phone_number, email)
VALUES ('BGkus5m0kIeqDUFnCC7NCSPfNzC2',
        'John Smith',
        'Cra 163 No. 34 - 67',
        '408-237-2345',
        'john.smith@example.com');