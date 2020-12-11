-- MySQL dump 10.13  Distrib 8.0.22, for Win64 (x86_64)
--
-- Host: localhost    Database: redrock_homework6_web_users_information
-- ------------------------------------------------------
-- Server version	8.0.22

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
-- Table structure for table `authority_relation`
--

DROP TABLE IF EXISTS `authority_relation`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `authority_relation` (
  `id` int NOT NULL,
  `group_id` int DEFAULT NULL,
  `leader_id` int DEFAULT NULL,
  `member_id` int DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `authority_relation`
--

LOCK TABLES `authority_relation` WRITE;
/*!40000 ALTER TABLE `authority_relation` DISABLE KEYS */;
INSERT INTO `authority_relation` VALUES (1,1,3,3),(2,1,3,1);
/*!40000 ALTER TABLE `authority_relation` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `message_bin`
--

DROP TABLE IF EXISTS `message_bin`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `message_bin` (
  `id` int NOT NULL,
  `previous_floor` int DEFAULT NULL,
  `time` varchar(255) DEFAULT NULL,
  `user` varchar(255) DEFAULT NULL,
  `information` varchar(255) DEFAULT NULL,
  `outside_name` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `message_bin`
--

LOCK TABLES `message_bin` WRITE;
/*!40000 ALTER TABLE `message_bin` DISABLE KEYS */;
/*!40000 ALTER TABLE `message_bin` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `message_favorite`
--

DROP TABLE IF EXISTS `message_favorite`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `message_favorite` (
  `id` int NOT NULL AUTO_INCREMENT,
  `floor` int DEFAULT NULL,
  `name` varchar(255) DEFAULT 'Unknown',
  `if_like` int DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `message_favorite`
--

LOCK TABLES `message_favorite` WRITE;
/*!40000 ALTER TABLE `message_favorite` DISABLE KEYS */;
INSERT INTO `message_favorite` VALUES (1,1,'tester',1),(2,1,'hong24',1),(3,2,'Unknown',1),(4,2,'hong24',0),(5,3,'tester',1);
/*!40000 ALTER TABLE `message_favorite` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `message_history`
--

DROP TABLE IF EXISTS `message_history`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `message_history` (
  `id` int NOT NULL AUTO_INCREMENT,
  `previous_floor` int DEFAULT '0',
  `time` varchar(255) DEFAULT NULL,
  `user` varchar(255) DEFAULT 'Unknown',
  `information` varchar(255) DEFAULT NULL,
  `outside_name` varchar(255) DEFAULT 'Unknown',
  `authority_id` int DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `message_history`
--

LOCK TABLES `message_history` WRITE;
/*!40000 ALTER TABLE `message_history` DISABLE KEYS */;
INSERT INTO `message_history` VALUES (1,0,'2020-11-03 21:42:03','Administrator','First floor is mine!','Administrator',0),(2,1,'2020-11-08 21:28:23','hong24','oh~','hong24',0),(3,1,'2020-11-22 10:50:51','hong24','emmmmmmm= .=','hong24',0),(4,0,'2020-11-26 05:50:59','qrg47','QAQ','qrg47',0),(5,4,'2020-12-01 11:51:13','qrg47','prprprprprprpr','Unknown',1),(6,5,'2020-12-03 21:50:21','Administrator','GOD...','Administrator',1),(7,0,'2020-12-03 23:08:59','Administrator','hello','Unknown',0);
/*!40000 ALTER TABLE `message_history` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `users_information`
--

DROP TABLE IF EXISTS `users_information`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `users_information` (
  `ID` int NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  `password` varchar(255) DEFAULT NULL,
  `signature` varchar(255) DEFAULT '这个人很懒，什么都没有留下',
  `telephone_number` bigint DEFAULT NULL,
  `MD5salt` bigint DEFAULT NULL,
  PRIMARY KEY (`ID`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `users_information`
--

LOCK TABLES `users_information` WRITE;
/*!40000 ALTER TABLE `users_information` DISABLE KEYS */;
INSERT INTO `users_information` VALUES (1,'Administrator','efbfbd421aa90e079fa326b6494f812ad13e79','这是我的地盘',20202020202,1607065568),(2,'hong24','efbfbd8b45316877beae7f0d4ddff900ba9083','丢人，马上给我滚出战场',12359310537,1607063860),(3,'qrg47','efbfbdcdd821b0cbe52170781fe3fb79d5be15','利兹与青鸟什么时候引进啊啊啊啊',15723593342,1607063503),(4,'tester','efbfbd25d55ad283aa400af464c76d713c07ad','这个人很懒，什么都没有留下',12453469775,1607612157);
/*!40000 ALTER TABLE `users_information` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2020-12-11 16:59:42
