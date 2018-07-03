
var Webservice = function (_Service) {
	this.Service = _Service;
	this.Init = function () {
		console.log('===正在连接服务器===');
		var self = this;
		if ('WebSocket' in window) {
			this.websocket = new WebSocket(this.Service);
		} else {
			alert('您的系统版本过低，请升级系统后运行')
		}

		//连接发生错误的回调方法
		this.websocket.onerror = function () {
			console.log("WebSocket连接发生错误");
		};

		//连接成功建立的回调方法
		this.websocket.onopen = function () {
			console.log("===WebSocket连接成功===");
			self.Open = true;
		}

		//接收到消息的回调方法
		this.websocket.onmessage = function (event) {
			var Data = {};
			
			try {
				var d = JSON.parse(event.data);
				
				if (d.action) {
					
					if(Config[d.action]&&(typeof Config[d.action] =="function")){
						
						Config[d.action](d);
					}
					else{
						console.log(d.action,"方法找不到")
					}
				}
				
			} catch (e) {
				console.error({ "执行错误": Data["Call"], "错误信息": e });
			}

		}

		//连接关闭的回调方法
		this.websocket.onclose = function () {
			self.Open = false;
			console.log("===WebSocket连接关闭===");
			self.Init();
		}
	};
	this.Send = function (message) {
		
		var self = this;
		if (this.Open && this.websocket.readyState === 1) {
			// console.log(JSON.stringify(message));
			this.websocket.send(JSON.stringify(message));
		} else {
			setTimeout(function () {
				self.Send(message);
			}, 1000);
		}
	};
	this.Close = function () {
		websocket.close();
	};
	this.Init();
}
window.webservice = new Webservice("ws://localhost:8082/WebSocket");
// window.webservice = new Webservice("ws://172.104.32.98:8082/WebSocket");