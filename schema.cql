CREATE KEYSPACE demo
WITH replication = {
'class':'SimpleStrategy',
'replication_factor':1
};

USE demo;

CREATE TABLE stock_data (
symbol TEXT,
datetime TIMESTAMP,
open DOUBLE,
high DOUBLE,
low DOUBLE,
close DOUBLE,
volume BIGINT,
PRIMARY KEY(symbol, datetime)
)
WITH CLUSTERING ORDER BY(datetime ASC);