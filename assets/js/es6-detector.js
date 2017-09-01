
//https://www.bram.us/2016/10/31/checking-if-a-browser-supports-es6/

var supportsES6 = function() {
    try {
        new Function("(a = 0) => a");
        return true;
    }
    catch (err) {
        return false;
    }
}();

if(!supportsES6){
    document.writeln(
        "<style> " +
        "a{" +
        "color:#575757;" +
        "}#theme-switcher{" +
        "color: #dedede;" +
        "}</style>");


}
