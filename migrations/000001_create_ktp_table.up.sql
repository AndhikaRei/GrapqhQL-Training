CREATE TABLE IF NOT EXISTS ktps (
    id INT UNIQUE AUTO_INCREMENT,
    nik VARCHAR(120),
    nama VARCHAR(120),
    agama VARCHAR(120),
    jenis_kelamin VARCHAR(120), 
    tanggal_lahir DATETIME(3),
    created_at DATETIME(3),
    updated_at DATETIME(3)
);
