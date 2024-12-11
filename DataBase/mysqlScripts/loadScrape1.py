import mysql.connector as con
from dotenv import load_dotenv
import json
import os
from pathlib import Path

load_dotenv()
current_dir = Path(__file__).parent

mydb = con.connect(host = "localhost" ,user = "root",password = os.getenv("mysqlPassword"))
dbcursor = mydb.cursor()
dbcursor.execute('use kumar;')


#******************************************************************************************************************************
# dbcursor.execute('create table backlog ( id int AUTO_INCREMENT primary key  , category varchar(50) , topic varchar(250) , link text );')

# backlog_table = {}

# json_path = current_dir / "../../webScraping/scrape1backlog.json"
# try:
#     with open(json_path,"r") as file:
#         backlog_table = json.load(file)  
# except FileNotFoundError:
#     print(f"File not found: {json_path}")
# except Exception as e:
#     print(f"An error occurred: {e}")

# cnt =0 
# for i,j in backlog_table.items():
#     for k,l in j.items():
#         print(type(i) , type(k) , type(l["backlogLink"]))
#         query = f'Insert into backlog(category, topic, link) values("{i}" , "{k}" , "{l["backlogLink"]}");'
#         print(query)
#         dbcursor.execute(query)

# dbcursor.execute('select * from backlog;')
# for i in dbcursor:
#     print(i)
#******************************************************************************************************************************

json_path = current_dir / "../../webScraping/main_scarped_1.json"
data = {}
with open(json_path,"r") as file:
    data = json.load(file)

dbcursor.execute('create table Scraped_data (id int auto_increment primary key , category varchar(50) , topic varchar(100) , content text )')

for category,val in data.items():
    





# dbcursor.execute('')
# dbcursor.execute('')
# dbcursor.execute('')
# dbcursor.execute('')
# dbcursor.execute('')
# dbcursor.execute('')



mydb.commit()