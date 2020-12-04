-- MySQL dump 10.13  Distrib 8.0.22, for Win64 (x86_64)
--
-- Host: localhost    Database: homework6_student_information_correction
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
-- Table structure for table `courses`
--

DROP TABLE IF EXISTS `courses`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `courses` (
  `ID` int NOT NULL AUTO_INCREMENT,
  `details` varchar(255) NOT NULL,
  PRIMARY KEY (`ID`),
  UNIQUE KEY `ID` (`ID`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `courses`
--

LOCK TABLES `courses` WRITE;
/*!40000 ALTER TABLE `courses` DISABLE KEYS */;
INSERT INTO `courses` VALUES (1,'art'),(2,'sports'),(3,'math');
/*!40000 ALTER TABLE `courses` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `specialty`
--

DROP TABLE IF EXISTS `specialty`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `specialty` (
  `ID` int NOT NULL AUTO_INCREMENT,
  `details` varchar(255) NOT NULL,
  PRIMARY KEY (`ID`),
  UNIQUE KEY `ID` (`ID`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `specialty`
--

LOCK TABLES `specialty` WRITE;
/*!40000 ALTER TABLE `specialty` DISABLE KEYS */;
INSERT INTO `specialty` VALUES (1,'math'),(2,'art');
/*!40000 ALTER TABLE `specialty` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `students_information`
--

DROP TABLE IF EXISTS `students_information`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `students_information` (
  `students_ID` int NOT NULL,
  `classNumber` int DEFAULT NULL,
  `name` varchar(20) DEFAULT 'name',
  `password` varchar(255) NOT NULL,
  `average_point` int DEFAULT NULL,
  PRIMARY KEY (`students_ID`),
  UNIQUE KEY `students_ID` (`students_ID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `students_information`
--

LOCK TABLES `students_information` WRITE;
/*!40000 ALTER TABLE `students_information` DISABLE KEYS */;
INSERT INTO `students_information` VALUES (2020233333,27,'王小美','Wangxiaomei12138',2);
/*!40000 ALTER TABLE `students_information` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `students_points_relation`
--

DROP TABLE IF EXISTS `students_points_relation`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `students_points_relation` (
  `ID` int NOT NULL AUTO_INCREMENT,
  `students_ID` int NOT NULL,
  `course_ID` int NOT NULL,
  `points` float DEFAULT NULL,
  PRIMARY KEY (`ID`),
  UNIQUE KEY `ID` (`ID`),
  KEY `students_ID` (`students_ID`),
  KEY `course_ID` (`course_ID`),
  CONSTRAINT `students_points_relation_ibfk_1` FOREIGN KEY (`students_ID`) REFERENCES `students_information` (`students_ID`),
  CONSTRAINT `students_points_relation_ibfk_2` FOREIGN KEY (`course_ID`) REFERENCES `courses` (`ID`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `students_points_relation`
--

LOCK TABLES `students_points_relation` WRITE;
/*!40000 ALTER TABLE `students_points_relation` DISABLE KEYS */;
INSERT INTO `students_points_relation` VALUES (1,2020233333,1,3.6),(2,2020233333,2,1),(3,2020233333,3,3.2);
/*!40000 ALTER TABLE `students_points_relation` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `students_specialty_relation`
--

DROP TABLE IF EXISTS `students_specialty_relation`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `students_specialty_relation` (
  `ID` int NOT NULL AUTO_INCREMENT,
  `students_ID` int NOT NULL,
  `specialty_aspect_ID` int NOT NULL,
  PRIMARY KEY (`ID`),
  UNIQUE KEY `ID` (`ID`),
  KEY `students_ID` (`students_ID`),
  KEY `specialty_aspect_ID` (`specialty_aspect_ID`),
  CONSTRAINT `students_specialty_relation_ibfk_1` FOREIGN KEY (`students_ID`) REFERENCES `students_information` (`students_ID`),
  CONSTRAINT `students_specialty_relation_ibfk_2` FOREIGN KEY (`specialty_aspect_ID`) REFERENCES `specialty` (`ID`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `students_specialty_relation`
--

LOCK TABLES `students_specialty_relation` WRITE;
/*!40000 ALTER TABLE `students_specialty_relation` DISABLE KEYS */;
INSERT INTO `students_specialty_relation` VALUES (1,2020233333,1),(2,2020233333,2);
/*!40000 ALTER TABLE `students_specialty_relation` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2020-12-03 16:22:33
