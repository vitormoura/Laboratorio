CREATE TABLE CAD_BOOKS (

	COD_BOOK CHAR(64) 			NOT NULL PRIMARY KEY,
	NOM_BOOK VARCHAR(255) 		NOT NULL,
	TXT_PATH_BOOK VARCHAR(500) 	NOT NULL,
	DTH_CADASTRO DATETIME 		NOT NULL,
	QTD_BYTES_BOOK INT			NOT NULL
);