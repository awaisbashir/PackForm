import psycopg2
import glob
conn = psycopg2.connect(database="hague",
                        user='postgres', password='Techleadz12*', 
                        host='localhost', port='5432'
)

conn.autocommit = False
cur = conn.cursor()
csvPath = "./test_data/"

try:

    timeZ = 'SET TimeZone = "Australia/Melbourne"'
    cur.execute(timeZ)

    # Loop through each CSV
    for filename in glob.glob(csvPath+"*.csv"):
        
        # Create a table name
        table = filename.replace("./test_data\\", "").replace("Test task - Postgres -", "").replace(".csv", "").strip()
        print(table)
        
        # Open file
        with open(filename, 'r') as f:
        
            # Extract the headers
            headers = f.readline().strip()
        
            # Split columns into array [...]
            columns = headers.split(",")
            
            # Create query to drop or create table
            tblQuery = 'DROP TABLE IF EXISTS '+ table + ";\n"
            tblQuery += 'CREATE TABLE '+ table + "("
            
            # Loop through columns definition
            for column in columns:
                if column == 'id':
                    tblQuery += column + " INT PRIMARY KEY,\n"
                elif column.find('quantity') != -1 or column == 'price_per_unit':
                    tblQuery += column + " NUMERIC NULL,\n"
                elif column == 'created_at':
                    tblQuery += column + " timestamp with time zone,\n"
                elif column == 'user_id' or column == 'customer_id':
                    tblQuery += column + " VARCHAR(255),\n"
                elif column.find('id') != -1:
                    if table == 'customer_companies':
                        tblQuery += column + " INT PRIMARY KEY,\n"
                    else:
                        tblQuery += column + " INT NOT NULL,\n"
                else:
                    tblQuery += column + " VARCHAR(255),\n"
                    
            tblQuery = tblQuery[:-2]
            tblQuery += ");"
            
            cur.execute(tblQuery)
            
            copy_cmd = f"copy {table}({headers}) from stdout (format csv)"
            cur.copy_expert(copy_cmd, f)
            
            conn.commit()

except Exception as err:
    # reverting changes because of exception
    conn.rollback()
    
    print("Failed to update record to database rollback: {}".format(err))
    
finally:
    cur.close()
    conn.close()
    print("connection is closed")
    
    
    