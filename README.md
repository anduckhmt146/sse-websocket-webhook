# Sever-side Event (SSE), Webhook, WebSocket in Golang

In backend server-side, you can implement a Kafka consumer (backend service) to consume message and processing server communication real-time.

But of course, in client-side service, we are not recommend to consume directly to Kafka. If you still need to real time update from server-side to client-side, you can use SSE pattern (with HTTP Flusher).

## When to use SSE, when to use Webhook, when to use WebSocket

- **Webhook:** 

    - **Usage:** When merchant to make a HTTP request (POST, GET) to your server to confirm their events have completed. It is called webhook.

    - **Example:** You are building an admin dashboard for an e-commerce platform. You want real-time order updates. Customer buys a product → Payment gateway (Stripe/PayPal) triggers a Webhook to your server. Your server processes the webhook, confirms order, updates database. 

- **SSE:**

    - **Usage:** When to need to use a real-time update in client-side without polling.

    - **Example:** You have a social media platform where users need to be notified of new likes, comments, or mentions in real-time. When another user interacts with their content, your server sends an SSE event to update their notification feed instantly, without requiring page refresh. This is more efficient than polling the server every few seconds for updates.

- **WebSocket:**

    - **Usage:** When you need bi-directional real-time communication between client and server. Both server and client can send messages to each other at any time.

    - **Example:** Chat applications, real-time gaming, collaborative editing where both client and server need to push updates to each other frequently.

## How to run

### Start server

```bash
go run main.go
```

### Start client

You can using feature `Start live server` or `Open this file html to browser` to execute.

## Result

### Curl SSE

```bash
curl http://localhost:8080/sse/send
```

### Curl WebSocket

```bash
curl http://localhost:8080/ws/send
```

### Client SSE

```javascript
 const evtSource = new EventSource("http://localhost:8080/sse");

    evtSource.onmessage = function(event) {
        console.log("📢 New Notification:", event.data);
        const p = document.createElement('p');
        p.innerText = event.data;
        document.body.appendChild(p);
    };
```

![](/assets/sse.png)

### Client WebSocket

```javascript
     const socket = new WebSocket("ws://localhost:8080/ws");

      socket.onopen = function() {
        console.log("Connected to WebSocket server");
      };

      socket.onmessage = function(event) {
        console.log("📢 WebSocket message:", event.data);
        const p = document.createElement('p');
        p.innerText = event.data;
        document.body.appendChild(p);
      };

      socket.onclose = function() {
        console.log("Disconnected from WebSocket server");
      };
```

![](/assets/ws.png)