const classMap = {
    h1: 'text-2xl font-medium bg-green-100 mb-8 p-2',
    h2: 'text-xl font-medium bg-indigo-100 mt-6 mb-3',
    ul: 'list-disc list-inside',
    li: '',
    pre: 'font-mono bg-gray-200 shadow-lg rounded-md p-2 mb-2 mt-1'
}

const bindings = Object.keys(classMap)
    .map(key => ({
        type: 'output',
        regex: new RegExp(`<${key}(.*)>`, 'g'),
        replace: `<${key} class="${classMap[key]}" $1>`
    }))

const task = "1-1"


var converter = new showdown.Converter({
    extensions: [...bindings]
})

window.onload = function () {
    // var md = "# Variablen\n## Schritt 1\nMach dies und das.\n - Hallo\n - Hallo 2\n## Schritt 2\nNun kommt der wichtige Teil\n - Mittag essen\n - Zvieri essen"

    superagent.get("/lesson/1/1").query({text: 'lessons\\1-Basics\\2-Variablen\\text.md'})
    .then(res => {
        var html = converter.makeHtml(res.body.text)
        $("#explanation").html(html)     
    })
}

function handleCodeRun() {
    var code = editor.getValue()
    
    var statusIcon = $("#statusicon_code")
    statusIcon.removeClass("fa-info-circle")
    statusIcon.addClass("fa-pulse fa-spinner")

    var statusText = $("#statusmessage_code")
    statusText.html("Code wird ausgeführt ...")

    codeoutput.setValue("")

    superagent.post(`/ps/${task}`).send(code).then(res => {
        console.log(res.body)

        statusIcon.addClass("fa-info-circle")
        statusIcon.removeClass("fa-pulse fa-spinner")
        statusText.html(res.body.status)
        codeoutput.setValue(res.body.response)
    })
}

