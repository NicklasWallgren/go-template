DROP PROCEDURE IF EXISTS go_template_test.truncate_tables;
DELIMITER ;;
CREATE PROCEDURE go_template_test.truncate_tables(theDb varchar(64))
BEGIN
    declare tname varchar(64);
    declare done TINYINT DEFAULT FALSE;
    declare tcursor CURSOR FOR
SELECT table_name FROM information_schema.tables WHERE table_type <> 'VIEW' AND table_schema = theDb AND table_name != 'goose_db_version';
DECLARE CONTINUE HANDLER FOR NOT FOUND SET done = true;
    SET FOREIGN_KEY_CHECKS = 0;

OPEN tcursor;
l1: LOOP
        FETCH tcursor INTO tname;
        if done or tname = NULL then leave l1; end if;
        set @sql = CONCAT('truncate `', theDB, '`.`', tname, '`');
PREPARE stmt from @sql;
EXECUTE stmt;
DEALLOCATE PREPARE stmt;
END LOOP l1;
CLOSE tcursor;
SET FOREIGN_KEY_CHECKS = 1;
END ;;
DELIMITER ;