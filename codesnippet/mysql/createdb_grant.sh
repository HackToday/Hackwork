mysql -u root -p -e 'create database mydb'
mysql -u root -p -e 'GRANT ALL PRIVILEGES ON mydb.* TO "root"@"%" IDENTIFIED BY "password";'
