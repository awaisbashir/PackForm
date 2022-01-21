# Packform web app with PostgreSQL, golang and vuejs

## Table of contents 👀
* [Initialize DB](#general-info)
* [The GOPOST SERVER](#technologies)
* [Vue Client](#blog)
* [Setup](#setup)

### Directory structure should be same as given

### Initialize DB
To initialize DB, you must have install postgresql, and python on your local machine. 
You can install python via ananconda. After installation, install `pip`, and then run these two commands.

```
pip install psycopg2
pip install glob
```

You must have postgresql installed. And a database created with name `packform`

```
database: packform
user: postgresql
password: Techleadz12*
port: 5432
```
### If you have different configurations, please modify credentials accordingly in `seeds.py` `line 3`:
### All the csv files should be located in ./test_data folder with the same name as provided.

The script is written in python, and is located in packform folder with name `seeds.py`
You can run the script in `spyder` if you have install ananconda.

#### The GOPOST SERVER 🍵
You must have go installed on your local machine.
https://go.dev/doc/install

You must have the above mentioned DB credentials, otherwise you will need to modify `.env` file connection string location in `go-post` main folder.

To run the server, please use the following command;

```
go run main.go
```

### Server must be running on port 8080

#### Vue Client ⚡
To start the client, you must have yarn/npm installed. You must have latest vue cli configured.

You will need to run the following commands in order;

```
yarn install
npm install --save material-design-icons-iconfont
npm i vue-time-date-range-picker
yarn run serve
```

### Application would be running at port 8081
### You can navigate to the application via the following link.
http://localhost:8081/