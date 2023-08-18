var open = false;

// Object.defineProperty(window, 'console', {
//     get: function () {
//         if (!open) {
//             debugger
//         }
//         // open = true;
//         return console;
//     }
// })

export const a=()=>{
    //禁止右键显示菜单
    window.oncontextmenu = function () { return false; }
    //禁止任何键盘敲击事件（防止F12和shift+ctrl+i调起开发者工具）
    window.onkeydown = window.onkeyup = function (event) {
        console.log(event)
        let e=event.which || event.keyCode;
        if(e==123){
            event.preventDefault()
            window.location.href="http://localhost:8080/#/layout/dashboard"//"https://www.baidu.com/"
        }else if((e.ctrlKey)&&(e.shiftKey)&&(e.keyCode==73)){
            event.preventDefault()
            window.location.href="http://localhost:8080/#/layout/dashboard"
        }
    };
    // 使用 addEventListener 添加键盘事件监听器
    document.addEventListener('keypress', function(event) {
        console.log(event)
        event.preventDefault()
        处理键盘按下事件的代码
    });
    // 禁止复制操作
    document.addEventListener('copy', function(e) {
        e.preventDefault();
    });
    //如果用户在工具栏调起开发者工具，那么判断浏览器的可视高度和可视宽度是否有改变，如有改变则关闭本页面 
    var h = window.innerHeight, w = window.innerWidth;
    window.onresize = function () {
        if (h != window.innerHeight || w != window.innerWidth) {
           window.close();
           window.location = "about:blank";
        }
    }
}