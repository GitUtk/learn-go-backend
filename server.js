const http = require("http");
const querystring = require("querystring");

const server = http.createServer((req, res) => {
  if (req.method === "GET" && req.url === "/get") {
    res.end("GET endpoint");
    return;
  }

  if (req.method === "POST" && req.url === "/post") {
    let body = "";

    req.on("data", (chunk) => {
      body += chunk;
    });

    req.on("end", () => {
      console.log("Received data:");
      console.log(body);

      res.end("Received Data : " + body);
    });

    return;
  }

  // Handle Go http.PostForm request
  if (req.method === "POST" && req.url === "/postform") {
    let body = "";

    req.on("data", (chunk) => {
      body += chunk;
    });

    req.on("end", () => {
      const formData = querystring.parse(body);

      console.log("Received form data:");
      console.log(formData);

      res.writeHead(200, {
        "Content-Type": "application/json",
      });

      res.end(
        JSON.stringify({
          message: "Form data received successfully",
          data: formData,
        })
      );
    });

    return;
  }

  res.statusCode = 404;
  res.end("Not Found");
});

server.listen(3000, () => {
  console.log("Server running on http://localhost:3000");
});