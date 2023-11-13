#!/bin/bash

# Define the host and port
HOST="localhost"
PORT=4000  # Replace with your specific port

MESSAGE="Sending some other message!\n"
ITERATIONS=10000

# Loop to send the message multiple times
for ((i = 1; i <= ITERATIONS; i++)); do
    printf "$MESSAGE" | nc $HOST $PORT
done