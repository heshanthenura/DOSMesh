<img src="banner.png">

## **DoS Mesh**

**DoS Mesh** is a lightweight distributed denial-of-service (DDoS) simulation system built in Go. It includes:

- A **Server** to control multiple connected clients (bots)
- **Client Bots** that receive commands and execute ICMP flood attacks

> ⚠️ **Disclaimer:** This project is strictly for **educational** and **research** purposes **only**. **Do not** use this software on networks or devices without **explicit permission**. Misuse is illegal and unethical.

---

### ✅ Prerequisites

- [Go](https://golang.org/dl/) must be installed (version **1.18** or higher recommended)

---

### 🚀 How to Run and Use

#### 🖥️ Server

1. Open `server.go` located at:

   [`server/cmd/server/server.go`](https://github.com/heshanthenura/DOSMesh/blob/main/server/cmd/server/server.go)

2. Change the port number in the following line if needed:

   ```go
   socketserver.SocketServer(8080) // Replace with your desired port
   ```

3. Run the server:

   ```bash
   go run server.go
   ```

4. You can now type commands into the server terminal:

   - `start` – Start the ICMP flood attack on all connected bots
   - `stop` – Stop the flood attack
   - `c` – Show the number of connected bots

#### 🤖 Client Bot

1. Open `bot.go` located at:

   [`bot/cmd/bot/bot.go`](https://github.com/heshanthenura/DOSMesh/blob/main/bot/cmd/bot/bot.go)

2. Set the target IP address for the attack by modifying this line:

   ```go
   go socket.RunFlood(controlChan, "192.168.1.101", 0*time.Millisecond) // Replace IP with your target
   ```

3. Build and run the bot (use `sudo` to allow raw ICMP packet sending):

   ```bash
   go run bot.go
   ```

   or

   ```bash
   ./bot
   ```

4. Deploy bots only on machines you own or have explicit permission to use for testing purposes.

---

### 📊 Bot Traffic Behavior

Each bot in the **DoS Mesh** network is designed to simulate lightweight, realistic DDoS behavior. By default:

- 📶 **Each bot generates approximately 500 Kbps** (kilobits per second) of ICMP traffic
- ⚙️ The flood behavior is controlled by a delay and packet size defined in `flood.SendICMPFlood()`
- 🔧 You can adjust the flood intensity by modifying:

  - The `sleepTime` duration between packets
  - The size of the `Data` payload in the ICMP packet

This helps simulate a distributed attack that remains under the radar and avoids drawing immediate attention during tests.

---

### 🚧 Project Status & Future Improvements

> ❗ **Note:** This project is still under development and not production-ready.

While **DoS Mesh** currently supports basic command-and-control for ICMP flood simulations, there are several areas that need improvement:

- 🔒 **Error handling:** Current error reporting is minimal. More robust logging and recovery mechanisms are needed.
- 🌐 **Dynamic target updates:** Bots currently require manual changes to set the target. A command to change targets remotely is planned.
- 🧠 **Smarter flood logic:** Improve flood behavior to better simulate realistic attack patterns.
- 📦 **Binary distribution:** Build scripts and proper cross-platform support.

Contributions and suggestions are welcome!
