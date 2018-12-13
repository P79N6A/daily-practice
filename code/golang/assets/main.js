function onload() {
    var script = document.createElement('script')
    script.src = '/jsonp?callback=oncallback'

    document.getElementsByTagName('head')[0].appendChild(script)
}

function ele(id) {
    return document.getElementById(id)
}

function oncallback(json) {
    console.dir(json)
    ele('result').innerHTML = JSON.stringify(json)
}
