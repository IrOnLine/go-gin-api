FROM mariadb:latest

# Set environment variables for MariaDB
ENV MYSQL_ROOT_PASSWORD=mysecretpassword
ENV MYSQL_DATABASE=mydatabase
ENV MYSQL_USER=myuser
ENV MYSQL_PASSWORD=mypassword

# Copy custom configuration file to the container
# COPY my.cnf /etc/mysql/my.cnf

# volume $PWD/my_container_config:/etc/mysql/conf.d:z
# VOLUME ./db/init.sql:/tmp/init.sql
VOLUME /mysql/:/var/lib/mysql/

# Expose the default MariaDB port
EXPOSE 3306

# Start the MariaDB service when the container launches
CMD ["mysqld"]