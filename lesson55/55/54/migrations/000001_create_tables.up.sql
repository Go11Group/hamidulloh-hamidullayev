CREATE TABLE restaurants (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR NOT NULL,
    address VARCHAR NOT NULL,
    total_avb_seats INTEGER NOT NULL,
    phone_number VARCHAR,
    description VARCHAR,
    created_at TIMESTAMP DEFAULT current_timestamp,
    updated_at TIMESTAMP DEFAULT current_timestamp,
    deleted_at TIMESTAMP
);

CREATE TABLE reservations (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    restaurant_id UUID REFERENCES restaurants(id),
    user_id UUID,
    arriving_time TIMESTAMP,
    number_of_seats INTEGER,
    created_at TIMESTAMP DEFAULT current_timestamp,
    updated_at TIMESTAMP DEFAULT current_timestamp,
    deleted_at TIMESTAMP
);

CREATE TABLE menu (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    item_type VARCHAR NOT NULL,
    name VARCHAR NOT NULL,
    price DECIMAL NOT NULL,
    description TEXT,
    restaurant_id UUID REFERENCES restaurants(id),
    created_at TIMESTAMP DEFAULT current_timestamp,
    updated_at TIMESTAMP DEFAULT current_timestamp,
    deleted_at TIMESTAMP 
);

CREATE TABLE orders (
    id SERIAL PRIMARY KEY,
    reservation_id UUID REFERENCES reservations(id),
    menu_item_id UUID REFERENCES menu(id),
    quantity INTEGER NOT NULL,
    created_at TIMESTAMP DEFAULT current_timestamp,
    updated_at TIMESTAMP DEFAULT current_timestamp,
    deleted_at TIMESTAMP
);