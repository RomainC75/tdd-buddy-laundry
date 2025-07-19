DELETE FROM reservation;

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
    '550e8400-e29b-41d4-a716-446655440000',  
    '2025-06-23 14:00:00',                  
    54,                                     
    'bob1@email.com',                        
    'abc1',                                 
    'A10',                                  
    NOW(),                                  
    NOW()                                   
);