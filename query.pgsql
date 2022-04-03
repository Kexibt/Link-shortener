CREATE TABLE Urls (
    ID      INTEGER     PRIMARY KEY,
    longURL     NCHAR()     NOT NULL
); 

INSERT INTO urls(id, longURL) 
VALUES (4, 'https://docs.google.com');

INSERT INTO urls(id, longURL) 
VALUES (2, 'https://route256.ozon.ru');

ALTER TABLE urls 
ADD COLUMN longURL NCHAR(1000) NOT NULL;

DELETE FROM urls
WHERE id = 4;

SELECT * FROM urls 
WHERE longurl LIKE '%https://route256.ozon.ru%'