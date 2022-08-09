A CRUD application with `gorm.io/gorm` and `gorilla/mux`.

### References
- [Build a REST API in Golang with MySQL, GORM and Gorilla Mux](https://www.soberkoder.com/go-rest-api-mysql-gorm/)

Notes: Follow the steps in [Tutorial: Accessing a relational database - Set up a database](https://go.dev/doc/tutorial/database-access#set_up_database) to set up the database. Change the table name `album` to `albums` as `gorm.DB.AutoMigrate()` create table name in plurals. 