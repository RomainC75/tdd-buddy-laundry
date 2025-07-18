CREATE TABLE reservation (
    id   uuid PRIMARY KEY,
    reservation_date TIMESTAMP NOT NULL,
    reservation_time INT NOT NULL, 
    email VARCHAR NOT NULL,
    pin VARCHAR NOT NULL,
    machine_num VARCHAR NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);


