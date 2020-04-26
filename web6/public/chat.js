$(function(){
    if(!window.EventSource){ //윈도우에서 이 EventSource 지원 확인
        alert("No EventSource!");
        return;
    }

    var $chatlog = $('#chat-log'); //$( ) jq(JQuery)양식. 해당 ID에 대한 정보 가져옴.
    var $chatmsg = $('#chat-msg');

    //== 2개 값만 ===3개 타입까지 비교
    //=== 타입이 텍스트 + 빈문자열일 경우 트루 반환
    var isBlank = function(string){
        return string == null || string.trim() === "";
    };
    
    var username;
    while(isBlank(username)){
        username = prompt("What's your name?");
        if( !isBlank(username)){
            $('#user-name').html('<b><h1>' + username + '</b>');
        }
    }

    //input-form을 누르면 동작하는 핸들러 등록
    $('#input-form').on('submit', function(e){
        //jq post를 쓰겠다.
        $.post('/messages', {
            msg: $chatmsg.val(),
            name: username
        });
        //초기화 및 포커싱
        $chatmsg.val("");
        $chatmsg.focus();

        return false;//다른 페이지로 넘어가지 않으려면 false
    });

    //메세지 화면에 추가
    var addMessage = function(data) {
        var text = "";
        if(!isBlank(data.name)){
            text += '<strong>' + data.name + ':</strong> ';
        }
        text += data.msg;
        $chatlog.prepend('<div><span>' + text + '</span></div>');     
    };

    //스트림 경로 요청..
    var es = new EventSource('/eventsource');
    //es 오픈 핸들러 등록
    es.onopen = function(e) {
        $.post('users/', {
            name: username
        });
    };
    //es 메시지 핸들러 등록
    es.onmessage = function(e) {
        var msg = JSON.parse(e.data);
        addMessage(msg);
    };

    window.onbeforeunload = function() {
        $.ajax({
            url:"/users?username=" + username,
            type:"DELETE"
        });
        es.close();
    };
})