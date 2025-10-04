INSERT INTO todos (title, description, priority, is_completed, category_id, created_at, updated_at)
VALUES
('Belajar Go', 'Selesaikan modul Go dasar dan latihan CRUD', 'high', false, (SELECT id FROM categories ORDER BY RANDOM() LIMIT 1), NOW(), NOW()),
('Belanja Mingguan', 'Beli sayur, buah, dan kebutuhan rumah tangga', 'medium', false, (SELECT id FROM categories ORDER BY RANDOM() LIMIT 1), NOW(), NOW()),
('Meeting Tim', 'Rapat mingguan membahas roadmap proyek', 'high', true, (SELECT id FROM categories ORDER BY RANDOM() LIMIT 1), NOW(), NOW()),
('Olahraga Pagi', 'Lari 5km di taman dekat rumah', 'medium', false, (SELECT id FROM categories ORDER BY RANDOM() LIMIT 1), NOW(), NOW()),
('Baca Buku', 'Baca 50 halaman buku tentang produktivitas', 'low', false, (SELECT id FROM categories ORDER BY RANDOM() LIMIT 1), NOW(), NOW()),
('Update CV', 'Perbarui CV untuk apply kerja remote', 'high', false, (SELECT id FROM categories ORDER BY RANDOM() LIMIT 1), NOW(), NOW()),
('Cek Email', 'Balas email penting yang masuk hari ini', 'medium', true, (SELECT id FROM categories ORDER BY RANDOM() LIMIT 1), NOW(), NOW()),
('Design UI', 'Membuat prototype halaman dashboard', 'high', false, (SELECT id FROM categories ORDER BY RANDOM() LIMIT 1), NOW(), NOW()),
('Belajar React', 'Latihan membuat komponen dan state management', 'high', false, (SELECT id FROM categories ORDER BY RANDOM() LIMIT 1), NOW(), NOW()),
('Bayar Tagihan', 'Bayar listrik, air, dan internet sebelum tanggal 10', 'medium', true, (SELECT id FROM categories ORDER BY RANDOM() LIMIT 1), NOW(), NOW()),
('Meditasi', 'Meditasi 15 menit untuk relaksasi', 'low', false, (SELECT id FROM categories ORDER BY RANDOM() LIMIT 1), NOW(), NOW()),
('Koding Project', 'Implementasi fitur search pada project Go', 'high', false, (SELECT id FROM categories ORDER BY RANDOM() LIMIT 1), NOW(), NOW()),
('Belajar SQL', 'Latihan query JOIN dan aggregate function', 'medium', false, (SELECT id FROM categories ORDER BY RANDOM() LIMIT 1), NOW(), NOW()),
('Persiapan Presentasi', 'Siapkan slide untuk presentasi tim', 'high', false, (SELECT id FROM categories ORDER BY RANDOM() LIMIT 1), NOW(), NOW()),
('Buat ToDo List', 'Membuat daftar tugas mingguan dengan prioritas', 'medium', false, (SELECT id FROM categories ORDER BY RANDOM() LIMIT 1), NOW(), NOW());
