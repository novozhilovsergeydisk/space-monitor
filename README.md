# Space Monitor

This is a Go application that monitors disk space usage on a Linux server and sends alerts via Telegram and email when usage exceeds 85%.

## Features

- Disk space monitoring
- Telegram notifications
- Email alerts
- Logging to file and database
- Systemd service for automatic startup
- Deployment automation script

## Installation

1. Clone the repository
2. Set up environment variables
3. Run the deployment script

```bash
./scripts/deploy.sh
```

## Configuration

Set the following environment variables:

- `TELEGRAM_BOT_TOKEN`: Your Telegram bot token
- `EMAIL_FROM`: Email address to send from
- `EMAIL_PASSWORD`: Email account password
- `EMAIL_TO`: Email address to send to

## Usage

The service will start automatically after installation. You can check its status with:

```bash
sudo systemctl status space-monitor
```
