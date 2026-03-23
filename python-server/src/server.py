import os
from http.server import HTTPServer, BaseHTTPRequestHandler


def response_message():
    feeling = os.getenv("FEELING", "$FEELING")
    awesomness = os.getenv("AWESOMNESS", "$AWESOMNESS")
    return f'Howdy! I feeeeel {feeling} today. I will tell you my secret of awesomeness: "{awesomness}".'

class HelloHandler(BaseHTTPRequestHandler):
    def do_GET(self):
        response = response_message()
        
        self.send_response(200)
        self.send_header("Content-Type", "text/plain")
        self.end_headers()
        self.wfile.write(response.encode())
    
    def log_message(self, format, *args):
        pass  # Suppress default logging

if __name__ == "__main__":
    port = int(os.getenv("PORT", 8080))
    server_address = ("0.0.0.0", port)
    httpd = HTTPServer(server_address, HelloHandler)
    
    print(f"Server running on http://0.0.0.0:{port}")
    httpd.serve_forever()
