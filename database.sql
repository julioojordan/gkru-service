CREATE TABLE users
(
    id INT NOT NULL AUTO_INCREMENT,
    username VARCHAR(30) NOT NULL,
    password TEXT,
    PRIMARY KEY (id)
)ENGINE InnoDB;

CREATE TABLE peserta
(
    id INT NOT NULL AUTO_INCREMENT,
    username VARCHAR(30) NOT NULL,
    PRIMARY KEY (id)
)ENGINE InnoDB;

drop table IF EXISTS data_keluarga, data_anggota, lingkungan, wilayah, keluarga_anggota_rel;


CREATE TABLE wilayah (
    id INT PRIMARY KEY NOT NULL AUTO_INCREMENT,
    kode_wilayah VARCHAR(255),
    nama_wilayah VARCHAR(255)
);

CREATE TABLE lingkungan (
    id INT PRIMARY KEY NOT NULL AUTO_INCREMENT,
    kode_lingkungan VARCHAR(255),
    nama_lingkungan VARCHAR(255),
    id_wilayah INT,
    FOREIGN KEY (id_wilayah) REFERENCES wilayah(id)
);


CREATE TABLE data_anggota (
    id INT PRIMARY KEY NOT NULL AUTO_INCREMENT,
    nama_lengkap VARCHAR(255),
    tanggal_lahir DATE,
    tanggal_baptis DATE,
    keterangan VARCHAR(255)
);

CREATE TABLE keluarga_anggota_rel (
    id INT NOT NULL AUTO_INCREMENT,
    id_keluarga INT,
    id_anggota INT,
    hubungan VARCHAR(255),
    PRIMARY KEY (id),
    FOREIGN KEY (id_anggota) REFERENCES data_anggota(id),
    UNIQUE (id_keluarga, id_anggota)
);


CREATE TABLE data_keluarga (
    id INT NOT NULL AUTO_INCREMENT,
    id_wilayah INT,
    id_lingkungan INT,
    nomor INT,
    id_kepala_keluarga INT,
    id_keluarga_anggota_rel INT,
    alamat VARCHAR(255),
    FOREIGN KEY (id_wilayah) REFERENCES wilayah(id),
    FOREIGN KEY (id_lingkungan) REFERENCES lingkungan(id),
    FOREIGN KEY (id_kepala_keluarga) REFERENCES data_anggota(id),
    PRIMARY KEY (id)
);


-- dummy data
INSERT INTO wilayah (kode_wilayah, nama_wilayah) VALUES
('W1', 'Wilayah 1'),
('W2', 'Wilayah 2');


INSERT INTO lingkungan (kode_lingkungan, nama_lingkungan, id_wilayah) VALUES
('Lingkungan A', 'Lingkungan A', 1),
('Lingkungan B', 'Lingkungan B', 1),
('Lingkungan C', 'Lingkungan C', 2);

-- Insert data untuk setiap anggota keluarga (total 36 anggota)
INSERT INTO data_anggota (nama_lengkap, tanggal_lahir, tanggal_baptis, keterangan)
VALUES
-- Keluarga 1
('Kepala Keluarga 1', '1980-01-01', '1985-01-01', 'Keterangan Kepala Keluarga 1'),
('Istri Keluarga 1', '1982-01-01', '1987-01-01', 'Keterangan Istri Keluarga 1'),
('Anak 1 Keluarga 1', '2000-01-01', '2005-01-01', 'Keterangan Anak 1 Keluarga 1'),
('Anak 2 Keluarga 1', '2002-01-01', '2007-01-01', 'Keterangan Anak 2 Keluarga 1'),

-- Keluarga 2
('Kepala Keluarga 2', '1975-01-01', '1980-01-01', 'Keterangan Kepala Keluarga 2'),
('Istri Keluarga 2', '1978-01-01', '1983-01-01', 'Keterangan Istri Keluarga 2'),
('Anak 1 Keluarga 2', '1996-01-01', '2001-01-01', 'Keterangan Anak 1 Keluarga 2'),
('Anak 2 Keluarga 2', '1998-01-01', '2003-01-01', 'Keterangan Anak 2 Keluarga 2'),

-- Keluarga 3
('Kepala Keluarga 3', '1983-01-01', '1988-01-01', 'Keterangan Kepala Keluarga 3'),
('Istri Keluarga 3', '1985-01-01', '1990-01-01', 'Keterangan Istri Keluarga 3'),
('Anak 1 Keluarga 3', '2005-01-01', '2010-01-01', 'Keterangan Anak 1 Keluarga 3'),
('Anak 2 Keluarga 3', '2008-01-01', '2013-01-01', 'Keterangan Anak 2 Keluarga 3'),

-- dan seterusnya untuk 36 anggota
('Kepala Keluarga 4', '1970-01-01', '1975-01-01', 'Keterangan Kepala Keluarga 4'),
('Istri Keluarga 4', '1973-01-01', '1978-01-01', 'Keterangan Istri Keluarga 4'),
('Anak 1 Keluarga 4', '1992-01-01', '1997-01-01', 'Keterangan Anak 1 Keluarga 4'),
('Anak 2 Keluarga 4', '1994-01-01', '1999-01-01', 'Keterangan Anak 2 Keluarga 4'),

-- dan seterusnya untuk 36 anggota
('Kepala Keluarga 5', '1988-01-01', '1993-01-01', 'Keterangan Kepala Keluarga 5'),
('Istri Keluarga 5', '1990-01-01', '1995-01-01', 'Keterangan Istri Keluarga 5'),
('Anak 1 Keluarga 5', '2010-01-01', '2015-01-01', 'Keterangan Anak 1 Keluarga 5'),
('Anak 2 Keluarga 5', '2012-01-01', '2017-01-01', 'Keterangan Anak 2 Keluarga 5'),

-- dan seterusnya untuk 36 anggota
('Kepala Keluarga 6', '1972-01-01', '1977-01-01', 'Keterangan Kepala Keluarga 6'),
('Istri Keluarga 6', '1974-01-01', '1979-01-01', 'Keterangan Istri Keluarga 6'),
('Anak 1 Keluarga 6', '1994-01-01', '1999-01-01', 'Keterangan Anak 1 Keluarga 6'),
('Anak 2 Keluarga 6', '1996-01-01', '2001-01-01', 'Keterangan Anak 2 Keluarga 6'),

-- dan seterusnya untuk 36 anggota
('Kepala Keluarga 7', '1985-01-01', '1990-01-01', 'Keterangan Kepala Keluarga 7'),
('Istri Keluarga 7', '1988-01-01', '1993-01-01', 'Keterangan Istri Keluarga 7'),
('Anak 1 Keluarga 7', '2008-01-01', '2013-01-01', 'Keterangan Anak 1 Keluarga 7'),
('Anak 2 Keluarga 7', '2010-01-01', '2015-01-01', 'Keterangan Anak 2 Keluarga 7'),

-- dan seterusnya untuk 36 anggota
('Kepala Keluarga 8', '1977-01-01', '1982-01-01', 'Keterangan Kepala Keluarga 8'),
('Istri Keluarga 8', '1980-01-01', '1985-01-01', 'Keterangan Istri Keluarga 8'),
('Anak 1 Keluarga 8', '2000-01-01', '2005-01-01', 'Keterangan Anak 1 Keluarga 8'),
('Anak 2 Keluarga 8', '2002-01-01', '2007-01-01', 'Keterangan Anak 2 Keluarga 8'),

-- dan seterusnya untuk 36 anggota
('Kepala Keluarga 9', '1979-01-01', '1984-01-01', 'Keterangan Kepala Keluarga 9'),
('Istri Keluarga 9', '1982-01-01', '1987-01-01', 'Keterangan Istri Keluarga 9'),
('Anak 1 Keluarga 9', '2002-01-01', '2007-01-01', 'Keterangan Anak 1 Keluarga 9'),
('Anak 2 Keluarga 9', '2004-01-01', '2009-01-01', 'Keterangan Anak 2 Keluarga 9');

-- Insert relasi anggota keluarga untuk setiap keluarga
INSERT INTO keluarga_anggota_rel (id_keluarga, id_anggota, hubungan) VALUES
-- Keluarga 1
(1, 1, 'Kepala Keluarga'),
(1, 2, 'Istri'),
(1, 3, 'Anak'),
(1, 4, 'Anak'),
-- Keluarga 2
(2, 5, 'Kepala Keluarga'),
(2, 6, 'Istri'),
(2, 7, 'Anak'),
(2, 8, 'Anak'),
-- Keluarga 3
(3, 9, 'Kepala Keluarga'),
(3, 10, 'Istri'),
(3, 11, 'Anak'),
(3, 12, 'Anak'),
-- Keluarga 4
(4, 13, 'Kepala Keluarga'),
(4, 14, 'Istri'),
(4, 15, 'Anak'),
(4, 16, 'Anak'),
-- Keluarga 5
(5, 17, 'Kepala Keluarga'),
(5, 18, 'Istri'),
(5, 19, 'Anak'),
(5, 20, 'Anak'),
-- Keluarga 6
(6, 21, 'Kepala Keluarga'),
(6, 22, 'Istri'),
(6, 23, 'Anak'),
(6, 24, 'Anak'),
-- Keluarga 7
(7, 25, 'Kepala Keluarga'),
(7, 26, 'Istri'),
(7, 27, 'Anak'),
(7, 28, 'Anak'),
-- Keluarga 8
(8, 29, 'Kepala Keluarga'),
(8, 30, 'Istri'),
(8, 31, 'Anak'),
(8, 32, 'Anak'),
-- Keluarga 9
(9, 33, 'Kepala Keluarga'),
(9, 34, 'Istri'),
(9, 35, 'Anak'),
(9, 36, 'Anak');


-- Insert data untuk setiap keluarga (total 9 keluarga)
INSERT INTO data_keluarga (id_wilayah, id_lingkungan, nomor, id_kepala_keluarga, id_keluarga_anggota_rel, alamat) VALUES
-- Keluarga di Wilayah 1 - Lingkungan A
(1, 1, 1, 1, 1, 'Alamat 1A'), -- id_kepala_keluarga = 1, id_keluarga_anggota_rel = 1
(1, 1, 2, 5, 5, 'Alamat 2A'), -- id_kepala_keluarga = 5, id_keluarga_anggota_rel = 5
(1, 1, 3, 9, 9, 'Alamat 3A'), -- id_kepala_keluarga = 9, id_keluarga_anggota_rel = 9
-- Keluarga di Wilayah 1 - Lingkungan B
(1, 2, 4, 13, 13, 'Alamat 1B'), -- id_kepala_keluarga = 13, id_keluarga_anggota_rel = 13
(1, 2, 5, 17, 17, 'Alamat 2B'), -- id_kepala_keluarga = 17, id_keluarga_anggota_rel = 17
(1, 2, 6, 21, 21, 'Alamat 3B'), -- id_kepala_keluarga = 21, id_keluarga_anggota_rel = 21
-- Keluarga di Wilayah 2 - Lingkungan C
(2, 3, 7, 25, 25, 'Alamat 1C'), -- id_kepala_keluarga = 25, id_keluarga_anggota_rel = 25
(2, 3, 8, 29, 29, 'Alamat 2C'), -- id_kepala_keluarga = 29, id_keluarga_anggota_rel = 29
(2, 3, 9, 33, 33, 'Alamat 3C'); -- id_kepala_keluarga = 33, id_keluarga_anggota_rel = 33

CREATE TABLE wealth (
    id INT AUTO_INCREMENT PRIMARY KEY,
    total FLOAT DEFAULT 0,
    lingkungan INT NOT NULL,
    wilayah INT NOT NULL
);

INSERT INTO wealth (total, lingkungan, wilayah) VALUES
(50000, 1, 1),
(70000, 2, 1),
(100000, 3, 2);

CREATE TABLE riwayat_transaksi (
    id INT AUTO_INCREMENT PRIMARY KEY,
    nominal FLOAT DEFAULT 0,
    sumber_dana INT NOT NULL, -- dari wealth id mana
    sumber_pemasukan INT NOT NULL, -- dari id peserta id yang memasukan dana
    keterangan VARCHAR(255) NOT NULL
);

INSERT INTO riwayat_transaksi (nominal, sumber_dana, sumber_pemasukan, keterangan) VALUES
(50000, 1, 1, 'IN'),
(70000, 2, 13, 'IN'),
(100000, 3, 25, 'IN');



note
coba buatkan data dummy untuk masing-masing table etersebut sehingga saya dapat mengujinya
dengan jumlah data
1. data keluarga ada 8 keluarga yang terdiri dari 
2 data anggota ada total 8x4 yaitu 32 data
3 data wilayah ada 3 wilayah masing masing memiliki
