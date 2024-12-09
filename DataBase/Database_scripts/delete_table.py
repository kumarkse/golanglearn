import pymysql
import os
from dotenv import load_dotenv
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

# Function to delete the interns table if it exists
def delete_table_if_exists(connection):
    try:
        with connection.cursor() as cursor:
            delete_table_query = "DROP TABLE IF EXISTS interns"
            cursor.execute(delete_table_query)
        connection.commit()
        print("Table 'interns' deleted successfully if it existed.")
    except Exception as e:
        print(f"Error deleting table: {str(e)}")

# Main function for local testing
if __name__ == "__main__":
    try:
        # Connect to the database
        connection = get_db_connection()

        # Delete the interns table if it exists
        delete_table_if_exists(connection)

        # Close database connection
        connection.close()

    except Exception as e:
        print(f"Error in main execution: {str(e)}")
