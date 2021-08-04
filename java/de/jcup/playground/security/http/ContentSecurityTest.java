package de.jcup.playground.security.http;

import java.io.IOException;
import java.io.OutputStream;
import java.net.InetSocketAddress;
import java.util.Arrays;

import com.sun.net.httpserver.HttpExchange;
import com.sun.net.httpserver.HttpHandler;
import com.sun.net.httpserver.HttpServer;


public class ContentSecurityTest {

    public static void main(String[] args) throws Exception {
        HttpServer server = HttpServer.create(new InetSocketAddress(8000), 0);
        server.createContext("/test", new SimpleResponseHandler());
        server.setExecutor(null); // creates a default executor
        server.start();
    }

    static class SimpleResponseHandler implements HttpHandler {
        @Override
        public void handle(HttpExchange t) throws IOException {
            /* @formatter:off */
            String response = 
                    "<html>"+
                         "<style>.myStyle { color: red }</style>"+
                         "<body class='myStyle'>Hello world...</body>"+
                    "</html";
            /* @formatter:on */
            
            t.getResponseHeaders().put("Content-Security-Policy", Arrays.asList("default-src localhost"));
            // when next line is NOT commented, the the "Hello World will be red... other wise not because of CSP set
//            t.getResponseHeaders().put("Content-Security-Policy", Arrays.asList("style-src 'unsafe-inline'"));

            OutputStream os = t.getResponseBody();
            t.sendResponseHeaders(200, response.length());
            
            os.write(response.getBytes());
            os.close();
        }
    }

}
