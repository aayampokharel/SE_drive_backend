create database if not exists ctrlPlusSave;
use ctrlPlusSave;

create table if not exists UserInfoTable(
email varchar(60) Primary key,
userName varchar(60) not null,
userPassword varchar(60) not null,
isSubscribed bool not null,
token varchar(50) unique not null
);

insert into UserInfoTable values ("aayam.pokharel@gmail.com","aayam Pokharel","ktm123");
-- when user is subscribed , dont refer to this table , if not then only refer here . 


create table if not exists TrialsTable(
token varchar(50) primary key,
trialsLeft int not null

);

create table if not exists VideoTable(
count int AUTO_INCREMENT primary key,
token varchar(50) not null,
videoFileName varchar(150) not null

);

create table if not exists PhotoTable(
count int AUTO_INCREMENT primary key,
token varchar(50) not null,
PhotoFileName varchar(150) not null

);

create table if not exists PdfTable(
count int AUTO_INCREMENT primary key,
token varchar(50) not null,
pdfFileName varchar(150) not null

);
create table if not exists AudioTable(
count int AUTO_INCREMENT primary key,
token varchar(50) not null,
audioFileName varchar(150) not null

);
create table if not exists TextTable(
count int AUTO_INCREMENT primary key,
token varchar(50) not null,
textFileName varchar(150) not null

);

select * from UserInfoTable;

drop table UserInfoTable;
