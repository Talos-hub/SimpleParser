#  Simple Concurrent Web Parser in Go

A lightweight, concurrent web parser built in Go that extracts data from multiple URLs using regular expressions and saves the results to a file. Perfect for learning core Go concepts!

##  Features

-  **Concurrent Processing**: Fetches and parses multiple URLs simultaneously
-  **Robust Error Handling**: Comprehensive error checking with specific case handling
-  **Configurable Pattern Matching**: Regex-based content extraction
-  **Thread-Safe Operations**: Mutex-protected file writing
-  **Structured Logging**: JSON logging to file/stdout with configurable levels
-  **Proper Timeouts**: HTTP client with timeout settings

## 
Getting Started

### Prerequisites

- Go 1.18+

### Installation

1. Clone the repository:
```bash
git clone https://github.com/your-username/simple-go-parser.git
cd simple-go-parser
