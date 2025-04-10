# Weather CLI

A command-line weather application that demonstrates working with external APIs, JSON parsing, and error handling in Go.

## Concepts Covered

- HTTP requests
- JSON parsing
- Command-line flags
- Environment variables
- Error handling
- API integration
- Struct tags
- Formatting output

## Features

- Get current weather for any location
- Display temperature in both Celsius and Fahrenheit
- Show weather conditions
- Display additional weather information:
  - Humidity
  - Wind speed
  - UV index
  - "Feels like" temperature
- Support for API key via flag or environment variable

## Prerequisites

1. Go installed on your system
2. API key from [WeatherAPI](https://www.weatherapi.com/)
   - Sign up for a free account
   - Get your API key from the dashboard

## How to Run

1. Navigate to the project directory:
   ```bash
   cd level-2/03-weather-cli
   ```

2. Set your API key (either method):
   ```bash
   # Method 1: Environment variable
   export WEATHER_API_KEY=your_api_key_here

   # Method 2: Command-line flag
   go run main.go -key=your_api_key_here -location="London"
   ```

3. Run the program:
   ```bash
   go run main.go -location="London"
   ```

## Usage Example

```bash
$ go run main.go -location="New York"

Weather for New York, United States
----------------------------------------
Temperature: 22.5°C (72.5°F)
Feels like: 23.2°C (73.8°F)
Condition: Partly cloudy
Humidity: 65%
Wind Speed: 12.6 km/h
UV Index: 5.0
```

## Project Structure

```
03-weather-cli/
├── main.go      # Main program file
└── README.md    # Project documentation
