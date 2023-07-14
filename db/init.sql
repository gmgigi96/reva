--
-- SQL script to create the database structure used by Reva from scratch
--
-- The script was obtained by dumping the CERNBox production database
-- and keeping the tables effectively referenced in the Reva code base

-- MySQL dump 10.14  Distrib 5.5.68-MariaDB, for Linux (x86_64)
--
/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `cbox_metadata`
--

DROP TABLE IF EXISTS `cbox_metadata`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `cbox_metadata` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `item_type` int(11) NOT NULL,
  `uid` varchar(64) NOT NULL,
  `fileid_prefix` varchar(255) NOT NULL,
  `fileid` varchar(255) NOT NULL,
  `tag_key` varchar(255) NOT NULL,
  `tag_val` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=11623 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `cbox_otg_ocis`
--

DROP TABLE IF EXISTS `cbox_otg_ocis`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `cbox_otg_ocis` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `message` text NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `cernbox_project_mapping`
--

DROP TABLE IF EXISTS `cernbox_project_mapping`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `cernbox_project_mapping` (
  `project_name` varchar(50) NOT NULL,
  `eos_relative_path` varchar(50) NOT NULL,
  `project_owner` varchar(50) NOT NULL DEFAULT 'temp',
  PRIMARY KEY (`project_name`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `notification_recipients`
--

DROP TABLE IF EXISTS `notification_recipients`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `notification_recipients` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `notification_id` int(11) NOT NULL,
  `recipient` varchar(320) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `notification_recipients_ix0` (`notification_id`),
  KEY `notification_recipients_ix1` (`recipient`),
  CONSTRAINT `notification_recipients_ibfk_1` FOREIGN KEY (`notification_id`) REFERENCES `notifications` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `notifications`
--

DROP TABLE IF EXISTS `notifications`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `notifications` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `ref` varchar(3072) NOT NULL,
  `template_name` varchar(320) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `ref` (`ref`),
  KEY `notifications_ix0` (`ref`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `oc_preferences`
--

DROP TABLE IF EXISTS `oc_preferences`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `oc_preferences` (
  `userid` varchar(64) COLLATE utf8_bin NOT NULL DEFAULT '',
  `appid` varchar(32) COLLATE utf8_bin NOT NULL DEFAULT '',
  `configkey` varchar(64) COLLATE utf8_bin NOT NULL DEFAULT '',
  `configvalue` longtext COLLATE utf8_bin,
  PRIMARY KEY (`userid`,`appid`,`configkey`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `oc_sciencemesh`
--

DROP TABLE IF EXISTS `oc_sciencemesh`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `oc_sciencemesh` (
  `iopurl` varchar(255) COLLATE utf8_bin NOT NULL DEFAULT 'http://localhost:10999',
  `country` varchar(2) COLLATE utf8_bin NOT NULL,
  `hostname` varchar(255) COLLATE utf8_bin NOT NULL DEFAULT 'example.org',
  `sitename` varchar(255) COLLATE utf8_bin NOT NULL DEFAULT 'CERNBox',
  `siteurl` varchar(255) COLLATE utf8_bin NOT NULL DEFAULT 'http://localhost',
  `numusers` bigint(20) unsigned DEFAULT '0',
  `numfiles` bigint(20) unsigned DEFAULT '0',
  `numstorage` bigint(20) unsigned DEFAULT '0'
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `oc_share`
--

DROP TABLE IF EXISTS `oc_share`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `oc_share` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `share_type` smallint(6) NOT NULL DEFAULT '0',
  `share_with` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `uid_owner` varchar(64) COLLATE utf8_bin NOT NULL DEFAULT '',
  `uid_initiator` varchar(64) COLLATE utf8_bin DEFAULT NULL,
  `parent` int(11) DEFAULT NULL,
  `item_type` varchar(64) COLLATE utf8_bin NOT NULL DEFAULT '',
  `item_source` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `item_target` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `file_source` bigint(20) unsigned DEFAULT NULL,
  `file_target` varchar(512) COLLATE utf8_bin DEFAULT NULL,
  `permissions` smallint(6) NOT NULL DEFAULT '0',
  `stime` bigint(20) NOT NULL DEFAULT '0',
  `accepted` smallint(6) NOT NULL DEFAULT '0',
  `expiration` datetime DEFAULT NULL,
  `token` varchar(32) COLLATE utf8_bin DEFAULT NULL,
  `mail_send` smallint(6) NOT NULL DEFAULT '0',
  `fileid_prefix` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `orphan` tinyint(4) DEFAULT NULL,
  `share_name` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `quicklink` tinyint(1) NOT NULL DEFAULT '0',
  `description` varchar(1024) COLLATE utf8_bin NOT NULL DEFAULT '',
  `internal` tinyint(1) NOT NULL DEFAULT '0',
  `notify_uploads_extra_recipients` varchar(2048) COLLATE utf8_bin DEFAULT NULL,
  `notify_uploads` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  KEY `item_share_type_index` (`item_type`,`share_type`),
  KEY `file_source_index` (`file_source`),
  KEY `token_index` (`token`),
  KEY `index_internal` (`internal`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=336200 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `oc_share_status`
--

DROP TABLE IF EXISTS `oc_share_status`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `oc_share_status` (
  `id` int(11) NOT NULL,
  `recipient` varchar(255) NOT NULL,
  `state` int(11) DEFAULT '0',
  PRIMARY KEY (`id`,`recipient`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `ocm_access_method_webapp`
--

DROP TABLE IF EXISTS `ocm_access_method_webapp`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `ocm_access_method_webapp` (
  `ocm_access_method_id` int(11) NOT NULL,
  `view_mode` int(11) NOT NULL,
  PRIMARY KEY (`ocm_access_method_id`),
  CONSTRAINT `ocm_access_method_webapp_ibfk_1` FOREIGN KEY (`ocm_access_method_id`) REFERENCES `ocm_shares_access_methods` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `ocm_access_method_webdav`
--

DROP TABLE IF EXISTS `ocm_access_method_webdav`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `ocm_access_method_webdav` (
  `ocm_access_method_id` int(11) NOT NULL,
  `permissions` int(11) NOT NULL,
  PRIMARY KEY (`ocm_access_method_id`),
  CONSTRAINT `ocm_access_method_webdav_ibfk_1` FOREIGN KEY (`ocm_access_method_id`) REFERENCES `ocm_shares_access_methods` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `ocm_protocol_transfer`
--

DROP TABLE IF EXISTS `ocm_protocol_transfer`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `ocm_protocol_transfer` (
  `ocm_protocol_id` int(11) NOT NULL,
  `source_uri` varchar(255) NOT NULL,
  `shared_secret` varchar(255) NOT NULL,
  `size` int(11) NOT NULL,
  PRIMARY KEY (`ocm_protocol_id`),
  CONSTRAINT `ocm_protocol_transfer_ibfk_1` FOREIGN KEY (`ocm_protocol_id`) REFERENCES `ocm_received_share_protocols` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `ocm_protocol_webapp`
--

DROP TABLE IF EXISTS `ocm_protocol_webapp`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `ocm_protocol_webapp` (
  `ocm_protocol_id` int(11) NOT NULL,
  `uri_template` varchar(255) NOT NULL,
  `view_mode` int(11) NOT NULL,
  PRIMARY KEY (`ocm_protocol_id`),
  CONSTRAINT `ocm_protocol_webapp_ibfk_1` FOREIGN KEY (`ocm_protocol_id`) REFERENCES `ocm_received_share_protocols` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `ocm_protocol_webdav`
--

DROP TABLE IF EXISTS `ocm_protocol_webdav`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `ocm_protocol_webdav` (
  `ocm_protocol_id` int(11) NOT NULL,
  `uri` varchar(255) NOT NULL,
  `shared_secret` text NOT NULL,
  `permissions` int(11) NOT NULL,
  PRIMARY KEY (`ocm_protocol_id`),
  CONSTRAINT `ocm_protocol_webdav_ibfk_1` FOREIGN KEY (`ocm_protocol_id`) REFERENCES `ocm_received_share_protocols` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `ocm_received_share_protocols`
--

DROP TABLE IF EXISTS `ocm_received_share_protocols`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `ocm_received_share_protocols` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `ocm_received_share_id` int(11) NOT NULL,
  `type` tinyint(4) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `ocm_received_share_id` (`ocm_received_share_id`,`type`),
  CONSTRAINT `ocm_received_share_protocols_ibfk_1` FOREIGN KEY (`ocm_received_share_id`) REFERENCES `ocm_received_shares` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=210 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `ocm_received_shares`
--

DROP TABLE IF EXISTS `ocm_received_shares`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `ocm_received_shares` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `item_type` tinyint(4) NOT NULL,
  `share_with` varchar(255) NOT NULL,
  `owner` varchar(255) NOT NULL,
  `initiator` varchar(255) NOT NULL,
  `ctime` int(11) NOT NULL,
  `mtime` int(11) NOT NULL,
  `expiration` int(11) DEFAULT NULL,
  `type` tinyint(4) NOT NULL,
  `state` tinyint(4) NOT NULL,
  `remote_share_id` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=136 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `ocm_remote_users`
--

DROP TABLE IF EXISTS `ocm_remote_users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `ocm_remote_users` (
  `initiator` varchar(255) NOT NULL,
  `opaque_user_id` varchar(255) NOT NULL,
  `idp` varchar(255) NOT NULL,
  `email` varchar(255) NOT NULL,
  `display_name` varchar(255) NOT NULL,
  PRIMARY KEY (`initiator`,`opaque_user_id`,`idp`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `ocm_shares`
--

DROP TABLE IF EXISTS `ocm_shares`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `ocm_shares` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `token` varchar(255) NOT NULL,
  `fileid_prefix` varchar(64) NOT NULL,
  `item_source` varchar(64) NOT NULL,
  `name` text NOT NULL,
  `share_with` varchar(255) NOT NULL,
  `owner` varchar(255) NOT NULL,
  `initiator` text NOT NULL,
  `ctime` int(11) NOT NULL,
  `mtime` int(11) NOT NULL,
  `expiration` int(11) DEFAULT NULL,
  `type` tinyint(4) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `token` (`token`),
  UNIQUE KEY `fileid_prefix` (`fileid_prefix`,`item_source`,`share_with`,`owner`)
) ENGINE=InnoDB AUTO_INCREMENT=192 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `ocm_shares_access_methods`
--

DROP TABLE IF EXISTS `ocm_shares_access_methods`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `ocm_shares_access_methods` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `ocm_share_id` int(11) NOT NULL,
  `type` tinyint(4) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `ocm_share_id` (`ocm_share_id`,`type`),
  CONSTRAINT `ocm_shares_access_methods_ibfk_1` FOREIGN KEY (`ocm_share_id`) REFERENCES `ocm_shares` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=333 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `ocm_tokens`
--

DROP TABLE IF EXISTS `ocm_tokens`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `ocm_tokens` (
  `token` varchar(255) NOT NULL,
  `initiator` varchar(255) NOT NULL,
  `expiration` datetime NOT NULL,
  `description` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`token`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

