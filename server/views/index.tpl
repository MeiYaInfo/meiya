<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Document</title>
    <script src="http://cdn.bootcss.com/jquery/3.1.1/jquery.min.js"></script>
</head>
<body>
    <!-- <img src="favicon.ico" alt="test"> -->
    <form>
        <div>
            <label for="type">类型：</label>
            <select id="type">
                <option value="ninepic">九图</option>
                <option value="wildpic">野表情</option>
                <option value="petpic">萌宠搞事</option>
            </select>
        </div>
        <div>
            <label for="pid">pid：</label>
            <input id="pid" type="text" value="a15b4afegw1f938pm4oo6j20ki03kmxr">
        </div>
        <div>
            <label for="startTime">startTime：</label>
            <input id="startTime" type="number" value="1484469117">
        </div>
        <div>
            <label for="endTime">endTime：</label>
            <input id="endTime" type="number" value="1485469117">
        </div>
        <button id="add">提交</button>
        <div>提交的数据为：<span id="postData"></span></div>
        <div>返回的数据为：<span id="response"></div></div>
    </form>
    <script>
        $('#add').click(function (e) {
            e.preventDefault();
            var postData = {
                type: $('#type').val(),
                pid: $('#pid').val(),
                startTime: parseInt($('#startTime').val()),
                endTime: parseInt($('#endTime').val())
            };
            $('#postData').append(JSON.stringify(postData));
            $.post('http://192.168.31.97:8080/editAd', postData, function (res) {
                $('#response').append(JSON.stringify(res));
            })

        })
    </script>
</body>
</html>