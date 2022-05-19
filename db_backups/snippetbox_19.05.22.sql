-- phpMyAdmin SQL Dump
-- version 5.1.0
-- https://www.phpmyadmin.net/
--
-- Хост: localhost
-- Время создания: Май 19 2022 г., 14:44
-- Версия сервера: 5.7.26
-- Версия PHP: 7.4.2

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- База данных: `snippetbox`
--

-- --------------------------------------------------------

--
-- Структура таблицы `snippets`
--

CREATE TABLE `snippets` (
  `id` int(11) NOT NULL,
  `title` varchar(100) NOT NULL,
  `content` text NOT NULL,
  `created` datetime NOT NULL,
  `expires` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Дамп данных таблицы `snippets`
--

INSERT INTO `snippets` (`id`, `title`, `content`, `created`, `expires`) VALUES
(1, 'An old silent pond', 'An old silent pond...\nA frog jumps into the pond,\nsplash! Silence again.\n\n– Matsuo Bashō', '2022-04-11 13:39:10', '2023-04-11 13:39:10'),
(2, 'Over the wintry forest', 'Over the wintry\nforest, winds howl in rage\nwith no leaves to blow.\n\n– Natsume Soseki', '2022-04-11 13:39:10', '2023-04-11 13:39:10'),
(3, 'First autumn morning', 'First autumn morning\nthe mirror I stare into\nshows my father\'s face.\n\n– Murakami Kijo', '2022-04-11 13:39:10', '2022-05-28 13:39:10'),
(4, 'O snail', 'O snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n– Kobayashi Issa', '2022-04-11 14:35:17', '2022-06-18 14:35:17'),
(5, 'O Gomer Micolas', 'To be or not to be.....\r\n\r\nGreat question', '2022-04-11 14:35:17', '2022-05-18 10:35:17');

--
-- Индексы сохранённых таблиц
--

--
-- Индексы таблицы `snippets`
--
ALTER TABLE `snippets`
  ADD PRIMARY KEY (`id`),
  ADD KEY `idx_snippets_created` (`created`);

--
-- AUTO_INCREMENT для сохранённых таблиц
--

--
-- AUTO_INCREMENT для таблицы `snippets`
--
ALTER TABLE `snippets`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=6;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
