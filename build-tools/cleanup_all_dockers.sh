#!/bin/bash


echo "🔍 Checking for running Docker containers..."
running_containers=$(docker ps -q)

if [ -z "$running_containers" ]; then
    echo "✅ No running containers found. Nothing to stop or remove."
    exit 0
fi

echo "🛑 Stopping all running Docker containers..."
docker stop $(docker ps -q)

echo "🗑️ Removing all stopped containers..."
docker rm $(docker ps -aq)

echo "✅ Done! All containers have been stopped and removed."

echo "🧹 Removing all Docker volumes..."
docker volume prune -f

echo "🔌 Removing all unused Docker networks..."
docker network prune -f
