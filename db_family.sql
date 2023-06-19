-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Host: localhost
-- Generation Time: Jun 19, 2023 at 04:08 PM
-- Server version: 10.4.28-MariaDB
-- PHP Version: 8.1.17

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `db_family`
--

-- --------------------------------------------------------

--
-- Table structure for table `customer`
--

CREATE TABLE `customer` (
  `cst_id` int(11) NOT NULL,
  `nationality_id` int(11) NOT NULL,
  `cst_dob` date NOT NULL,
  `cst_phone_num` varchar(20) NOT NULL,
  `cst_email` varchar(50) NOT NULL,
  `cst_name` char(50) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `customer`
--

INSERT INTO `customer` (`cst_id`, `nationality_id`, `cst_dob`, `cst_phone_num`, `cst_email`, `cst_name`, `created_at`, `updated_at`) VALUES
(1, 3, '1997-05-23', '0895367298151', 'ompusungguandry@gmail.com', 'Andry Halomoan Ompusunggu', '2023-06-17 12:24:27', '2023-06-17 23:06:23');

-- --------------------------------------------------------

--
-- Table structure for table `family_list`
--

CREATE TABLE `family_list` (
  `fl_id` int(11) NOT NULL,
  `cst_id` int(11) NOT NULL,
  `fl_relation` varchar(50) NOT NULL,
  `fl_dob` date NOT NULL,
  `fl_name` varchar(50) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `family_list`
--

INSERT INTO `family_list` (`fl_id`, `cst_id`, `fl_relation`, `fl_dob`, `fl_name`, `created_at`, `updated_at`) VALUES
(6, 1, 'Bapak', '1988-01-14', 'Bornok Ompusunggu', '2023-06-17 09:25:38', '2023-06-17 16:25:38');

-- --------------------------------------------------------

--
-- Table structure for table `nationality`
--

CREATE TABLE `nationality` (
  `nationality_id` int(11) NOT NULL,
  `nationality_name` varchar(50) NOT NULL,
  `nationality_code` char(2) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `nationality`
--

INSERT INTO `nationality` (`nationality_id`, `nationality_name`, `nationality_code`) VALUES
(1, 'Indonesia', 'ID'),
(2, 'Japan', 'JP'),
(3, 'United States of America', 'US'),
(4, 'Australia', 'AU'),
(5, 'South Africa', 'SA');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `customer`
--
ALTER TABLE `customer`
  ADD PRIMARY KEY (`cst_id`),
  ADD KEY `nationality_id` (`nationality_id`);

--
-- Indexes for table `family_list`
--
ALTER TABLE `family_list`
  ADD PRIMARY KEY (`fl_id`),
  ADD KEY `cst_id` (`cst_id`);

--
-- Indexes for table `nationality`
--
ALTER TABLE `nationality`
  ADD PRIMARY KEY (`nationality_id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `customer`
--
ALTER TABLE `customer`
  MODIFY `cst_id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=9;

--
-- AUTO_INCREMENT for table `family_list`
--
ALTER TABLE `family_list`
  MODIFY `fl_id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=7;

--
-- AUTO_INCREMENT for table `nationality`
--
ALTER TABLE `nationality`
  MODIFY `nationality_id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=6;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `customer`
--
ALTER TABLE `customer`
  ADD CONSTRAINT `customer_ibfk_1` FOREIGN KEY (`nationality_id`) REFERENCES `nationality` (`nationality_id`);

--
-- Constraints for table `family_list`
--
ALTER TABLE `family_list`
  ADD CONSTRAINT `family_list_ibfk_1` FOREIGN KEY (`cst_id`) REFERENCES `customer` (`cst_id`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
