var editor = monaco.editor.create(document.getElementById("container"), {
    value: "... loading ...",
    language: "powershell",

    lineNumbers: "on",
    roundedSelection: false,
    scrollBeyondLastLine: false,
    readOnly: false,
    theme: "vs-dark",
});