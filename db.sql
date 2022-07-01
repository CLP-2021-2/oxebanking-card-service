create table card (
    num_card    int(16)      unsigned NOT NULL AUTO_INCREMENT,
    cod_seg     int(3)       unsigned NOT NULL,
    name        varchar(30)  NOT NULL,
    date_venc   varchar(10)  NOT NULL,
    status      varchar(10)  NOT NULL,
    PRIMARY KEY   (`num_card`)
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=latin1;