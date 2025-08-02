#!/bin/bash

# Railway Deployment Script for TrackMe Go Application
# This script handles the deployment process to Railway

set -e

echo "🚀 Starting Railway deployment for TrackMe..."

# Check if Railway CLI is installed
if ! command -v railway &> /dev/null; then
    echo "❌ Railway CLI is not installed. Please install it first:"
    echo "   npm install -g @railway/cli"
    echo "   or visit: https://docs.railway.app/develop/cli"
    exit 1
fi

# Check if logged in to Railway
if ! railway whoami &> /dev/null; then
    echo "❌ Not logged in to Railway. Please login first:"
    echo "   railway login"
    exit 1
fi

# Ensure we're in the project root
if [ ! -f "go.mod" ]; then
    echo "❌ go.mod not found. Please run this script from the project root."
    exit 1
fi

echo "✅ Railway CLI is installed and authenticated"

# Build the application locally to check for errors
echo "🔨 Building application locally..."
go mod tidy
go build -o bin/trackme ./cmd/trackme

if [ $? -ne 0 ]; then
    echo "❌ Build failed. Please fix the errors before deploying."
    exit 1
fi

echo "✅ Local build successful"

# Deploy to Railway
echo "🚀 Deploying to Railway..."

# Check if Railway project exists and has a service
railway_status=$(railway status 2>&1)
if echo "$railway_status" | grep -q "No service could be found"; then
    echo "📝 No Railway service found. Creating new service..."
    railway service new
elif echo "$railway_status" | grep -q "No project linked"; then
    echo "📝 No Railway project found. Creating new project..."
    railway new
fi

# Check for required environment variables
echo "🔧 Checking environment variables..."
missing_vars=()

# Check if MONGODB_URI is set in Railway
if ! railway variables | grep -q "MONGODB_URI"; then
    missing_vars+=("MONGODB_URI")
fi

# Check if MONGODB_NAME is set in Railway
if ! railway variables | grep -q "MONGODB_NAME"; then
    missing_vars+=("MONGODB_NAME")
fi

if [ ${#missing_vars[@]} -gt 0 ]; then
    echo "⚠️  Missing required environment variables:"
    printf '   - %s\n' "${missing_vars[@]}"
    echo ""
    echo "Please set these in Railway dashboard or use railway variables:"
    echo "   railway variables set MONGODB_URI=your_mongodb_connection_string"
    echo "   railway variables set MONGODB_NAME=your_database_name"
    echo ""
    echo "Railway Dashboard: https://railway.app/dashboard"
    echo ""
    read -p "Continue with deployment? (y/N): " -n 1 -r
    echo
    if [[ ! $REPLY =~ ^[Yy]$ ]]; then
        echo "Deployment cancelled."
        exit 1
    fi
fi

# Deploy the application
railway up

if [ $? -eq 0 ]; then
    echo "✅ Deployment successful!"
    echo "🌐 Your application should be available at your Railway domain"
    echo "📊 Check deployment status: railway status"
    echo "📝 View logs: railway logs"
    echo "🔧 Open dashboard: railway open"
else
    echo "❌ Deployment failed. Check the logs for more details:"
    echo "   railway logs"
    exit 1
fi

echo "🎉 Railway deployment completed!"