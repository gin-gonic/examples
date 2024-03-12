package main

import (
	"flag"
	"log"
	"text/template"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var addr = flag.String("addr", ":8080", "http service address")

var upgrader = websocket.Upgrader{} // use default option

func echo(ctx *gin.Context) {
	w, r := ctx.Writer, ctx.Request
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("upgrade:", err)
		return
	}
	defer c.Close()
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv:%s", message)
		err = c.WriteMessage(mt, message)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}

func home(c *gin.Context) {
	homeTemplate.Execute(c.Writer, "ws://"+c.Request.Host+"/echo")
}

func main() {
	flag.Parse()
	log.SetFlags(0)
	r := gin.Default()
	r.GET("/echo", echo)
	r.GET("/", home)
	log.Fatal(r.Run(*addr))
}

var homeTemplate = template.Must(template.New("").Parse(`
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="utf-8">
    <script>
        window.addEventListener("load", ()=> {
            let output = document.getElementById("output");
            let input = document.getElementById("input");
            let ws;
            let print = function (message) {
                let d = document.createElement("div");
                d.textContent = message;
                output.appendChild(d);
                output.scroll(0, output.scrollHeight);
            };
            document.getElementById("open").onclick = ()=> {
                if (ws) {
                    return;
                }
                ws = new WebSocket("{{.}}");
                ws.onopen = ()=> {
                    print("OPEN");
                }
                ws.onclose = ()=>{
                    print("CLOSE");
                    ws = null;
                }
                ws.onmessage = (evt)=> {
                    print("RESPONSE: " + evt.data);
                }
                ws.onerror = (evt)=> {
                    print("ERROR: " + evt.data);
                }
            };
            document.getElementById("send").onclick = ()=> {
                if (!ws) {
                    return;
                }
                print("SEND: " + input.value);
                ws.send(input.value);
				input.value = '';
            };
            document.getElementById("close").onclick = ()=> {
                if (!ws) {
                    return;
                }
                ws.close();
            };
        });
    </script>
    <title>WS</title>
</head>
<body>
<table>
    <tr><td valign="top" width="50%">
        <p>
            Click "Open" to create a connection to the server,
            "Send" to send a message to the server and "Close" to close the connection.
            You can change the message and send multiple times.
        </p>
        <button id="open">Open</button>
        <button id="close">Close</button>
        <p>
            <label for="input"></label>
            <input id="input" type="text" value="Hello world!"><button id="send">Send</button>
        </p>
    </td>
    <td valign="top" width="50%">
        <div id="output" style="max-height: 70vh;overflow-y: scroll;"></div>
    </td></tr>
</table>
</body>
</html>
`))
