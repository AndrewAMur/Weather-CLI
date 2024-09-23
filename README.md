# Weather CLI

A simple command-line tool written in Go that fetches current weather data for a specified city using the [WeatherAPI.com](https://www.weatherapi.com) service. This project demonstrates how to use external APIs, manage environment variables with `.env` files, and handle JSON data.

## Features

- Fetches the current weather for any city.
- Displays temperature, weather conditions, wind speed, and humidity.
- Uses a `.env` file to securely manage your API key.

## Prerequisites

Before running this project, ensure you have the following:

- Go installed on your system. [Install Go](https://golang.org/doc/install).
- A valid API key from [WeatherAPI.com](https://www.weatherapi.com).

## Installation

Follow these steps to get the project running on your local machine:

### 1. Clone the repository

```bash
git clone https://github.com/AndrewAMur/Weather-CLI.git
cd weathercli
```

### 2. Initialize Go modules

Inside the project directory, initialize Go modules:

```bash
go mod init weathercli
```

### 3. Install dependencies

Install the `godotenv` package for managing environment variables:

```bash
go get github.com/joho/godotenv
```

### 4. Set up the `.env` file

Create a `.env` file in the project root to store your API key. Replace `your-api-key-here` with your actual API key from WeatherAPI.com:

```
API_KEY=your-api-key-here
```

### 5. Run the program

Run the program by specifying the name of the city you'd like to get the weather for:

```bash
go run main.go <city-name>
```

Example:

```bash
go run main.go Toronto
```

This will display the current weather conditions for the specified city.

## Example Output

```bash
Weather in Toronto, Canada:
Temperature: 18.3°C
Condition: Partly cloudy
Wind: 12.5 kph
Humidity: 60%
```

## Project Structure

```
weathercli/
├── .env               # Environment variables (API key)
├── main.go            # Main Go program for fetching weather
├── go.mod             # Go module file
└── README.md          # This README file
```

## License

This project is licensed under the GNU License. See the `LICENSE` file for more details.
