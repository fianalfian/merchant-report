CREATE TABLE IF NOT EXISTS `Outlets` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT,
    `merchant_id` bigint(20) NOT NULL,
    `outlet_name` varchar(40) NOT NULL,
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `created_by` bigint(20) NOT NULL,
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_by` bigint(20) NOT NULL,
    PRIMARY KEY (`id`)
  ) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=latin1;