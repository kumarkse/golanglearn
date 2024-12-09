import pymysql
import os
from dotenv import load_dotenv


#pip install python-dotenv
#pip install PyMySQL


# Function to test database connectivity
def test_database_connectivity():
    load_dotenv()
    
    db_host = os.getenv('DB_HOST')
    db_user = os.getenv('DB_USER')
    db_password = os.getenv('DB_PASSWORD')
    db_name = os.getenv('DB_NAME')
    
    try:
        connection = pymysql.connect(
            host=db_host,
            user=db_user,
            password=db_password,
            database=db_name
        )
        with connection.cursor() as cursor:
            cursor.execute("SELECT 1")
            result = cursor.fetchone()
        connection.close()
        return True, 'Database connectivity is working.'
    except Exception as e:
        return False, f'Failed to connect to the database: {str(e)}'

# For local testing
if __name__ == "__main__":
    db_status, db_message = test_database_connectivity()
    print(f'Database status: {db_status}, Message: {db_message}')
