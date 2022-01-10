create table admin
(
    column_1   int          null,
    id         int          not null
        primary key,
    number     varchar(255) null comment '管理员编号',
    name       varchar(255) null comment '姓名',
    email      varchar(255) null comment '邮箱',
    phone      varchar(255) null comment '电话号码',
    password   varchar(255) null comment '密码',
    created_at timestamp    null,
    updated_at timestamp    null,
    deleted_at timestamp    null,
    last_modify_admin_id    int         null comment '最后修改记录管理员id',
    create_admin_id    int         null comment '创建记录的管理员id'
)
    comment '管理员';

create table article
(
    di         int          not null
        primary key,
    title      varchar(255) null comment '标题',
    author     int          null comment '作者，对应于用户表的主键id',
    created_at timestamp    null,
    updated_at timestamp    null,
    deleted_at timestamp    null,
    tag        varchar(255) null comment '标签简介',
    code       varchar(255) null comment '编号',
    visibility int          null comment '可见性，自己可见：1，公共:2',
     last_modify_admin_id    int         null comment '最后修改记录管理员id',
    create_admin_id    int         null comment '创建记录的管理员id'
)
    comment '文章';

create table comment
(
    id           int auto_increment
        primary key,
    code         varchar(255) null comment '编号',
    introduction varchar(255) null comment '简介',
    article_id   int          null comment '文章id',
     created_at timestamp    null,
    updated_at timestamp    null,
    deleted_at timestamp    null,
     last_modify_admin_id    int         null comment '最后修改记录管理员id',
    create_admin_id    int         null comment '创建记录的管理员id'
)
    comment '评论';

create table examine
(
    id         int auto_increment
        primary key,
    admin_id   int          null,
    article_id int          null,
    result     int          null comment '审核结果，1：通过，2：不通过',
    reason     varchar(255) null comment '原因',
     created_at timestamp    null,
    updated_at timestamp    null,
    deleted_at timestamp    null,
     last_modify_admin_id    int         null comment '最后修改记录管理员id',
    create_admin_id    int         null comment '创建记录的管理员id'
)
    comment '审核记录';

create table `grant`
(
    id            int auto_increment
        primary key,
    admin_id      int null comment '管理员id',
    permission_id int null comment '权限id',
     created_at timestamp    null,
    updated_at timestamp    null,
    deleted_at timestamp    null,
     last_modify_admin_id    int         null comment '最后修改记录管理员id',
    create_admin_id    int         null comment '创建记录的管理员id'
)
    comment '权限';

create table notice
(
    id           int auto_increment
        primary key,
    introduction varchar(255) null comment '简介',
    content      varchar(255) null,
    user_id      int          null,
    admin_id     int          null,
     created_at timestamp    null,
    updated_at timestamp    null,
    deleted_at timestamp    null,
     last_modify_admin_id    int         null comment '最后修改记录管理员id',
    create_admin_id    int         null comment '创建记录的管理员id'
)
    comment '通知，通知是管理员或系统给用户发的信息';

create table permission
(
    id           int auto_increment
        primary key,
    introduction int null comment '权限简介',
    name         int null comment '权限名',
     created_at timestamp    null,
    updated_at timestamp    null,
    deleted_at timestamp    null,
     last_modify_admin_id    int         null comment '最后修改记录管理员id',
    create_admin_id    int         null comment '创建记录的管理员id'
)
    comment '权限';

create table private_chat
(
    id          int auto_increment
        primary key,
    sender_id   int          null comment '发送方id',
    accepter_id int          null,
    content     varchar(255) null,
     created_at timestamp    null,
    updated_at timestamp    null,
    deleted_at timestamp    null,
     last_modify_admin_id    int         null comment '最后修改记录管理员id',
    create_admin_id    int         null comment '创建记录的管理员id'
)
    comment '私信';

create table reply
(
    id           int auto_increment
        primary key,
    comment_id   int          null comment '评论编号',
    introduction varchar(255) null comment '简介',
    code         varchar(255) null comment '编号',
     created_at timestamp    null,
    updated_at timestamp    null,
    deleted_at timestamp    null,
     last_modify_admin_id    int         null comment '最后修改记录管理员id',
    create_admin_id    int         null comment '创建记录的管理员id'
)
    comment '回复';

create table user
(
    id         int auto_increment
        primary key,
    number     varchar(255) null comment '用户编号',
    name       varchar(255) null comment '姓名
',
    email      varchar(255) null comment '邮箱',
    phone      varchar(255) null comment '电话号码
',
    created_at timestamp    null,
    updated_at timestamp    null,
    deleted_at timestamp    null,
    password   varchar(255) null comment '密码',
    portrait   varchar(255) null comment '头像url',
     last_modify_admin_id    int         null comment '最后修改记录管理员id',
    create_admin_id    int         null comment '创建记录的管理员id'
)
    comment '博客用户';

