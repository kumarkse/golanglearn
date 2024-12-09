import pymysql
import os
from dotenv import load_dotenv

# Load environment variables from .env file
load_dotenv()

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

# Function to create interns table if it does not exist
def create_T1_categories_table(connection):
    try:
        with connection.cursor() as cursor:
            create_table_query = """
                CREATE TABLE IF NOT EXISTS T1_categories (
                            id INT AUTO_INCREMENT PRIMARY KEY,
                            topic_name VARCHAR(255) NOT NULL,
                            topic_url TEXT NOT NULL,
                            name VARCHAR(255),
                            profile_url TEXT,
                            image_url TEXT
                )
            """
            cursor.execute(create_table_query)
        connection.commit()
        print("Table T1_categories created or already exists.")
    except Exception as e:
        print(f"Error creating table: {str(e)}")

# Main function for local testing
if __name__ == "__main__":
    try:
        # Connect to the database
        connection = get_db_connection()

        # Create T1_categories table if not exists
        create_T1_categories_table(connection)

        # Close database connection
        connection.close()

    except Exception as e:
        print(f"Error in main execution: {str(e)}")
