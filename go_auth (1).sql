-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Aug 08, 2023 at 10:45 PM
-- Server version: 10.4.28-MariaDB
-- PHP Version: 8.2.4

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `go_auth`
--

DELIMITER $$
--
-- Procedures
--
CREATE DEFINER=`root`@`localhost` PROCEDURE `generar_horariosss` ()   BEGIN
    DECLARE hora_inicio TIME;
    SET hora_inicio = '09:00:00';

    WHILE hora_inicio <= '17:00:00' DO
        INSERT INTO horario (hora_inicio, hora_salida,id_medico)
        VALUES (hora_inicio, ADDTIME(hora_inicio, '00:30:00'),20);

        SET hora_inicio = ADDTIME(hora_inicio, '00:30:00');
    END WHILE;
END$$

DELIMITER ;

-- --------------------------------------------------------

--
-- Table structure for table `citas_creadas`
--

CREATE TABLE `citas_creadas` (
  `Id` int(11) NOT NULL,
  `id_paciente` int(20) UNSIGNED NOT NULL,
  `Descripcion` varchar(200) NOT NULL,
  `Especialidad` varchar(150) NOT NULL,
  `id_medico` int(10) NOT NULL,
  `Fecha` date NOT NULL,
  `Hora` time NOT NULL,
  `Estado` varchar(60) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `citas_creadas`
--

INSERT INTO `citas_creadas` (`Id`, `id_paciente`, `Descripcion`, `Especialidad`, `id_medico`, `Fecha`, `Hora`, `Estado`) VALUES
(6, 1, '', 'Pediatría', 15, '2023-08-17', '16:30:00', 'Pendiente'),
(7, 1, 'Dolor de cabeza', 'Pediatría', 15, '2023-09-01', '15:00:00', 'Pendiente'),
(8, 1, 'Fiebre y vomito', 'Pediatría', 15, '2023-08-17', '17:00:00', 'Pendiente'),
(9, 1, '', 'Pediatría', 15, '2023-08-17', '10:30:00', 'Pendiente'),
(10, 1, 'Tos y gripe ', 'Pediatría', 15, '2023-08-17', '10:00:00', 'Pendiente'),
(11, 1, 'No se lo que tengo', 'Pediatría', 15, '2023-08-17', '09:00:00', 'Pendiente'),
(12, 1, 'Post cirugia', 'Cardiología', 19, '2023-08-24', '12:00:00', 'Pendiente'),
(13, 1, 'Cirugia', 'Cirugía', 19, '2023-08-24', '17:00:00', 'Pendiente'),
(14, 1, 'Dolor al orinar ', 'Pediatría', 16, '2023-08-24', '13:00:00', 'Pendiente'),
(15, 1, 'No lo se', 'Urología', 16, '2023-08-24', '09:00:00', 'Pendiente');

-- --------------------------------------------------------

--
-- Table structure for table `dia`
--

CREATE TABLE `dia` (
  `ID_dia` int(20) NOT NULL,
  `Dia` date NOT NULL,
  `id_horario` int(20) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `dia`
--

INSERT INTO `dia` (`ID_dia`, `Dia`, `id_horario`) VALUES
(22, '2023-08-23', 76),
(23, '2023-08-16', 71),
(24, '2023-08-16', 71),
(25, '2023-08-09', 86),
(26, '2023-08-10', 79),
(27, '2023-08-10', 75),
(28, '2023-08-17', 86),
(29, '2023-09-01', 83),
(30, '2023-08-17', 87),
(31, '2023-08-17', 74),
(32, '2023-08-17', 73),
(33, '2023-08-17', 71),
(34, '2023-08-24', 145),
(35, '2023-08-24', 155),
(36, '2023-08-24', 113),
(38, '2023-08-24', 105);

-- --------------------------------------------------------

--
-- Table structure for table `especialidades`
--

CREATE TABLE `especialidades` (
  `Id_especialidad` int(10) NOT NULL,
  `Nombre` varchar(100) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `especialidades`
--

INSERT INTO `especialidades` (`Id_especialidad`, `Nombre`) VALUES
(8, 'Pediatría'),
(9, 'Cirugía'),
(10, 'Cardiología'),
(11, 'Gastroenterología'),
(12, 'Dermatología'),
(13, 'Neurología'),
(14, 'Urología');

-- --------------------------------------------------------

--
-- Table structure for table `horario`
--

CREATE TABLE `horario` (
  `Id` int(10) NOT NULL,
  `hora_inicio` time NOT NULL,
  `hora_salida` time NOT NULL,
  `id_medico` int(20) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `horario`
--

INSERT INTO `horario` (`Id`, `hora_inicio`, `hora_salida`, `id_medico`) VALUES
(71, '09:00:00', '09:30:00', 15),
(72, '09:30:00', '10:00:00', 15),
(73, '10:00:00', '10:30:00', 15),
(74, '10:30:00', '11:00:00', 15),
(75, '11:00:00', '11:30:00', 15),
(76, '11:30:00', '12:00:00', 15),
(77, '12:00:00', '12:30:00', 15),
(78, '12:30:00', '13:00:00', 15),
(79, '13:00:00', '13:30:00', 15),
(80, '13:30:00', '14:00:00', 15),
(81, '14:00:00', '14:30:00', 15),
(82, '14:30:00', '15:00:00', 15),
(83, '15:00:00', '15:30:00', 15),
(84, '15:30:00', '16:00:00', 15),
(85, '16:00:00', '16:30:00', 15),
(86, '16:30:00', '17:00:00', 15),
(87, '17:00:00', '17:30:00', 15),
(88, '09:00:00', '09:30:00', 18),
(89, '09:30:00', '10:00:00', 18),
(90, '10:00:00', '10:30:00', 18),
(91, '10:30:00', '11:00:00', 18),
(92, '11:00:00', '11:30:00', 18),
(93, '11:30:00', '12:00:00', 18),
(94, '12:00:00', '12:30:00', 18),
(95, '12:30:00', '13:00:00', 18),
(96, '13:00:00', '13:30:00', 18),
(97, '13:30:00', '14:00:00', 18),
(98, '14:00:00', '14:30:00', 18),
(99, '14:30:00', '15:00:00', 18),
(100, '15:00:00', '15:30:00', 18),
(101, '15:30:00', '16:00:00', 18),
(102, '16:00:00', '16:30:00', 18),
(103, '16:30:00', '17:00:00', 18),
(104, '17:00:00', '17:30:00', 18),
(105, '09:00:00', '09:30:00', 16),
(106, '09:30:00', '10:00:00', 16),
(107, '10:00:00', '10:30:00', 16),
(108, '10:30:00', '11:00:00', 16),
(109, '11:00:00', '11:30:00', 16),
(110, '11:30:00', '12:00:00', 16),
(111, '12:00:00', '12:30:00', 16),
(112, '12:30:00', '13:00:00', 16),
(113, '13:00:00', '13:30:00', 16),
(114, '13:30:00', '14:00:00', 16),
(115, '14:00:00', '14:30:00', 16),
(116, '14:30:00', '15:00:00', 16),
(117, '15:00:00', '15:30:00', 16),
(118, '15:30:00', '16:00:00', 16),
(119, '16:00:00', '16:30:00', 16),
(120, '16:30:00', '17:00:00', 16),
(121, '17:00:00', '17:30:00', 16),
(122, '09:00:00', '09:30:00', 17),
(123, '09:30:00', '10:00:00', 17),
(124, '10:00:00', '10:30:00', 17),
(125, '10:30:00', '11:00:00', 17),
(126, '11:00:00', '11:30:00', 17),
(127, '11:30:00', '12:00:00', 17),
(128, '12:00:00', '12:30:00', 17),
(129, '12:30:00', '13:00:00', 17),
(130, '13:00:00', '13:30:00', 17),
(131, '13:30:00', '14:00:00', 17),
(132, '14:00:00', '14:30:00', 17),
(133, '14:30:00', '15:00:00', 17),
(134, '15:00:00', '15:30:00', 17),
(135, '15:30:00', '16:00:00', 17),
(136, '16:00:00', '16:30:00', 17),
(137, '16:30:00', '17:00:00', 17),
(138, '17:00:00', '17:30:00', 17),
(139, '09:00:00', '09:30:00', 19),
(140, '09:30:00', '10:00:00', 19),
(141, '10:00:00', '10:30:00', 19),
(142, '10:30:00', '11:00:00', 19),
(143, '11:00:00', '11:30:00', 19),
(144, '11:30:00', '12:00:00', 19),
(145, '12:00:00', '12:30:00', 19),
(146, '12:30:00', '13:00:00', 19),
(147, '13:00:00', '13:30:00', 19),
(148, '13:30:00', '14:00:00', 19),
(149, '14:00:00', '14:30:00', 19),
(150, '14:30:00', '15:00:00', 19),
(151, '15:00:00', '15:30:00', 19),
(152, '15:30:00', '16:00:00', 19),
(153, '16:00:00', '16:30:00', 19),
(154, '16:30:00', '17:00:00', 19),
(155, '17:00:00', '17:30:00', 19),
(156, '09:00:00', '09:30:00', 20),
(157, '09:30:00', '10:00:00', 20),
(158, '10:00:00', '10:30:00', 20),
(159, '10:30:00', '11:00:00', 20),
(160, '11:00:00', '11:30:00', 20),
(161, '11:30:00', '12:00:00', 20),
(162, '12:00:00', '12:30:00', 20),
(163, '12:30:00', '13:00:00', 20),
(164, '13:00:00', '13:30:00', 20),
(165, '13:30:00', '14:00:00', 20),
(166, '14:00:00', '14:30:00', 20),
(167, '14:30:00', '15:00:00', 20),
(168, '15:00:00', '15:30:00', 20),
(169, '15:30:00', '16:00:00', 20),
(170, '16:00:00', '16:30:00', 20),
(171, '16:30:00', '17:00:00', 20),
(172, '17:00:00', '17:30:00', 20);

-- --------------------------------------------------------

--
-- Table structure for table `medicos`
--

CREATE TABLE `medicos` (
  `id` int(10) NOT NULL,
  `DNI` int(30) NOT NULL,
  `Nombre` varchar(200) NOT NULL,
  `id_especialidad` int(10) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `medicos`
--

INSERT INTO `medicos` (`id`, `DNI`, `Nombre`, `id_especialidad`) VALUES
(15, 876666, 'Yubiny Sibrian', 8),
(16, 897, 'Yulissa Garcia', 14),
(17, 112000022, 'Brayan Cano miralda', 8),
(18, 122000033, 'Yubiny Romero', 8),
(19, 232411112, 'Flor Romero Vasquez', 9),
(20, 123516252, 'Miguel Salinas', 9);

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE `users` (
  `id` int(10) UNSIGNED NOT NULL,
  `full_name` varchar(200) NOT NULL,
  `email` varchar(200) NOT NULL,
  `username` varchar(80) NOT NULL,
  `user_type` varchar(50) NOT NULL,
  `password` varchar(100) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `users`
--

INSERT INTO `users` (`id`, `full_name`, `email`, `username`, `user_type`, `password`) VALUES
(1, 'jaime sibrian', 'jaime@gmail.com', 'jaime123', 'Paciente', '1234'),
(2, 'jaime', 'yyyyyyy', 'uuuuuuu', 'Doctor', '1111'),
(3, 'Nodis orellana', 'nodis@gmail.com', 'nodis12', 'Paciente', '1111'),
(4, 'Ivan Rivera', 'ivan@unah.hn', 'Ivan321', 'Doctor', '7777');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `citas_creadas`
--
ALTER TABLE `citas_creadas`
  ADD PRIMARY KEY (`Id`),
  ADD KEY `fk_pacientes` (`id_paciente`),
  ADD KEY `fk_med` (`id_medico`);

--
-- Indexes for table `dia`
--
ALTER TABLE `dia`
  ADD PRIMARY KEY (`ID_dia`),
  ADD KEY `Fk_horas` (`id_horario`);

--
-- Indexes for table `especialidades`
--
ALTER TABLE `especialidades`
  ADD PRIMARY KEY (`Id_especialidad`);

--
-- Indexes for table `horario`
--
ALTER TABLE `horario`
  ADD PRIMARY KEY (`Id`);

--
-- Indexes for table `medicos`
--
ALTER TABLE `medicos`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `citas_creadas`
--
ALTER TABLE `citas_creadas`
  MODIFY `Id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=16;

--
-- AUTO_INCREMENT for table `dia`
--
ALTER TABLE `dia`
  MODIFY `ID_dia` int(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=39;

--
-- AUTO_INCREMENT for table `especialidades`
--
ALTER TABLE `especialidades`
  MODIFY `Id_especialidad` int(10) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=15;

--
-- AUTO_INCREMENT for table `horario`
--
ALTER TABLE `horario`
  MODIFY `Id` int(10) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=173;

--
-- AUTO_INCREMENT for table `medicos`
--
ALTER TABLE `medicos`
  MODIFY `id` int(10) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=21;

--
-- AUTO_INCREMENT for table `users`
--
ALTER TABLE `users`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=5;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `citas_creadas`
--
ALTER TABLE `citas_creadas`
  ADD CONSTRAINT `fk_med` FOREIGN KEY (`id_medico`) REFERENCES `medicos` (`id`);

--
-- Constraints for table `dia`
--
ALTER TABLE `dia`
  ADD CONSTRAINT `Fk_horas` FOREIGN KEY (`id_horario`) REFERENCES `horario` (`Id`);

--
-- Constraints for table `horario`
--
ALTER TABLE `horario`
  ADD CONSTRAINT `Fk_Horarios` FOREIGN KEY (`id_medico`) REFERENCES `medicos` (`id`);

--
-- Constraints for table `medicos`
--
ALTER TABLE `medicos`
  ADD CONSTRAINT `fk_medic` FOREIGN KEY (`id_especialidad`) REFERENCES `especialidades` (`Id_especialidad`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
