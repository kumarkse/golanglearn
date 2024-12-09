import pymysql
import os
import uuid
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

# Function to insert a record into the interns table
def insert_record(connection, topic_name, topic_url, name, profile_url, image_url):
    try:
        with connection.cursor() as cursor:
            insert_query = """
                INSERT INTO T1_categories(topic_name, topic_url, name, profile_url, image_url)
                VALUES (%s, %s, %s, %s, %s)
            """
            cursor.execute(insert_query, (topic_name, topic_url, name, profile_url, image_url))
        connection.commit()
        print(f"Record inserted successfully for topic: {topic_name}")
    except Exception as e:
        print(f"Error inserting record: {str(e)}")

# Function to read and parse the data from the .txt file
def read_scraped_data(file_path):
    data = []
    with open(file_path, 'r') as file:
        lines = file.readlines()

    topic_name, topic_url = None, None
    name, profile_url, image_url = None, None, None

    for line in lines:
        if line.startswith("Topic Name:"):
            topic_name = line.split(": ", 1)[1].strip()
        elif line.startswith("Topic URL:"):
            topic_url = line.split(": ", 1)[1].strip() if ": " in line else None
            # When a new topic starts, save the previous topic if it exists
            if topic_name and topic_url:
                data.append((topic_name, topic_url, name, profile_url, image_url))
            # Reset poster details for the new topic
            poster_name, poster_profile_url, poster_image_url = None, None, None
        elif line.startswith("  Name:"):
            name = line.split(": ", 1)[1].strip()
        elif line.startswith("  Profile URL:"):
            profile_url = line.split(": ", 1)[1].strip() if ": " in line else None
        elif line.startswith("  Image URL:"):
            image_url = line.split(": ", 1)[1].strip() if ": " in line else None

    # Append the last topic
    if topic_name and topic_url:
        data.append((topic_name, topic_url, name, profile_url, image_url))

    return data


# Function to process all .txt files in a specified directory
def process_all_files(directory):
    try:
        # Connect to the database
        connection = get_db_connection()

        # List all .txt files in the directory
        for filename in os.listdir(directory):
            if filename.endswith(".txt"):
                file_path = os.path.join(directory, filename)
                print(f"Processing file: {file_path}")

                # Read scraped data from the .txt file
                scraped_data = read_scraped_data(file_path)

                # Insert the scraped data into the topics_posters table
                for record in scraped_data:
                    insert_record(connection, *record)

        # Close database connection
        connection.close()

    except Exception as e:
        print(f"Error processing files: {str(e)}")

# Main function for local testing
if __name__ == "__main__":
    # Directory containing all .txt files
    directory = "c:\\Users\\Vaishnavi M\\OneDrive\\Desktop\\Prodigal_AI\\Task2(upd)\\all_categories" 
    # Process all .txt files in the specified directory
    process_all_files(directory)
