INSERT INTO categories (name, color, created_at)
VALUES
('Personal', 'RED', NOW()),
('Work', 'BLUE', NOW()),
('Shopping', 'GREEN', NOW())
ON CONFLICT (name) DO NOTHING;
