CREATE
DATABASE IF NOT EXISTS workspace_booking;

USE workspace_booking;

CREATE TABLE `rooms`
(
    `id`           INT AUTO_INCREMENT PRIMARY KEY,
    `created_at`   DATETIME,
    `updated_at`   DATETIME,
    `deleted_at`   DATETIME,
    `name`         VARCHAR(255),
    `MaxSeats`     INT
);

CREATE TABLE `workspaces`
(
    `id`                    INT AUTO_INCREMENT PRIMARY KEY,
    `created_at`            DATETIME,
    `updated_at`            DATETIME,
    `deleted_at`            DATETIME,
    `room_Id`                INT,
    `name`                  VARCHAR(255),
    `docking_station_present` TINYINT(1) DEFAULT 0,
    `adjustable_desk_present` TINYINT(1) DEFAULT 0,
    `has_two_screens`         TINYINT(1) DEFAULT 0,
    FOREIGN KEY (`room_Id`) REFERENCES `rooms` (`id`)
);

CREATE TABLE `bookings`
(
    `id`           INT AUTO_INCREMENT PRIMARY KEY,
    `created_at`   DATETIME,
    `updated_at`   DATETIME,
    `deleted_at`   DATETIME,
    `workspace_Id`  INT,
    `person_Id`     VARCHAR(255),
    `date`         DATETIME,
    FOREIGN KEY (`workspace_Id`) REFERENCES `workspaces` (`id`)

);

INSERT INTO `rooms` (`Id`, `Name`, `MaxSeats`)
VALUES (1, 'Raum 01', 10),
       (2, 'Raum 02', 10),
       (3, 'Raum 03', 10),
       (4, 'Raum 04', 10),
       (5, 'Raum 05', 10);

INSERT INTO `workspaces` (`id`, `room_Id`, `Name`, `docking_station_present`, `adjustable_desk_present`,
                          `has_two_screens`)
VALUES (1, 1, 'Arbeitsplatz 01', 1, 1, 1),
       (2, 1, 'Arbeitsplatz 02', 0, 1, 0),
       (3, 1, 'Arbeitsplatz 03', 1, 1, 0),
       (4, 1, 'Arbeitsplatz 04', 0, 0, 1),
       (5, 1, 'Arbeitsplatz 05', 1, 1, 0),
       (6, 1, 'Arbeitsplatz 06', 0, 0, 0),
       (7, 1, 'Arbeitsplatz 07', 1, 1, 1),
       (8, 1, 'Arbeitsplatz 08', 0, 0, 0),
       (9, 1, 'Arbeitsplatz 09', 1, 1, 1),
       (10, 1, 'Arbeitsplatz 10', 0, 0, 1),
       (11, 2, 'Arbeitsplatz 01', 1, 1, 1),
       (12, 2, 'Arbeitsplatz 02', 1, 0, 0),
       (13, 2, 'Arbeitsplatz 03', 0, 1, 1),
       (14, 2, 'Arbeitsplatz 04', 0, 1, 1),
       (15, 2, 'Arbeitsplatz 05', 1, 0, 0),
       (16, 2, 'Arbeitsplatz 06', 1, 1, 1),
       (17, 2, 'Arbeitsplatz 07', 0, 0, 1),
       (18, 2, 'Arbeitsplatz 08', 1, 0, 0),
       (19, 2, 'Arbeitsplatz 09', 1, 0, 1),
       (20, 2, 'Arbeitsplatz 10', 1, 0, 1),
       (21, 3, 'Arbeitsplatz 01', 1, 1, 1),
       (22, 3, 'Arbeitsplatz 02', 0, 1, 0),
       (23, 3, 'Arbeitsplatz 03', 0, 1, 1),
       (24, 3, 'Arbeitsplatz 04', 0, 1, 1),
       (25, 3, 'Arbeitsplatz 05', 1, 1, 0),
       (26, 3, 'Arbeitsplatz 06', 1, 0, 0),
       (27, 3, 'Arbeitsplatz 07', 1, 1, 1),
       (28, 3, 'Arbeitsplatz 08', 1, 1, 0),
       (29, 3, 'Arbeitsplatz 09', 0, 1, 0),
       (30, 3, 'Arbeitsplatz 10', 0, 1, 1),
       (31, 4, 'Arbeitsplatz 01', 0, 1, 1),
       (32, 4, 'Arbeitsplatz 02', 1, 1, 0),
       (33, 4, 'Arbeitsplatz 03', 1, 1, 0),
       (34, 4, 'Arbeitsplatz 04', 1, 1, 1),
       (35, 4, 'Arbeitsplatz 05', 1, 0, 0),
       (36, 4, 'Arbeitsplatz 06', 0, 1, 1),
       (37, 4, 'Arbeitsplatz 07', 0, 0, 1),
       (38, 4, 'Arbeitsplatz 08', 1, 1, 0),
       (39, 4, 'Arbeitsplatz 09', 1, 1, 0),
       (40, 4, 'Arbeitsplatz 10', 1, 1, 1),
       (41, 5, 'Arbeitsplatz 01', 0, 0, 1),
       (42, 5, 'Arbeitsplatz 02', 1, 1, 0),
       (43, 5, 'Arbeitsplatz 03', 1, 0, 0),
       (44, 5, 'Arbeitsplatz 04', 1, 1, 1),
       (45, 5, 'Arbeitsplatz 05', 1, 1, 0),
       (46, 5, 'Arbeitsplatz 06', 0, 1, 1),
       (47, 5, 'Arbeitsplatz 07', 0, 1, 1),
       (48, 5, 'Arbeitsplatz 08', 0, 1, 0),
       (49, 5, 'Arbeitsplatz 09', 1, 1, 1),
       (50, 5, 'Arbeitsplatz 10', 1, 0, 1);