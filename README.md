# 📦🎉 Made In My Room - Party Game!
A real-time online multiplayer browser party game inspired by [Jackbox.TV](https://jackbox.tv/), and entirely *made in my room*!!

# ⚙️ Prerequisites
- Go 1.25.3  
[Install Golang Here!](https://go.dev/doc/install)

- Node.js 24.9.0  
[Install Node.js Here!](https://nodejs.org/en/download)

# 💻 Setup & Usage
Follow these steps to get your development environment set up and operational:  
1. **Clone the Repository**
    ```bash
    git clone https://github.com/dotping-me/made-in-my-room.git
    cd made-in-my-room/
    ```
2. **Install Go Dependencies**
    ```bash
    cd backend/
    go mod tidy
    ```
3. **Build & Run Go Binaries**
    ```bash
    go build -o ./main
    ./main
    ```

---
***Note:*** *Follow the steps below on another terminal!*

4. **Install Node.js Dependencies**
    ```bash
    cd frontend/
    npm install
    ```

5. **Start Frontend**
    ```bash
    npm run dev
    ```