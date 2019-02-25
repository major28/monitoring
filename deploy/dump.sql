# ************************************************************
# Sequel Pro SQL dump
# Версия 4541
#
# http://www.sequelpro.com/
# https://github.com/sequelpro/sequelpro
#
# Адрес: 127.0.0.1 (MySQL 5.5.62)
# Схема: monitoring
# Время создания: 2019-02-24 20:56:54 +0000
# ************************************************************


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

# Дамп таблицы services
# ------------------------------------------------------------

DROP TABLE IF EXISTS `services`;

CREATE TABLE `services` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL DEFAULT '',
  `host` varchar(255) NOT NULL DEFAULT '',
  `is_available` tinyint(1) DEFAULT NULL,
  `response_time` int(11) unsigned DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `host` (`host`),
  KEY `rt` (`response_time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

LOCK TABLES `services` WRITE;
/*!40000 ALTER TABLE `services` DISABLE KEYS */;

INSERT INTO `services` (`id`, `name`, `host`, `is_available`, `response_time`, `created_at`, `updated_at`, `deleted_at`)
VALUES
	(1,'google.com','https://google.com',0,280184169,'2019-02-23 19:05:35','2019-02-24 19:55:18',NULL),
	(52,'youtube.com','https://youtube.com',1,236963992,'2019-02-23 20:36:42','2019-02-24 19:54:38',NULL),
	(53,'facebook.com','https://facebook.com',0,504277493,'2019-02-23 20:36:42','2019-02-24 19:54:39',NULL),
	(54,'baidu.com','https://baidu.com',0,503824302,'2019-02-23 20:36:42','2019-02-24 19:54:40',NULL),
	(55,'wikipedia.org','https://wikipedia.org',1,397540818,'2019-02-23 20:36:42','2019-02-24 19:54:40',NULL),
	(56,'qq.com','https://qq.com',0,505122717,'2019-02-23 20:36:42','2019-02-24 19:54:40',NULL),
	(57,'taobao.com','https://taobao.com',0,505440024,'2019-02-23 20:36:42','2019-02-24 19:54:41',NULL),
	(58,'yahoo.com','https://yahoo.com',0,504239962,'2019-02-23 20:36:42','2019-02-24 19:54:41',NULL),
	(59,'tmall.com','https://tmall.com',0,505482886,'2019-02-23 20:36:42','2019-02-24 19:54:41',NULL),
	(60,'amazon.com','https://amazon.com',0,503671079,'2019-02-23 20:36:42','2019-02-24 19:54:42',NULL),
	(61,'google.co.in','https://google.co.in',1,198517778,'2019-02-23 20:36:42','2019-02-24 19:54:42',NULL),
	(62,'twitter.com','https://twitter.com',1,357742975,'2019-02-23 20:36:42','2019-02-24 19:54:43',NULL),
	(63,'sohu.com','https://sohu.com',0,502616452,'2019-02-23 20:36:42','2019-02-24 19:54:43',NULL),
	(64,'jd.com','https://jd.com',0,505495963,'2019-02-23 20:36:42','2019-02-24 19:54:43',NULL),
	(65,'live.com','https://live.com',1,286640671,'2019-02-23 20:36:42','2019-02-24 19:54:44',NULL),
	(66,'instagram.com','https://instagram.com',0,501601938,'2019-02-23 20:36:42','2019-02-24 19:54:44',NULL),
	(67,'sina.com.cn','https://sina.com.cn',0,503775099,'2019-02-23 20:36:42','2019-02-24 19:54:45',NULL),
	(68,'weibo.com','https://weibo.com',0,502638995,'2019-02-23 20:36:42','2019-02-24 19:54:45',NULL),
	(69,'google.co.jp','https://google.co.jp',1,188266635,'2019-02-23 20:36:42','2019-02-24 19:54:45',NULL),
	(70,'reddit.com','https://reddit.com',0,399621043,'2019-02-23 20:36:42','2019-02-24 19:54:46',NULL),
	(71,'vk.com','https://vk.com',1,152441419,'2019-02-23 20:36:42','2019-02-24 19:54:46',NULL),
	(72,'360.cn','https://360.cn',0,501722059,'2019-02-23 20:36:42','2019-02-24 19:54:47',NULL),
	(73,'login.tmall.com','https://login.tmall.com',0,503108611,'2019-02-23 20:36:42','2019-02-24 19:54:47',NULL),
	(74,'blogspot.com','https://blogspot.com',0,503232486,'2019-02-23 20:36:42','2019-02-24 19:54:47',NULL),
	(75,'yandex.ru','https://yandex.ru',1,227589002,'2019-02-23 20:36:42','2019-02-24 19:54:48',NULL),
	(76,'google.com.hk','https://google.com.hk',1,202334583,'2019-02-23 20:36:42','2019-02-24 19:54:48',NULL),
	(77,'netflix.com','https://netflix.com',0,504023474,'2019-02-23 20:36:42','2019-02-24 19:54:49',NULL),
	(78,'linkedin.com','https://linkedin.com',1,174861572,'2019-02-23 20:36:42','2019-02-24 19:54:49',NULL),
	(79,'pornhub.com','https://pornhub.com',1,242388884,'2019-02-23 20:36:42','2019-02-24 19:54:49',NULL),
	(80,'google.com.br','https://google.com.br',1,214807899,'2019-02-23 20:36:42','2019-02-24 19:54:50',NULL),
	(81,'twitch.tv','https://twitch.tv',1,260644529,'2019-02-23 20:36:42','2019-02-24 19:54:50',NULL),
	(82,'pages.tmall.com','https://pages.tmall.com',0,501801780,'2019-02-23 20:36:42','2019-02-24 19:54:51',NULL),
	(83,'csdn.net','https://csdn.net',0,504041456,'2019-02-23 20:36:42','2019-02-24 19:54:51',NULL),
	(84,'yahoo.co.jp','https://yahoo.co.jp',0,500540624,'2019-02-23 20:36:42','2019-02-24 19:54:51',NULL),
	(85,'mail.ru','https://mail.ru',1,88525766,'2019-02-23 20:36:42','2019-02-24 19:54:51',NULL),
	(86,'aliexpress.com','https://aliexpress.com',0,502884674,'2019-02-23 20:36:42','2019-02-24 19:54:52',NULL),
	(87,'alipay.com','https://alipay.com',0,505214537,'2019-02-23 20:36:42','2019-02-24 19:54:53',NULL),
	(88,'office.com','https://office.com',0,504428180,'2019-02-23 20:36:42','2019-02-24 19:54:53',NULL),
	(89,'google.fr','https://google.fr',1,215656001,'2019-02-23 20:36:42','2019-02-24 19:54:53',NULL),
	(90,'google.ru','https://google.ru',0,509713384,'2019-02-23 20:36:42','2019-02-24 19:54:54',NULL),
	(91,'google.co.uk','https://google.co.uk',1,218654098,'2019-02-23 20:36:42','2019-02-24 19:54:54',NULL),
	(92,'microsoftonline.com','https://microsoftonline.com',0,504466127,'2019-02-23 20:36:42','2019-02-24 19:54:55',NULL),
	(93,'google.de','https://google.de',1,194092058,'2019-02-23 20:36:42','2019-02-24 19:54:55',NULL),
	(94,'ebay.com','https://ebay.com',0,501981052,'2019-02-23 20:36:42','2019-02-24 19:54:55',NULL),
	(95,'microsoft.com','https://microsoft.com',0,505445174,'2019-02-23 20:36:42','2019-02-24 19:54:56',NULL),
	(96,'livejasmin.com','https://livejasmin.com',0,166651039,'2019-02-23 20:36:42','2019-02-24 19:54:56',NULL),
	(97,'t.co','https://t.co',0,612156,'2019-02-23 20:36:42','2019-02-24 19:54:56',NULL),
	(98,'bing.com','https://bing.com',1,419862990,'2019-02-23 20:36:42','2019-02-24 19:54:37',NULL),
	(99,'xvideos.com','https://xvideos.com',0,504062472,'2019-02-23 20:36:42','2019-02-24 19:54:37',NULL),
	(100,'google.ca','https://google.ca',1,340536197,'2019-02-23 20:36:42','2019-02-24 19:54:38',NULL);

/*!40000 ALTER TABLE `services` ENABLE KEYS */;
UNLOCK TABLES;


# Дамп таблицы users
# ------------------------------------------------------------

DROP TABLE IF EXISTS `users`;

CREATE TABLE `users` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(255) NOT NULL DEFAULT '',
  `password` varchar(255) NOT NULL DEFAULT '',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `username` (`username`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;

INSERT INTO `users` (`id`, `username`, `password`, `created_at`, `updated_at`, `deleted_at`)
VALUES
	(2,'test','test','2019-02-24 19:18:00',NULL,NULL);

/*!40000 ALTER TABLE `users` ENABLE KEYS */;
UNLOCK TABLES;


# Дамп таблицы users_history_requests
# ------------------------------------------------------------

DROP TABLE IF EXISTS `users_history_requests`;

CREATE TABLE `users_history_requests` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `fk_user` int(11) unsigned NOT NULL,
  `request` text NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;




/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
