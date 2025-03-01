-- Create and use the database
create database if not exists ctrlPlusSave;
use ctrlPlusSave;

-- Create the UserInfoTable
create table if not exists UserInfoTable(
    email varchar(60) Primary key,
    userName varchar(60) not null,
    userPassword varchar(60) not null,
    isSubscribed bool not null,
    token varchar(50) unique not null
);

-- Insert data into UserInfoTable
INSERT INTO UserInfoTable (email, userName, userPassword, isSubscribed, token) 
VALUES 
("testuser1@gmail.com", "Test User 1", "pass123", FALSE, "token123"),
("testuser2@gmail.com", "Test User 2", "pass456", TRUE, "token456");

-- Create the TrialsTable
create table if not exists TrialsTable(
    token varchar(50) primary key,
    trialsLeft int not null
);

-- Insert data into TrialsTable
INSERT INTO TrialsTable (token, trialsLeft) 
VALUES 
("token123", 5),
("token456", 10);

-- Create VideoTable
create table if not exists VideoTable(
    count int AUTO_INCREMENT primary key,
    token varchar(50) not null,
    originalVideoFileName varchar(150) unique not null,
    outputVideoFileName varchar(150) unique not null
);

-- Insert data into VideoTable
INSERT INTO VideoTable (token, originalVideoFileName, outputVideoFileName) 
VALUES 
("token123", "original_video1.mp4", "output_video1.mp4"),
("token456", "original_video2.mp4", "output_video2.mp4");

-- Create PhotoTable
create table if not exists PhotoTable(
    count int AUTO_INCREMENT primary key,
    token varchar(50) not null,
    originalPhotoFileName varchar(150)unique not null,
    outputPhotoFileName varchar(150) unique not null
);

-- Insert data into PhotoTable
INSERT INTO PhotoTable (token, originalPhotoFileName, outputPhotoFileName) 
VALUES 
("token123", "original_photo1.jpg", "output_photo1.jpg"),
("token456", "original_photo2.jpg", "output_photo2.jpg");

-- Create PdfTable
create table if not exists PdfTable(
    count int AUTO_INCREMENT primary key,
    token varchar(50) not null,
    originalPdfFileName varchar(150)unique not null,
    outputPdfFileName varchar(150) unique not null
);

-- Insert data into PdfTable
INSERT INTO PdfTable (token, originalPdfFileName, outputPdfFileName) 
VALUES 
("token123", "original_doc1.pdf", "output_doc1.pdf"),
("token456", "original_doc2.pdf", "output_doc2.pdf");

-- Create AudioTable
create table if not exists AudioTable(
    count int AUTO_INCREMENT primary key,
    token varchar(50) not null,
    originalAudioFileName varchar(150)unique not null,
    outputAudioFileName varchar(150)unique not null
);

-- Insert data into AudioTable
INSERT INTO AudioTable (token, originalAudioFileName, outputAudioFileName) 
VALUES 
("token123", "original_audio1.mp3", "output_audio1.mp3"),
("token456", "original_audio2.mp3", "output_audio2.mp3");

-- Create TextTable
create table if not exists TextTable(
    count int AUTO_INCREMENT primary key,
    token varchar(50) not null,
    originalTextFileName varchar(150)unique not null,
    outputTextFileName varchar(150)unique not null
);

-- Insert data into TextTable
INSERT INTO TextTable (token, originalTextFileName, outputTextFileName) 
VALUES 
("token123", "original_text1.txt", "output_text1.txt"),
("token456", "original_text2.txt", "output_text2.txt");

INSERT INTO TextTable (token, originalTextFileName, outputTextFileName) 
VALUES 
("token123", "original_importantFile.txt", "output_importantFile.txt");
-- Select query to see file names based on subscription status
select * from usert;
SELECT 
    u.userName, 
    u.isSubscribed,
    u.token,  
    CASE 
        WHEN u.isSubscribed = FALSE THEN v.outputVideoFileName
        ELSE v.originalVideoFileName
    END AS VideoFileName,

    CASE 
        WHEN u.isSubscribed = FALSE THEN p.outputPhotoFileName
        ELSE p.originalPhotoFileName
    END AS PhotoFileName,

    CASE 
        WHEN u.isSubscribed = FALSE THEN pdf.outputPdfFileName
        ELSE pdf.originalPdfFileName
    END AS PdfFileName,

    CASE 
        WHEN u.isSubscribed = FALSE THEN a.outputAudioFileName
        ELSE a.originalAudioFileName
    END AS AudioFileName,

    CASE 
        WHEN u.isSubscribed = FALSE THEN t.outputTextFileName
        ELSE t.originalTextFileName
    END AS TextFileName

FROM UserInfoTable u
LEFT JOIN VideoTable v ON u.token = v.token
LEFT JOIN PhotoTable p ON u.token = p.token
LEFT JOIN PdfTable pdf ON u.token = pdf.token
LEFT JOIN AudioTable a ON u.token = a.token
LEFT JOIN TextTable t ON u.token = t.token
WHERE u.email = "testuser1@gmail.com";  -- Change to the desired email
select * from userinfotable;
-- drop database ctrlplussave;