#!/bin/bash

# Function to display usage
usage() {
    echo "Usage: $0 [option]"
    echo "Options:"
    echo "  1  Run backend only"
    echo "  2  Run frontend only"
    echo "  3  Run both backend and frontend"
    exit 1
}

# Check if an argument is provided
if [ -z "$1" ]; then
    usage
fi

# Start the backend server if option is 1 or 3
if [ "$1" == "1" ] || [ "$1" == "3" ]; then
    echo "Starting backend server..."
    cd api
    go run cmd/api/main.go &
    BACKEND_PID=$!
    cd - 
fi

# Start the frontend server if option is 2 or 3
if [ "$1" == "2" ] || [ "$1" == "3" ]; then
    echo "Starting frontend server..."
    cd frontend/my-app
    npm run dev &
    FRONTEND_PID=$!
    cd -
fi

# Wait for both processes to finish if both are started
if [ "$1" == "3" ]; then
    wait $BACKEND_PID $FRONTEND_PID
elif [ "$1" == "1" ]; then
    wait $BACKEND_PID
elif [ "$1" == "2" ]; then
    wait $FRONTEND_PID
fi
