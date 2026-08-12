const http = require("http");

const server = http.createServer((req, res) => {
  if (req.method === "GET" && req.url === "/get") {
    res.end("GET endpoint");
  }

  if (req.method === "POST" && req.url === "/post") {
    res.end("POST endpoint");
  }
});

server.listen(3000, () => {
  console.log("Server running on http://localhost:3000");
});