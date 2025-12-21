# LinkedIn Automation Tool

## Features
- Authentication System
- Search & Targeting
- Connection Requests
- Messaging System

### Authentication System
<input type="checkbox"> Login with LinkedIn credentials using environment variables. <br>
<input type="checkbox"> detect and handle login challenges (e.g., CAPTCHA, 2FA).<br>
<input type="checkbox"> Identify security checkpoints (2FA, captcha).<br>
<input type="checkbox"> Persist session cookies for seamless re-authentication.<br>

### Search & Targeting
<input type="checkbox"> Search for users based on job title,company,keywords, locations. <br>
<input type="checkbox"> Parse and collect user profile URLs efficiently.<br>
<input type="checkbox"> Handle pagination accross search results.<br>
<input type="checkbox"> Implement duplicate profile detection and filtering.<br>

### Connection Requests
<input type="checkbox"> Navigate to user profiles programmatically.<br>
<input type="checkbox"> Click Connect button with precise targeting.<br>
<input type="checkbox"> Send personalized notes within character limits.<br>
<input type="checkbox"> Track sent requests and enforce daily limits.<br>

### Messaging System
<input type="checkbox"> Detect newly accepted connections<br>
<input type="checkbox"> Send follow-up messages automatically<br>
<input type="checkbox"> Support templates with dynamic variables<br>
<input type="checkbox"> Maintain comprehensive message tracking

## Setup

### Prerequisites
- [Go](https://go.dev/dl/) installed on your machine.
- A LinkedIn account.

### Installation

1.  **Clone the repository**:
    ```bash
    git clone <repository-url>
    cd linkedin-automation
    ```

2.  **Install dependencies**:
    ```bash
    go mod tidy
    ```

3.  **Configure Environment Variables**:
    Create a `.env` file in the root directory and add your LinkedIn credentials:
    ```env
    username=your_email@example.com
    password=your_password
    ```

4.  **Run the Bot**:
    ```bash
    go run cmd/app/main.go
    ```
    Or build and run:
    ```bash
    go build -o linkedin-bot cmd/app/main.go
    ./linkedin-bot
    ```