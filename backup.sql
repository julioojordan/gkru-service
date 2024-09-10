-- MySQL dump 10.13  Distrib 8.0.36, for Win64 (x86_64)
--
-- Host: localhost    Database: gkru_app
-- ------------------------------------------------------
-- Server version	8.0.36

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `data_anggota`
--

DROP TABLE IF EXISTS `data_anggota`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `data_anggota` (
  `id` int NOT NULL AUTO_INCREMENT,
  `nama_lengkap` varchar(255) DEFAULT NULL,
  `tanggal_lahir` date DEFAULT NULL,
  `tanggal_baptis` date DEFAULT NULL,
  `keterangan` varchar(255) DEFAULT NULL,
  `status` varchar(100) DEFAULT 'HIDUP',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=39 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `data_anggota`
--

LOCK TABLES `data_anggota` WRITE;
/*!40000 ALTER TABLE `data_anggota` DISABLE KEYS */;
INSERT INTO `data_anggota` VALUES (1,'Kepala Keluarga 1','1980-01-01','1985-01-01','Keterangan Kepala Keluarga 1','HIDUP'),(2,'Istri Keluarga 1','1982-01-01','1987-01-01','','HIDUP'),(3,'Anak 1 Keluarga 1','2000-01-01','2005-01-01','Keterangan Anak 1 Keluarga 1','HIDUP'),(4,'Diablo','2024-09-28','2024-10-28','Anak update','HIDUP'),(5,'Kepala Keluarga 2','1975-01-01','1980-01-01','Keterangan Kepala Keluarga 2','HIDUP'),(6,'Istri Keluarga 2','1978-01-01','1983-01-01','Keterangan Istri Keluarga 2','HIDUP'),(7,'Anak 1 Keluarga 2','1996-01-01','2001-01-01','Keterangan Anak 1 Keluarga 2','HIDUP'),(8,'Anak 2 Keluarga 2','1998-01-01','2003-01-01','Keterangan Anak 2 Keluarga 2','HIDUP'),(9,'Kepala Keluarga 3','1983-01-01','1988-01-01','Keterangan Kepala Keluarga 3','HIDUP'),(10,'Istri Keluarga 3','1985-01-01','1990-01-01','Keterangan Istri Keluarga 3','HIDUP'),(11,'Anak 1 Keluarga 3','2005-01-01','2010-01-01','Keterangan Anak 1 Keluarga 3','HIDUP'),(12,'Anak 2 Keluarga 3','2008-01-01','2013-01-01','Keterangan Anak 2 Keluarga 3','HIDUP'),(13,'Kepala Keluarga 4','1970-01-01','1975-01-01','Keterangan Kepala Keluarga 4','HIDUP'),(14,'Istri Keluarga 4','1973-01-01','1978-01-01','Keterangan Istri Keluarga 4','HIDUP'),(15,'Anak 1 Keluarga 4','1992-01-01','1997-01-01','Keterangan Anak 1 Keluarga 4','HIDUP'),(16,'Anak 2 Keluarga 4','1994-01-01','1999-01-01','Keterangan Anak 2 Keluarga 4','HIDUP'),(17,'Kepala Keluarga 5','1988-01-01','1993-01-01','Keterangan Kepala Keluarga 5','HIDUP'),(18,'Istri Keluarga 5','1990-01-01','1995-01-01','Keterangan Istri Keluarga 5','HIDUP'),(19,'Anak 1 Keluarga 5','2010-01-01','2015-01-01','Keterangan Anak 1 Keluarga 5','HIDUP'),(20,'Anak 2 Keluarga 5','2012-01-01','2017-01-01','Keterangan Anak 2 Keluarga 5','HIDUP'),(21,'Kepala Keluarga 6','1972-01-01','1977-01-01','Keterangan Kepala Keluarga 6','HIDUP'),(22,'Istri Keluarga 6','1974-01-01','1979-01-01','Keterangan Istri Keluarga 6','HIDUP'),(23,'Anak 1 Keluarga 6','1994-01-01','1999-01-01','Keterangan Anak 1 Keluarga 6','HIDUP'),(24,'Anak 2 Keluarga 6','1996-01-01','2001-01-01','Keterangan Anak 2 Keluarga 6','HIDUP'),(25,'Kepala Keluarga 7','1985-01-01','1990-01-01','Keterangan Kepala Keluarga 7','HIDUP'),(26,'Istri Keluarga 7','1988-01-01','1993-01-01','Keterangan Istri Keluarga 7','HIDUP'),(27,'Anak 1 Keluarga 7','2008-01-01','2013-01-01','Keterangan Anak 1 Keluarga 7','HIDUP'),(28,'Anak 2 Keluarga 7','2010-01-01','2015-01-01','Keterangan Anak 2 Keluarga 7','HIDUP'),(29,'Kepala Keluarga 8','1977-01-01','1982-01-01','Keterangan Kepala Keluarga 8','HIDUP'),(30,'Istri Keluarga 8','1980-01-01','1985-01-01','Keterangan Istri Keluarga 8','HIDUP'),(31,'Anak 1 Keluarga 8','2000-01-01','2005-01-01','Keterangan Anak 1 Keluarga 8','HIDUP'),(32,'Anak 2 Keluarga 8','2002-01-01','2007-01-01','Keterangan Anak 2 Keluarga 8','HIDUP'),(33,'Kepala Keluarga 9','1979-01-01','1984-01-01','Keterangan Kepala Keluarga 9','HIDUP'),(34,'Istri Keluarga 9','1982-01-01','1987-01-01','Keterangan Istri Keluarga 9','HIDUP'),(35,'Anak 1 Keluarga 9','2002-01-01','2007-01-01','Keterangan Anak 1 Keluarga 9','HIDUP'),(36,'Anak 2 Keluarga 9','2004-01-01','2009-01-01','Keterangan Anak 2 Keluarga 9','HIDUP'),(37,'Milim Nava','2024-08-28','2024-09-28','Anak baru','Hidup'),(38,'Rimuru','2024-08-28','2024-08-28','Kepala Keluarga 10','HIDUP');
/*!40000 ALTER TABLE `data_anggota` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `data_keluarga`
--

DROP TABLE IF EXISTS `data_keluarga`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `data_keluarga` (
  `id` int NOT NULL AUTO_INCREMENT,
  `id_wilayah` int DEFAULT NULL,
  `id_lingkungan` int DEFAULT NULL,
  `nomor` varchar(255) DEFAULT NULL,
  `id_kepala_keluarga` int DEFAULT NULL,
  `id_keluarga_anggota_rel` int DEFAULT NULL COMMENT 'id keluarga di tabel keluarga_anggota_rel, kemungkinan ini tidak guna -> perlu di hapus  nanti',
  `alamat` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `id_wilayah` (`id_wilayah`),
  KEY `id_lingkungan` (`id_lingkungan`),
  KEY `id_kepala_keluarga` (`id_kepala_keluarga`),
  CONSTRAINT `data_keluarga_ibfk_1` FOREIGN KEY (`id_wilayah`) REFERENCES `wilayah` (`id`),
  CONSTRAINT `data_keluarga_ibfk_2` FOREIGN KEY (`id_lingkungan`) REFERENCES `lingkungan` (`id`),
  CONSTRAINT `data_keluarga_ibfk_3` FOREIGN KEY (`id_kepala_keluarga`) REFERENCES `data_anggota` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `data_keluarga`
--

LOCK TABLES `data_keluarga` WRITE;
/*!40000 ALTER TABLE `data_keluarga` DISABLE KEYS */;
INSERT INTO `data_keluarga` VALUES (1,2,3,'3322415',2,1,'tempest'),(2,1,1,'2',5,5,'Alamat 2A'),(3,1,1,'3',9,9,'Alamat 3A'),(4,1,2,'4',13,13,'Alamat 1B'),(5,1,2,'5',17,17,'Alamat 2B'),(6,1,2,'6',21,21,'Alamat 3B'),(7,2,3,'7',25,25,'Alamat 1C'),(8,2,3,'8',29,29,'Alamat 2C'),(9,2,3,'9',33,33,'Alamat 3C'),(10,2,3,'3322415',38,NULL,'tempest');
/*!40000 ALTER TABLE `data_keluarga` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `keluarga_anggota_rel`
--

DROP TABLE IF EXISTS `keluarga_anggota_rel`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `keluarga_anggota_rel` (
  `id` int NOT NULL AUTO_INCREMENT,
  `id_keluarga` int DEFAULT NULL,
  `id_anggota` int DEFAULT NULL,
  `hubungan` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id_keluarga` (`id_keluarga`,`id_anggota`),
  KEY `id_anggota` (`id_anggota`),
  CONSTRAINT `keluarga_anggota_rel_ibfk_1` FOREIGN KEY (`id_anggota`) REFERENCES `data_anggota` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=39 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `keluarga_anggota_rel`
--

LOCK TABLES `keluarga_anggota_rel` WRITE;
/*!40000 ALTER TABLE `keluarga_anggota_rel` DISABLE KEYS */;
INSERT INTO `keluarga_anggota_rel` VALUES (1,1,1,'Kepala Keluarga'),(2,1,2,'Kepala Keluarga'),(3,1,3,'Anak'),(4,1,4,'Anak'),(5,2,5,'Kepala Keluarga'),(6,2,6,'Istri'),(7,2,7,'Anak'),(8,2,8,'Anak'),(9,3,9,'Kepala Keluarga'),(10,3,10,'Istri'),(11,3,11,'Anak'),(12,3,12,'Anak'),(13,4,13,'Kepala Keluarga'),(14,4,14,'Istri'),(15,4,15,'Anak'),(16,4,16,'Anak'),(17,5,17,'Kepala Keluarga'),(18,5,18,'Istri'),(19,5,19,'Anak'),(20,5,20,'Anak'),(21,6,21,'Kepala Keluarga'),(22,6,22,'Istri'),(23,6,23,'Anak'),(24,6,24,'Anak'),(25,7,25,'Kepala Keluarga'),(26,7,26,'Istri'),(27,7,27,'Anak'),(28,7,28,'Anak'),(29,8,29,'Kepala Keluarga'),(30,8,30,'Istri'),(31,8,31,'Anak'),(32,8,32,'Anak'),(33,9,33,'Kepala Keluarga'),(34,9,34,'Istri'),(35,9,35,'Anak'),(36,9,36,'Anak'),(37,1,37,'Anak'),(38,0,38,'Kepala Keluarga');
/*!40000 ALTER TABLE `keluarga_anggota_rel` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `lingkungan`
--

DROP TABLE IF EXISTS `lingkungan`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `lingkungan` (
  `id` int NOT NULL AUTO_INCREMENT,
  `kode_lingkungan` varchar(255) DEFAULT NULL,
  `nama_lingkungan` varchar(255) DEFAULT NULL,
  `id_wilayah` int DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `id_wilayah` (`id_wilayah`),
  CONSTRAINT `lingkungan_ibfk_1` FOREIGN KEY (`id_wilayah`) REFERENCES `wilayah` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `lingkungan`
--

LOCK TABLES `lingkungan` WRITE;
/*!40000 ALTER TABLE `lingkungan` DISABLE KEYS */;
INSERT INTO `lingkungan` VALUES (1,'Lingkungan A','Lingkungan A',1),(2,'Lingkungan B','Lingkungan B',1),(3,'Lingkungan C','Lingkungan C',2);
/*!40000 ALTER TABLE `lingkungan` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `peserta`
--

DROP TABLE IF EXISTS `peserta`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `peserta` (
  `id` int NOT NULL AUTO_INCREMENT,
  `username` varchar(30) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `peserta`
--

LOCK TABLES `peserta` WRITE;
/*!40000 ALTER TABLE `peserta` DISABLE KEYS */;
INSERT INTO `peserta` VALUES (1,'coba 1'),(2,'coba 2');
/*!40000 ALTER TABLE `peserta` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `riwayat_transaksi`
--

DROP TABLE IF EXISTS `riwayat_transaksi`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `riwayat_transaksi` (
  `id` int NOT NULL AUTO_INCREMENT,
  `nominal` float DEFAULT '0',
  `wealth_id` int NOT NULL COMMENT 'lokasi tempat dana disimpan, ada di wealth mana gitu',
  `keterangan` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `created_by` int DEFAULT NULL COMMENT 'ID yang melakukan pengeluaran atau pemasukan dana misalkan admin/user non peserta',
  `tanggal` datetime DEFAULT CURRENT_TIMESTAMP,
  `id_wilayah` int DEFAULT NULL,
  `id_lingkungan` int DEFAULT NULL,
  `updated_by` int DEFAULT NULL COMMENT 'id user yang melakukan update pada rows',
  `sub_keterangan` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `riwayat_transaksi`
--

LOCK TABLES `riwayat_transaksi` WRITE;
/*!40000 ALTER TABLE `riwayat_transaksi` DISABLE KEYS */;
INSERT INTO `riwayat_transaksi` VALUES (1,50000,1,'IN',1,'2024-08-04 08:44:17',1,1,NULL,NULL),(2,70000,2,'IN',13,'2024-08-04 08:44:17',1,2,NULL,NULL),(3,100000,3,'IN',25,'2024-08-04 08:44:17',2,3,NULL,NULL);
/*!40000 ALTER TABLE `riwayat_transaksi` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `users`
--

DROP TABLE IF EXISTS `users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `users` (
  `id` int NOT NULL AUTO_INCREMENT,
  `username` varchar(30) NOT NULL,
  `password` text,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `users`
--

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` VALUES (1,'coba','1234');
/*!40000 ALTER TABLE `users` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `wealth`
--

DROP TABLE IF EXISTS `wealth`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `wealth` (
  `id` int NOT NULL AUTO_INCREMENT,
  `total` float DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `wealth`
--

LOCK TABLES `wealth` WRITE;
/*!40000 ALTER TABLE `wealth` DISABLE KEYS */;
INSERT INTO `wealth` VALUES (1,220000);
/*!40000 ALTER TABLE `wealth` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `wilayah`
--

DROP TABLE IF EXISTS `wilayah`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `wilayah` (
  `id` int NOT NULL AUTO_INCREMENT,
  `kode_wilayah` varchar(255) DEFAULT NULL,
  `nama_wilayah` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `wilayah`
--

LOCK TABLES `wilayah` WRITE;
/*!40000 ALTER TABLE `wilayah` DISABLE KEYS */;
INSERT INTO `wilayah` VALUES (1,'W1','Wilayah 1'),(2,'W2','Wilayah 2');
/*!40000 ALTER TABLE `wilayah` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2024-09-10 19:53:09
