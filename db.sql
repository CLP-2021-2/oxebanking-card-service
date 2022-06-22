CREATE TABLE `card` (
    `num_card`    int(16)      NOT NULL     unsigned    AUTO_INCREMENT,
    `cod_seg`     int(3)       NOT NULL     unsigned    AUTO_INCREMENT,
    `name`        varchar(30)  NOT NULL,
    `date_venc`   varchar(10)  NOT NULL,
    `status`      varchar(10)  NOT NULL,
    PRIMARY KEY   (`num_card`)
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=UTF-8;