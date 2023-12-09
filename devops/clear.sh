#!/bin/bash

# Find the process ID (PID) of the application running on port 8080
PID=$(lsof -t -i:8080)

# Check if a PID was found
if [ -z "$PID" ]; then
    echo "No application is running on port 8080."
else
    # Kill the process
    kill $PID
    if [ $? -eq 0 ]; then
        echo "Application on port 8080 has been terminated."
    else
        echo "Failed to terminate the application on port 8080. You might need to run this script as root."
    fi
fi
