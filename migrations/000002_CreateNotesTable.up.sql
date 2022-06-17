CREATE TABLE Notes(
    id uuid DEFAULT uuid_generate_v4 (),
    data varchar,
    user_id uuid,
    PRIMARY KEY(id),
    CONSTRAINT fk_user FOREIGN KEY(user_id) REFERENCES Users(id)
);
