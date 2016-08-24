#dump ways

mysqldump db_name table_name > table_name.sql

#dump from remote db

mysqldump -u username -h db_host -p db_name table_name > table_name.sql

#restore

mysql -u username -p db_name
mysql> source <full_path>/table_name.sql

#or
mysql -u username -p db_name < /path/to/table_name.sql

#dump restore from compressed

mysqldump db_name table_name | gzip > table_name.sql.gz

gunzip < table_name.sql.gz | mysql -u username -p db_name
