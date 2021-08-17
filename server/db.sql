
----------------------------------------------------------------
create table IF NOT EXISTS `t_ca_server_info` (
    id bigint(20) NOT NULL AUTO_INCREMENT,
    name varchar(255) NOT NULL COMMENT 'ca name',
    host varchar(255) COMMENT 'host',
    ip_address varchar(255) NOT NULL,
    created_at timestamp ,
    updated_at timestamp,
    user_id bigint(20) NOT NULL DEFAULT -1
) Engine=InnoDB DEFAULT CHARSET=utf8 ;

--------------

create table IF NOT EXISTS `t_org_info`(
    id bigint(20) NOT NULL AUTO_INCREMENT,
    user_id bigint(20) NOT NULL DEFAULT -1,
    org_domain varchar(255) NOT NULL,
    peer_name varchar(255) NOT NULL,
    created_at timestamp ,
    updated_at timestamp
)Engine=InnoDB DEFAULT CHARSET=utf8 ;

----------------------------------------------------------------

create table IF NOT EXISTS `t_channel`(
    id bigint(20) NOT NULL AUTO_INCREMENT,
    channel_id bigint(20) NOT NULL,
    orgs varchar(255) NOT NULL,
    channel_name varchar(255) NOT NULL,
    user_id bigint(20) NOT NULL DEFAULT -1,
    user_account varchar(255) NOT NULL,
    created_at timestamp ,
    updated_at timestamp
)Engine=InnoDB DEFAULT CHARSET=utf8 ;


--------------

create table IF NOT EXISTS `t_users`(
    id bigint(20) NOT NULL AUTO_INCREMENT,
    usernames varchar(255) NOT NULL,
    pwd varchar(255) NOT NULL,
    ca_name varchar(255), 
    ca_pwd varchar(255) ,
    created_at timestamp ,
    updated_at timestamp 
)Engine=InnoDB DEFAULT CHARSET=utf8 ;