# Read Me

## Google Sign-in

    1. rename creds_example.json to creds.json

    2. update "google_client_id" and "google_client_secret" in creds.json

## Database 

    PostgreSQL

## DB Schema
  
    CREATE TABLE USERS(
        ID         SERIAL  PRIMARY KEY,
        Email      TEXT    NOT NULL UNIQUE,
        Password   TEXT    NOT NULL,
        RxPoint    integer DEFAULT 0,
        TxPoint    integer DEFAULT 0
    );
    
## Other Config
    
    file : config.json
    
    ## Port for server running 
    "WebServer":{
		"Port":"8081"
    }

    ## Tag "db" ia a config for connecting database
    ## Tag "master" is a config of database for writing
    ## Tag "replication" is a config of database for reading

    "db":{
	    "master":{
			"server": "127.0.0.1",
			"port": "5432",
			"db": "gift",
			"username": "postgres",
			"password": "8888",
			"parameter":{
				"charset": "utf8",
				"parseTime": "True",
				"loc": "Local"
			},
			"maxidleconns": 500,
			"maxopenconns": 500,
			"connmaxlifetime": 90
		},
		"replication":[
			{
				"server": "127.0.0.1",
				"port": "5432",
				"db": "gift",
				"username": "postgres",
				"password": "8888",
				"parameter":{
					"charset": "utf8",
					"parseTime": "True",
					"loc": "Local"
				},
				"maxidleconns": 500,
				"maxopenconns": 500,
				"connmaxlifetime": 90
			},
			{
				"server": "127.0.0.1",
				"port": "3306",
				"db": "gift",
				"username": "root",
				"password": "8888",
				"parameter":{
					"charset": "utf8",
					"parseTime": "True",
					"loc": "Local"
				},
				"maxidleconns": 500,
				"maxopenconns": 500,
				"connmaxlifetime": 90
			},
			{
				"server": "127.0.0.1",
				"port": "5432",
				"db": "gift",
				"username": "postgres",
				"password": "8888",
				"parameter":{
					"charset": "utf8",
					"parseTime": "True",
					"loc": "Local"
				},
				"maxidleconns": 500,
				"maxopenconns": 500,
				"connmaxlifetime": 90
			}						
		]
	}

## Function of current version  ( point transaction)  

    must go to user page and click button on the page to make a point transaction