# Use an official Python runtime as a base image
FROM python:3.11.7-slim

# Set the working directory in the container to /app
WORKDIR /app

# Copy the current directory contents into the container at /app
COPY . /app

# Install Deps using Pipenv
RUN pip install pipenv
RUN cd /app && pipenv install --system --deploy --ignore-pipfile

# Install Gunicorn Webserver
RUN pip install gunicorn

# Install gevent for pseudo-threads
RUN pip install gevent

# Make port 5000 available to the world outside this container
EXPOSE 5000

# Define environment variables
ENV FLASK_APP=app.py
ENV FLASK_RUN_HOST=0.0.0.0
ENV FLASK_ENV=production

# Run the app using Gunicorn
## Recommended for number of workers to be (2 * num_cores) + 1
### Source: https://medium.com/building-the-system/gunicorn-3-means-of-concurrency-efbb547674b7
## Timeout set to 180 seconds
CMD ["pipenv", "run", "gunicorn", "--timeout", "180", "-b", "0.0.0.0:5000", "app:app", "--worker-class=gevent", "--worker-connections=1000", "--workers=5"]