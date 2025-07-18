-- name: GetReservation :one
SELECT * FROM reservation WHERE id = $1;

-- name: ListReservations :many
SELECT * FROM reservation ORDER BY reservation_date;

-- name: CreateReservation :one
INSERT INTO reservation (
    id,
    reservation_date,
    reservation_time,
    email,
    pin,
    machine_num,
    created_at,
    updated_at
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8
) RETURNING *;
