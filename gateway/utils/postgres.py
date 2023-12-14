import psycopg2
from psycopg2 import OperationalError

def connect_to_postgres(database_url):
    try:
        conn = psycopg2.connect(database_url)
        return conn
    except OperationalError as e:
        print(f"Unable to connect to PostgreSQL. Error: {e}")
        return None

def ping_postgres(conn):
    try:
        with conn.cursor() as cursor:
            cursor.execute("SELECT 1")
            result = cursor.fetchone()
            if result and result[0] == 1:
                return True
    except OperationalError as e:
        print(f"Error pinging PostgreSQL: {e}")
    return False

database_url = "postgres://postgres:pass@journaling-app-postgres:5432/journaling-app-db"
postgres_conn = connect_to_postgres(database_url)

if postgres_conn:
    if ping_postgres(postgres_conn):
        print("PostgreSQL is accessible.")
    else:
        print("Failed to ping PostgreSQL.")
    postgres_conn.close()
else:
    print("Unable to connect to PostgreSQL.")
