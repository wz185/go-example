CREATE DATABASE goexample;
USE goexample;

CREATE TABLE `value_product_pause` (
  `value_product_pause_id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `value_product_id` int(10) NOT NULL,
  `account_id` int(11) unsigned NOT NULL,
  `type` tinyint(3) unsigned NOT NULL,
  `status` tinyint(3) unsigned NOT NULL,
  `requested_start_date` date NOT NULL,
  `requested_end_date` date NOT NULL,
  `actual_end_date` datetime NOT NULL,
  `canceled_date` datetime NOT NULL,
  `created_by_date` datetime NOT NULL,
  `created_by` int(11) unsigned NOT NULL,
  `modified_by_date` datetime NOT NULL,
  `modified_by` int(11) unsigned NOT NULL,
  PRIMARY KEY (`value_product_pause_id`),
  KEY `value_product_id` (`value_product_id`),
  KEY `account_id` (`account_id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8;