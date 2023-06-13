CREATE
DATABASE IF NOT EXISTS workspace_booking;

USE workspace_booking;

CREATE TABLE `rooms`
(
    `RoomId`   INT AUTO_INCREMENT PRIMARY KEY,
    `Name`     VARCHAR(50) NOT NULL,
    `MaxSeats` INT         NOT NULL
);

CREATE TABLE `workspaces`
(
    `WorkspaceId`           INT AUTO_INCREMENT PRIMARY KEY,
    `RoomId`                INT,
    `Name`                  VARCHAR(50) NOT NULL,
    `DockingStationPresent` TINYINT(1) DEFAULT 0,
    `AdjustableDeskPresent` TINYINT(1) DEFAULT 0,
    `NumberOfMonitors`      INT DEFAULT 1,
    FOREIGN KEY (`RoomId`) REFERENCES `rooms` (`RoomId`)
);

INSERT INTO `rooms` (`RoomId`, `Name`, `MaxSeats`)
VALUES (1, 'Raum 01', 10),
       (2, 'Raum 02', 10),
       (3, 'Raum 03', 10),
       (4, 'Raum 04', 10),
       (5, 'Raum 05', 10);

INSERT INTO `workspaces` (`WorkspaceId`, `RoomId`, `Name`, `DockingStationPresent`, `AdjustableDeskPresent`,
                          `NumberOfMonitors`)
VALUES (1, 1, 'Arbeitsplatz 01', 1, 1, 1),
       (2, 1, 'Arbeitsplatz 02', 0, 1, 2),
       (3, 1, 'Arbeitsplatz 03', 1, 1, 2),
       (4, 1, 'Arbeitsplatz 04', 0, 0, 1),
       (5, 1, 'Arbeitsplatz 05', 1, 1, 2),
       (6, 1, 'Arbeitsplatz 06', 0, 0, 2),
       (7, 1, 'Arbeitsplatz 07', 1, 1, 1),
       (8, 1, 'Arbeitsplatz 08', 0, 0, 2),
       (9, 1, 'Arbeitsplatz 09', 1, 1, 1),
       (10, 1, 'Arbeitsplatz 10', 0, 0, 1),
       (11, 2, 'Arbeitsplatz 01', 1, 1, 1),
       (12, 2, 'Arbeitsplatz 02', 1, 0, 2),
       (13, 2, 'Arbeitsplatz 03', 0, 1, 3),
       (14, 2, 'Arbeitsplatz 04', 0, 1, 1),
       (15, 2, 'Arbeitsplatz 05', 1, 0, 2),
       (16, 2, 'Arbeitsplatz 06', 1, 1, 3),
       (17, 2, 'Arbeitsplatz 07', 0, 0, 1),
       (18, 2, 'Arbeitsplatz 08', 1, 0, 2),
       (19, 2, 'Arbeitsplatz 09', 1, 0, 1),
       (20, 2, 'Arbeitsplatz 10', 1, 0, 1),
       (21, 3, 'Arbeitsplatz 01', 1, 1, 1),
       (22, 3, 'Arbeitsplatz 02', 0, 1, 2),
       (23, 3, 'Arbeitsplatz 03', 0, 1, 1),
       (24, 3, 'Arbeitsplatz 04', 0, 1, 1),
       (25, 3, 'Arbeitsplatz 05', 1, 1, 2),
       (26, 3, 'Arbeitsplatz 06', 1, 0, 2),
       (27, 3, 'Arbeitsplatz 07', 1, 1, 1),
       (28, 3, 'Arbeitsplatz 08', 1, 1, 2),
       (29, 3, 'Arbeitsplatz 09', 0, 1, 2),
       (30, 3, 'Arbeitsplatz 10', 0, 1, 1),
       (31, 4, 'Arbeitsplatz 01', 0, 1, 1),
       (32, 4, 'Arbeitsplatz 02', 1, 1, 2),
       (33, 4, 'Arbeitsplatz 03', 1, 1, 2),
       (34, 4, 'Arbeitsplatz 04', 1, 1, 1),
       (35, 4, 'Arbeitsplatz 05', 1, 0, 2),
       (36, 4, 'Arbeitsplatz 06', 0, 1, 1),
       (37, 4, 'Arbeitsplatz 07', 0, 0, 1),
       (38, 4, 'Arbeitsplatz 08', 1, 1, 2),
       (39, 4, 'Arbeitsplatz 09', 1, 1, 2),
       (40, 4, 'Arbeitsplatz 10', 1, 1, 1),
       (41, 5, 'Arbeitsplatz 01', 0, 0, 1),
       (42, 5, 'Arbeitsplatz 02', 1, 1, 2),
       (43, 5, 'Arbeitsplatz 03', 1, 0, 2),
       (44, 5, 'Arbeitsplatz 04', 1, 1, 1),
       (45, 5, 'Arbeitsplatz 05', 1, 1, 2),
       (46, 5, 'Arbeitsplatz 06', 0, 1, 1),
       (47, 5, 'Arbeitsplatz 07', 0, 1, 1),
       (48, 5, 'Arbeitsplatz 08', 0, 1, 2),
       (49, 5, 'Arbeitsplatz 09', 1, 1, 1),
       (50, 5, 'Arbeitsplatz 10', 1, 0, 1);