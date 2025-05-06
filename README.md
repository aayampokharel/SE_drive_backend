# SE_drive Backend - Backend

## Overview

The **SE_drive** is a robust and scalable backend solution built using **Go (Golang)** and **MySQL**. It's designed to handle file uploads, storage, and retrieval efficiently. The backend mimics a file storage system, offering powerful features like **file management**, **history tracking**, and **dynamic resolution adjustment** of media files using **FFmpeg**. This solution ensures seamless interaction between the client and server while maintaining high performance, scalability, and security.

---

## Key Features

- **File Upload and Transfer**: Secure and efficient handling of various file types (PDFs, images, videos, etc.) through multipart request.

- **FFmpeg Integration**: Dynamically adjust the resolution of uploaded media files to optimize storage and enhance performance.

- **History Retrieval**: Fetch and display the history of uploaded and downloaded files from the database.

- **Database Management**: MySQL database used for storing user data, file metadata, and file history for easy retrieval.

- **RESTful API**: A clean, scalable API to handle file operations and history retrieval.

- **Scalable**: Optimized to handle multiple simultaneous uploads/downloads and user queries with ease.

---

## Features in Detail

### 1. **File Upload**

Users can upload various types of files (documents, images, videos, etc.). 

### 2. **FFmpeg for Media Optimization**

For media files such as images and videos, **FFmpeg** is used to adjust their resolution before storing them in the database. This ensures that the media files are stored in optimal sizes without compromising quality, optimizing both storage space and application performance.

### 3. **File History**

The backend allows users to fetch the history of files they’ve uploaded or downloaded during login. 
---




