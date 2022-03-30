create table test(
    -> id int unsigned primary key auto_increment not null,
    -> name varchar(20) default '',
    -> age tinyint unsigned default 0,
    ->  gender enum('男','女','中性','保密') default '保密'));