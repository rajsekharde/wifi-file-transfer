
// Display Files in uploads/ folder
async function loadFiles() {
    const res = await fetch("/files");
    const files = await res.json();

    const fileList = document.getElementById("fileList");
    fileList.innerHTML = '';

    files.forEach( f => {
        const li = document.createElement('li');
        li.innerHTML = `${f} <a href="/download/${f}">Download</a>`;
        fileList.appendChild(li);
    })
}

// Load files on app start up
loadFiles()