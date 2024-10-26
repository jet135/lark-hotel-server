chrome.runtime.onInstalled.addListener(() => {
  chrome.declarativeNetRequest.updateDynamicRules({
    addRules: [{
      "id": 1,
      "priority": 1,
      "action": {
        "type": "modifyHeaders",
        "requestHeaders": [
          { "header": "x-custom-header", "operation": "set", "value": "value" }
        ]
      },
      "condition": {
        "urlFilter": "*",
        "resourceTypes": ["main_frame", "sub_frame"]
      }
    }],
    removeRuleIds: [1]
  });
});

chrome.webRequest.onBeforeRequest.addListener(
  function(details) {
    console.log('Original Request:', details);

    let requestData = {};

    if (details.requestBody) {
      if (details.requestBody.raw) {
        // 处理 raw 类型的请求体
        let decoder = new TextDecoder("utf-8");
        details.requestBody.raw.forEach((element) => {
          let str = decoder.decode(element.bytes);
          Object.assign(requestData, JSON.parse(str));
        });
      }
    }

    // 构造你的新请求
    let newRequest = {
      url: "http://localhost:8808/gdzwfw/report",
      method: "POST",
      headers: {
        "Content-Type": "application/json"
      },
      body: JSON.stringify({
        originalUrl: details.url,
        method: details.method,
        requestBody: requestData
      })
    };

    // 发送新请求
    fetch(newRequest.url, {
      method: newRequest.method,
      headers: newRequest.headers,
      body: newRequest.body
    })
    .then(response => response.json())
    .then(data => console.log('New Request Response:', data))
    .catch(error => console.error('Error:', error));
  },
  // todo 你要监听的接口地址
  { urls: [""] },
  ["requestBody"]
);