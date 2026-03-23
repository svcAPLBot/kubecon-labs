use hyper::service::{make_service_fn, service_fn};
use hyper::{Body, Request, Response, Server};
use std::convert::Infallible;
use std::env;
use std::net::SocketAddr;

fn response_message() -> String {
    let feeling = env::var("FEELING").unwrap_or_else(|_| "$FEELING".to_string());
    let awesomness = env::var("AWESOMNESS").unwrap_or_else(|_| "$AWESOMNESS".to_string());

    format!(
        "Howdy! I feeeeel {} today. I will tell you my secret of awesomeness: \"{}\".",
        feeling, awesomness
    )
}

async fn handle(_req: Request<Body>) -> Result<Response<Body>, Infallible> {
    let response = response_message();
    Ok(Response::new(Body::from(response)))
}

#[cfg(test)]
mod tests {
    use super::response_message;

    #[test]
    fn response_uses_env_values() {
        unsafe {
            std::env::set_var("FEELING", "curious");
            std::env::set_var("AWESOMNESS", "snacks");
        }

        assert_eq!(
            response_message(),
            "Howdy! I feeeeel curious today. I will tell you my secret of awesomeness: \"snacks\"."
        );
    }

    #[test]
    fn response_defaults_to_placeholders() {
        unsafe {
            std::env::remove_var("FEELING");
            std::env::remove_var("AWESOMNESS");
        }

        assert_eq!(
            response_message(),
            "Howdy! I feeeeel $FEELING today. I will tell you my secret of awesomeness: \"$AWESOMNESS\"."
        );
    }
}

#[tokio::main]
async fn main() {
    let port = env::var("PORT").unwrap_or_else(|_| "8080".to_string());
    let addr: SocketAddr = format!("0.0.0.0:{}", port).parse().unwrap();

    let make_svc = make_service_fn(|_conn| async { Ok::<_, Infallible>(service_fn(handle)) });

    let server = Server::bind(&addr).serve(make_svc);

    println!("Server running on http://{}", addr);

    if let Err(e) = server.await {
        eprintln!("server error: {}", e);
    }
}
