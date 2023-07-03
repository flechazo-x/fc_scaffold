// Package htmlx
// @description
// @author      张盛钢
// @datetime    2023/5/8 18:27
package htmlx

const HtmlStr = `<!DOCTYPE html>
<html>

<head>
    <meta charset="UTF-8">
    <title>fc_scaffold</title>
    <style>
        body {
            background: #f3f3f3;
            margin: 0;
            padding: 0;
            font-family: Arial, sans-serif;
        }

        .welcome {
            margin-top: 50px;
            text-align: center;
        }

        .progress {
            width: 0%;
            height: 5px;
            background-color: #ef5350;
            /* 进度条颜色 */
            animation: progress 2s linear forwards;
        }

        .title {
            font-family: 'Montserrat', sans-serif;
            /* 使用Google Fonts中的Montserrat字体 */
            font-size: 3rem;
            /* 字体大小 */
            color: #4e4e4e;
            /* 字体颜色 */
            animation: show 4s ease-out forwards;
            opacity: 0;
        }

        @keyframes progress {
            from {
                width: 0%;
            }

            to {
                width: 100%;
            }
        }

        @keyframes show {
            0% {
                transform: translateX(-200%);
                opacity: 0;
            }

            30% {
                transform: translateX(20%);
                opacity: 1;
            }

            70% {
                transform: translateX(-10%);
            }

            100% {
                transform: translateX(0);
            }
        }


        .container {
            display: flex;
            flex-wrap: wrap;
            justify-content: space-between;
            align-items: center;
            max-width: 960px;
            margin: 0 auto;
            margin-top: 50px;
            padding: 20px;
            background: #fff;
        }

        .input-field {
            flex-basis: calc(33.33% - 10px);
            margin-bottom: 20px;
            position: relative;
        }

        .input-field input {
            width: 100%;
            padding: 10px;
            border: none;
            border-bottom: 1px solid #ddd;
            background: transparent;
            font-size: 16px;
            color: #333;
            outline: none;
            box-sizing: border-box;
        }

        .input-field input:focus,
        .input-field input:active {
            border-bottom-color: #5e98e9;
        }

        .label {
            position: absolute;
            left: 10px;
            top: -10px;
            font-size: 14px;
            color: #130202;
            background: #fff;
            padding: 0 5px;
        }

        .btn {
            display: inline-block;
            padding: 10px 20px;
            background: #5e98e9;
            color: #fff;
            border: none;
            border-radius: 3px;
            font-size: 16px;
            cursor: pointer;
            transition: background 0.3s ease-in-out;
        }

        .btn:hover {
            background: #4c7bd9;
        }
    </style>
</head>

<body>
<div class="welcome">
    <div class="progress"></div>
    <p>Welcome to use fc scaffolding tool.</p>
</div>
<div class="container">
    <div class="input-field">
        <label for="PkgName">PackageName</label>
        <input type="text" id="PkgName" name="PkgName" placeholder="请输入活动包名">
    </div>
    <div class="input-field">
        <label for="Id">ActivityID</label>
        <input type="text" id="Id" name="Id" placeholder="请输入活动ID">
    </div>
    <div class="input-field">
        <label for="StructName">StructName</label>
        <input type="text" id="StructName" name="StructName" placeholder="请输入结构体名">
    </div>
    <div class="input-field">
        <label for="Author">AuthorName</label>
        <input type="text" id="Author" name="Author" placeholder="请输入作者名字">
    </div>
    <div class="input-field">
        <label for="OutputPath">OutputPath</label>
        <input type="text" id="OutputPath" name="OutputPath" placeholder="请输入工作路径">
    </div>
    <div class="input-field">
        <label for="ProtoPath">ProtoPath</label>
        <input type="text" id="ProtoPath" name="ProtoPath" placeholder="请输入Proto路径">
    </div>
    <div class="input-field">
        <label for="ActivityChineseName">ActivityChineseName</label>
        <input type="text" id="ActivityChineseName" name="ActivityChineseName" placeholder="请输入活动中文名字">
    </div>
    <div class="input-field">
        <label for="IsSupportTasks">IsSupportTasks</label>
        <input type="number" id="IsSupportTasks" name="IsSupportTasks" placeholder="是否支持任务(0 No, 1 Yes)" min="0"
               max="1">
    </div>
    <div class="input-field">
        <label for="IsPaySupport">IsPaySupport</label>
        <input type="number" id="IsPaySupport" name="IsPaySupport" placeholder="是否支持支付(0 No, 1 Yes)" min="0"
               max="1">
    </div>
    <div class="input-field">
        <label for="IsTable">IsTable</label>
        <input type="number" id="IsTable" name="IsTable" placeholder="是否需要sql转Go(0 No, 1 Yes)" min="0" max="1">
    </div>
    <div class="input-field">
        <label for="IsTable">TableName</label>
        <input type="text" id="TableName" name="TableName" placeholder="数据库中的表名，多个用逗号分开">
    </div>
    <button class="btn" onclick="submitData()">提交</button>
</div>

<div class="container" style="background: #4c7bd9;">
    <p>举例:</p>
    <textarea name="" id="" cols="80" rows="12">
--PkgName=news
--Id=7099
--StructName=news
--Author=王二狗
--OutputPath="E:\GoProject\src\work\test_server\branches\beta"
--ProtoPath="E:\GoProject\src\work\docs\branches\beta"
--IsSupportTasks=0
--IsPaySupport=1
--ActivityChineseName="活动中心"
--IsTable=0
--TableName=venture,levelcharge
    </textarea>

</div>
<script>
    function checkComma(str) {
        return /^[^\u002c]+(\u002c)?[^\u002c]+$/.test(str);
    }
    function submitData() {
        let data = {
            PkgName: document.getElementById('PkgName').value,
            Id: document.getElementById('Id').value,
            StructName: document.getElementById('StructName').value,
            Author: document.getElementById('Author').value,
            OutputPath: document.getElementById('OutputPath').value,
            ProtoPath: document.getElementById('ProtoPath').value,
            ActivityChineseName: document.getElementById('ActivityChineseName').value,
            IsSupportTasks: document.getElementById('IsSupportTasks').value,
            IsPaySupport: document.getElementById('IsPaySupport').value,
            IsTable: document.getElementById('IsTable').value,
            TableName: document.getElementById('TableName').value
        };
        var xhr = new XMLHttpRequest();
        xhr.open('POST', 'http://127.0.0.1:8080/fc_scaffold');
        xhr.setRequestHeader('Content-Type', 'application/json');
        xhr.onload = function () {
            if (xhr.status === 200) {
                // 存储到localStorage中
                localStorage.setItem('OutputPath', document.getElementById('OutputPath').value);
                localStorage.setItem('ProtoPath', document.getElementById('ProtoPath').value);
                localStorage.setItem('Author', document.getElementById('Author').value);
                alert(xhr.responseText);
            } else {
                let err = xhr.statusText + ":" + xhr.responseText
                alert(err);
            }
        };
        xhr.onerror = function () {
            alert('请求错误');
        };
        data.IsPaySupport = Number(data.IsPaySupport);
        data.IsSupportTasks = Number(data.IsPaySupport);
        data.IsTable = Number(data.IsTable);

        if (data.IsTable === 1 && data.TableName.trim() === "") {
            alert("请输入表名");
            return;
        }

        // 查找逗号的位置
        const commaIndex = data.TableName.indexOf("，");
        if (commaIndex !== -1) {
            alert("输入的内容中包含非英文逗号!");
            return;
        }

        xhr.send(JSON.stringify(data));
    }

    // 从localStorage中获取数据（注意键名大小写）
    const OutputPathData = localStorage.getItem('OutputPath');
    const ProtoPathData = localStorage.getItem('ProtoPath');
    const AuthorData = localStorage.getItem('Author');
    // 如果数据存在，将其显示在输入框中
    if (OutputPathData) {
        document.getElementById('OutputPath').value = OutputPathData;
    }
    if (ProtoPathData) {
        document.getElementById('ProtoPath').value = ProtoPathData;
    }
    if (AuthorData) {
        document.getElementById('Author').value = AuthorData;
    }


</script>
</body>

</html>
`

const SqlToGo = `<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <title>SqlToGo</title>
    <style type="text/css">
        body {
            font-family: Arial, sans-serif;
            background-color: #f2f2f2;
            padding: 20px;
            overflow: hidden;
        }

        .form-container {
            max-width: 600px;
            margin: 0 auto;
            display: flex;
            flex-direction: column;
            justify-content: center;
            align-items: center;
            height: 100vh;
            opacity: 0;
            transform: translateY(-50px);
            animation: fade-up 1s forwards;
        }

        label {
            font-size: 20px;
            margin-bottom: 10px;
        }

        input[type="text"] {
            padding: 10px;
            font-size: 16px;
            border: none;
            border-radius: 4px;
            background-color: #e9e9e9;
            margin-bottom: 20px;
            width: 100%;
            box-sizing: border-box;
        }

        button {
            padding: 15px 30px;
            font-size: 20px;
            font-weight: bold;
            color: white;
            background-color: #3498db;
            border: none;
            border-radius: 4px;
            margin-top: 20px;
            cursor: pointer;
            transition: background-color 0.3s;
            box-shadow: 0px 3px 5px rgba(0, 0, 0, 0.3);
        }

        button:hover {
            background-color: #2980b9;
        }

        .text-animation {
            animation-duration: 2s;
            animation-name: slidein;
        }

        @keyframes slidein {
            from {
                margin-left: -100%;
                width: 300%;
                opacity: 0;
            }

            to {
                margin-left: 0%;
                width: 100%;
                opacity: 1;
            }
        }

        .title {
            font-size: 28px;
            font-weight: bold;
            text-align: center;
            margin-bottom: 30px;
            opacity: 0;
            transform: translateY(-50px);
            animation: fade-up 1s 0.5s forwards;
        }

        @keyframes fade-up {
            from {
                opacity: 0;
                transform: translateY(50px);
            }

            to {
                opacity: 1;
                transform: translateY(0);
            }
        }
    </style>
</head>
<body>
<div class="form-container">
    <h1 class="title text-animation">SqlToGo</h1>

    <label for="Filepath">文件输出路径:</label>
    <input type="text" id="Filepath" name="Filepath" placeholder="需要填写绝对路径">

    <label for="TableName">表名:</label>
    <input type="text" id="TableName" name="TableName" placeholder="数据库中的表名，多个请用英文逗号分开">
    <div style="display: flex; justify-content: space-between ; width: 100%;">
        <button type="submit" onclick="submitData()">提交</button>
        <form action="http://127.0.0.1:8080/scaffold" method="Get">
            <button type="submit">脚手架</button>
        </form>
    </div>
</div>

<script>
    function submitData() {
        let data = {
            Filepath: document.getElementById('Filepath').value,
            TableName: document.getElementById('TableName').value,
        };
        const xhr = new XMLHttpRequest();
        xhr.open('POST', 'http://127.0.0.1:8080/sqltogo');
        xhr.setRequestHeader('Content-Type', 'application/json');
        xhr.onload = function () {
            if (xhr.status === 200) {
                // 存储到localStorage中
                localStorage.setItem('Filepath', document.getElementById('Filepath').value);
                alert(xhr.responseText);
            } else {
                let err = xhr.statusText + ":" + xhr.responseText
                alert(err);
            }
        };
        xhr.onerror = function () {
            alert('请求错误');
        };


        if (data.Filepath.trim() === "") {
            alert("请输入路径");
            return;
        }

        if (data.TableName.trim() === "") {
            alert("请输入表名");
            return;
        }
        
        // 查找逗号的位置
        const commaIndex = data.TableName.indexOf("，");
        if (commaIndex !== -1) {
            alert("输入的内容中包含非英文逗号!");
            return;
        }
        xhr.send(JSON.stringify(data));
    }

    const Filepath = localStorage.getItem('Filepath');
    // 如果数据存在，将其显示在输入框中
    if (Filepath) {
        document.getElementById('Filepath').value = Filepath;
    }
</script>
</body>
</html>
`
