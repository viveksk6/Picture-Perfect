-- phpMyAdmin SQL Dump
-- version 4.7.4
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Nov 18, 2019 at 02:01 PM
-- Server version: 10.1.26-MariaDB
-- PHP Version: 7.1.9

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `pictureperfect`
--

-- --------------------------------------------------------

--
-- Table structure for table `cinemas`
--

CREATE TABLE `cinemas` (
  `cineplexId` int(5) NOT NULL,
  `cineplexName` varchar(50) NOT NULL,
  `city` varchar(50) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- --------------------------------------------------------

--
-- Table structure for table `movies`
--

CREATE TABLE `movies` (
  `movieId` int(10) NOT NULL,
  `title` varchar(50) NOT NULL,
  `summary` varchar(256) NOT NULL,
  `genre` varchar(10) NOT NULL,
  `img` varchar(256) NOT NULL,
  `language` varchar(10) NOT NULL,
  `certificate` varchar(5) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `movies`
--

INSERT INTO `movies` (`movieId`, `title`, `summary`, `genre`, `img`, `language`, `certificate`) VALUES
(1, 'War', 'An Indian soldier is assigned to eliminate his former mentor and he must keep his wits about him if he is to be successful in his mission. When the two men collide, it results in a barrage of battles and bullets.', 'Action', 'https://pictureperfectvivek.s3.amazonaws.com/movie2.jpg', 'Hindi', 'U/A'),
(2, 'Bahubali 2 - The Conclusion', 'When Bhallaladeva conspires against his brother to become the king of Mahishmati, he has him killed by Katappa and imprisons his wife. Years later, his brother\'s son returns to avenge his father\'s death.', 'Drama', 'https://pictureperfectvivek.s3.amazonaws.com/movie1.jpg', 'Telugu', 'U/A'),
(3, 'Joker', 'Forever alone in a crowd, failed comedian Arthur Fleck seeks connection as he walks the streets of Gotham City. ', 'Thriller', 'https://pictureperfectvivek.s3.amazonaws.com/movie3.jpg', 'English', 'A'),
(4, 'Kaithi', ' A prisoner on parole, who\'s on his way to meet his daughter, is enlisted by an injured police officer to protect a few cops.', 'Thriller', 'https://pictureperfectvivek.s3.amazonaws.com/movie4.jpg', 'Tamil', 'U/A');

-- --------------------------------------------------------

--
-- Table structure for table `review`
--

CREATE TABLE `review` (
  `movieId` int(5) NOT NULL,
  `userId` int(5) NOT NULL,
  `review` varchar(256) NOT NULL,
  `rating` float NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- --------------------------------------------------------

--
-- Table structure for table `screening`
--

CREATE TABLE `screening` (
  `movieId` int(5) NOT NULL,
  `cineplexId` int(5) NOT NULL,
  `startTime` varchar(10) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE `users` (
  `userId` int(5) NOT NULL,
  `name` varchar(50) NOT NULL,
  `password` varchar(20) NOT NULL,
  `role` varchar(20) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Indexes for dumped tables
--

--
-- Indexes for table `cinemas`
--
ALTER TABLE `cinemas`
  ADD PRIMARY KEY (`cineplexId`);

--
-- Indexes for table `movies`
--
ALTER TABLE `movies`
  ADD PRIMARY KEY (`movieId`),
  ADD KEY `movieId` (`movieId`);

--
-- Indexes for table `review`
--
ALTER TABLE `review`
  ADD PRIMARY KEY (`movieId`,`userId`),
  ADD KEY `userId` (`userId`),
  ADD KEY `movieId` (`movieId`);

--
-- Indexes for table `screening`
--
ALTER TABLE `screening`
  ADD PRIMARY KEY (`movieId`,`cineplexId`,`startTime`),
  ADD KEY `cineplexId` (`cineplexId`);

--
-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`userId`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `cinemas`
--
ALTER TABLE `cinemas`
  MODIFY `cineplexId` int(5) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `movies`
--
ALTER TABLE `movies`
  MODIFY `movieId` int(10) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `users`
--
ALTER TABLE `users`
  MODIFY `userId` int(5) NOT NULL AUTO_INCREMENT;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `review`
--
ALTER TABLE `review`
  ADD CONSTRAINT `review_ibfk_1` FOREIGN KEY (`movieId`) REFERENCES `movies` (`movieid`),
  ADD CONSTRAINT `review_ibfk_2` FOREIGN KEY (`userId`) REFERENCES `users` (`userId`);

--
-- Constraints for table `screening`
--
ALTER TABLE `screening`
  ADD CONSTRAINT `screening_ibfk_1` FOREIGN KEY (`movieId`) REFERENCES `movies` (`movieid`),
  ADD CONSTRAINT `screening_ibfk_2` FOREIGN KEY (`cineplexId`) REFERENCES `cinemas` (`cineplexId`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
