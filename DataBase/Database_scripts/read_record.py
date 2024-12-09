import pymysql
import os
from dotenv import load_dotenv
import random
from os.path import join, dirname

# Load environment variables from .env file
dotenv_path = join(dirname(__file__), '.env')
load_dotenv(dotenv_path)

# Function to establish database connection
def get_db_connection():
    try:
        conn = pymysql.connect(
            host=os.getenv('DB_HOST'),
            user=os.getenv('DB_USER'),
            passwd=os.getenv('DB_PASSWORD'),
            db=os.getenv('DB_NAME'),
            connect_timeout=5
        )
        print("Connection to MySQL database successful.")
        return conn
    except pymysql.MySQLError as e:
        print("ERROR: Unexpected error: Could not connect to MySQL instance.")
        print(e)
        raise

# Function to fetch one random record from the interns table
def fetch_random_record(connection):
    try:
        with connection.cursor() as cursor:
            # Get the total number of records
            count_query = "SELECT COUNT(*) FROM T1_categories"
            cursor.execute(count_query)
            total_records = cursor.fetchone()[0]
            
            if total_records == 0:
                print("No records found in the interns table.")
                return None

            # Fetch one random record
            random_offset = random.randint(0, total_records - 1)
            random_record_query = f"SELECT * FROM T1_categories LIMIT 1 OFFSET {random_offset}"
            cursor.execute(random_record_query)
            random_record = cursor.fetchone()
            
            return random_record
    except pymysql.MySQLError as e:
        print("ERROR: Could not fetch a random record.")
        print(e)
        raise

# Main function for local testing
if __name__ == "__main__":
    try:
        # Connect to the database
        connection = get_db_connection()

        # Fetch one random record from interns table
        random_record = fetch_random_record(connection)
        
        if random_record:
            print(f"Random record fetched from T1_categories table: {random_record}")
        
        # Close database connection
        connection.close()

    except Exception as e:
        print(f"Error in main execution: {str(e)}")
