-- Dummy Admin with username 'admin' and password 'admin'
INSERT INTO mst_admin (id, username, password) VALUES (
    '8020a4df-ede6-49be-a252-ebe8b4ed4e09', 
    'admin',
    '$2a$14$RP/3xXZRc/oZO53auBHGnOzXRBbdqtFv/xI6pqrhDRjsJRrXTDUoe',
);

-- Dummy Bank
INSERT INTO mst_bank(id, name) VALUES 
('b5b051ba-26aa-45e6-bbfc-800fbf947a01', 'BCA'),
('4279c931-de13-413e-bc31-0a03d2ed90d6', 'BRI'),
('efef51d4-1fe4-4eca-a749-6db8b2184968', 'BNI'),
('1b751130-6504-419b-b68c-54195df91ea5', 'Mandiri');

-- Dummy Merchant
INSERT INTO mst_merchant(id, name, address) VALUES 
('6ab5c8c5-1ea9-4cc8-b2b3-94492561db51', 'Good Food', '21, Jalan Teuku Umar, Menteng, Special Capital Region of Jakarta, 10350'),
('039e1080-571f-4e8c-b8e7-90c11b090b27', 'Lumbung Bakso', 'Km 1, Jalan Raya Cileungsi-Jonggol, West Java, 16820'),
('f7ebf3c8-93d8-4225-b11f-5a939798e61c', 'Sushinesia', '56, Jalan Raya Margonda, Depok, West Java, 16431'),
('a7f3a06e-b704-4487-828d-e3025435f11e', 'Ganyem', '109, Jalan Wonorejo Permai Selatan IX, Surabaya, East Java, 60296');