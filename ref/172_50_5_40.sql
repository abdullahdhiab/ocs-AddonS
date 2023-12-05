-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Host: 172.50.5.40
-- Generation Time: Dec 01, 2023 at 06:29 AM
-- Server version: 8.0.31-0ubuntu0.20.04.2
-- PHP Version: 8.2.8

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `addons`
--
CREATE DATABASE IF NOT EXISTS `addons` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci;
USE `addons`;

-- --------------------------------------------------------

--
-- Table structure for table `blackList`
--

CREATE TABLE `blackList` (
  `id` int NOT NULL,
  `msisdn` int NOT NULL,
  `Notes` text NOT NULL,
  `TS` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `blackList`
--

INSERT INTO `blackList` (`id`, `msisdn`, `Notes`, `TS`) VALUES
(88, 792010, '', '2023-11-28 15:39:05'),
(89, 792015, '', '2023-11-28 15:40:10'),
(90, 791010, '', '2023-11-28 15:47:48');

-- --------------------------------------------------------

--
-- Table structure for table `custProfile`
--

CREATE TABLE `custProfile` (
  `id` int UNSIGNED NOT NULL,
  `msisdn` bigint NOT NULL,
  `serviceID` int NOT NULL,
  `TS` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- --------------------------------------------------------

--
-- Table structure for table `services`
--

CREATE TABLE `services` (
  `id` int NOT NULL,
  `serviceName` varchar(256) NOT NULL,
  `serviceCode` varchar(256) NOT NULL,
  `TS` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `services`
--

INSERT INTO `services` (`id`, `serviceName`, `serviceCode`, `TS`) VALUES
(1000, 'CUG02', '2021', '2023-11-28 07:29:43');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `blackList`
--
ALTER TABLE `blackList`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `custProfile`
--
ALTER TABLE `custProfile`
  ADD PRIMARY KEY (`id`),
  ADD KEY `serviceID` (`serviceID`);

--
-- Indexes for table `services`
--
ALTER TABLE `services`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `srvCode_uiq` (`serviceCode`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `blackList`
--
ALTER TABLE `blackList`
  MODIFY `id` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=91;

--
-- AUTO_INCREMENT for table `custProfile`
--
ALTER TABLE `custProfile`
  MODIFY `id` int UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `services`
--
ALTER TABLE `services`
  MODIFY `id` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=1001;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `custProfile`
--
ALTER TABLE `custProfile`
  ADD CONSTRAINT `srcID_FK` FOREIGN KEY (`serviceID`) REFERENCES `services` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
