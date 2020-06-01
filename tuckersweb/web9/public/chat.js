$(function () {
    if (!window.WebSocket) {
        alert("No WebSocket!");
        return;
    }
    var $chatlog = $('#chat-log'); //$( ) jq(JQuery)양식. 해당 ID에 대한 정보 가져옴.
    var $chatmsg = $('#chat-msg');

    addmessage = function (data) {
        $chatlog.prepend("<div><span>" + data + "</span></div>")
    }
    connect = function () {
        ws = new WebSocket("ws://" + window.location.host + "/ws");
        ws.onopen = function (e) {
            console.log("onopen", arguments);
        }
        ws.onclose = function (e) {
            console.log("onclose", arguments);
        }
        ws.onmessage = function (e) {
            addmessage(e.data);
        }
    }
    connect();


    //== 2개 값만 ===3개 타입까지 비교
    //=== 타입이 텍스트 + 빈문자열일 경우 트루 반환
    var isBlank = function (string) {
        return string == null || string.trim() === "";
    };

    var username;
    while (isBlank(username)) {
        username = prompt("What's your name?");
        if (!isBlank(username)) {
            $('#user-name').html('<b><h1>' + username + '</b>');
        }
    }

    //input-form을 누르면 동작하는 핸들러 등록
    $('#input-form').on('submit', function (e) {
        if (ws.readyState === ws.OPEN) {
            ws.send(JSON.stringify({
                type: "msg",
                data: $chatmsg.val()
            }));
        }
        //초기화 및 포커싱
        $chatmsg.val("");
        $chatmsg.focus();
        return false;//다른 페이지로 넘어가지 않으려면 false
    });

})