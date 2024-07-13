# Goffer - Telegram Bot

'Goffer' is a Go-based Telegram bot that interacts with users by playing games (Rock-Paper-Scissors) and managing highscores using Redis. This project serves as an example of how to build a simple yet functional Telegram bot with Go and Redis.

## Table of Contents

- [Introduction](#introduction)
- [Features](#features)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Configuration](#configuration)
- [Usage](#usage)
- [Project Structure](#project-structure)
- [Contributing](#contributing)
- [License](#license)

## Introduction

Goffer is designed to be a fun and interactive Telegram bot. Users can start a game with the bot, and the bot will keep track of highscores using Redis. The bot is implemented in Go and leverages the Telegram Bot API.

## Features

- Start a game with the bot using `/start`
- Play a simple game and compete for highscores
- End the game with `/bye`
- Redis integration for storing and retrieving highscores

## Prerequisites

Before you begin, ensure you have the following installed:

- Go (version 1.18+)
- Redis server
- Git

## Installation

1. **Clone the repository:**

```bash
git clone https://github.com/Eastwesser/Goffer.git
cd Goffer
```

Install dependencies:

```bash
go mod tidy
```

Install Redis:

```bash
sudo apt update
sudo apt install redis-server
```

Start Redis server:

```bash
sudo service redis-server start
redis-cli ping
```

You should see PONG if Redis is running properly.

Configuration
Set up environment variables:

Create a .env file in the root directory of the project and add your Telegram bot token:

```plaintext
BOT_TOKEN=your_telegram_bot_token
```

## Usage

Run the bot:

```bash
go run cmd/main.go
```

### Interact with the bot:

Start a game: ```/start```

End the game: ```/bye```

Project Structure:
```bash
Goffer/
├── cmd/
│   └── main.go         # Main entry point of the application
├── internal/
│   ├── bot/
│   │   ├── mainloop.go # Main loop to handle Telegram updates
│   │   ├── helpers.go  # Helper functions for bot operations
│   │   ├── handlers.go # Handlers for bot actions and commands
│   │   └── commands.go # Command handlers
│   ├── game/
│   │   └── game.go     # Game logic
│   └── redis/
│       └── redis.go    # Redis client setup and highscore management
├── .env                # Environment variables
├── go.mod              # Go module file
└── README.md           # Project documentation
```

## Contributing

Contributions are welcome! Please open an issue or submit a pull request if you would like to contribute to this project.

Feel free to adjust any section based on specific details or preferences for your project.