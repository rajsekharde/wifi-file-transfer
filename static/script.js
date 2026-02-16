
// Display Files in uploads/ folder
async function loadFiles() {
    const res = await fetch("/files");
    const files = await res.json();

    const fileList = document.getElementById("fileList");

    files.forEach( f => {
        const fileDiv = document.createElement('div');
        const fileDetailsDiv = document.createElement('div');
        const fileName = document.createElement('p');
        const fileSize = document.createElement('p');
        const downloadButton = document.createElement('button');

        fileDiv.id = "fileDiv";
        fileDetailsDiv.id = "fileDetailsDiv";
        fileName.id = "fileName";
        downloadButton.id = "downloadButton";

        fileName.textContent = f;
        fileSize.textContent = "10000B"

        downloadButton.textContent = "Download";
        downloadButton.onclick = (() => {
            window.location.href = `/download/${f}`;
        });
        // fileDiv.innerHTML = `${f} <a href="/download/${f}">Download</a>`;
        
        fileDetailsDiv.append(fileName, fileSize);
        fileDiv.append(fileDetailsDiv, downloadButton);
        fileList.appendChild(fileDiv);
    })
}

// Load files on app start up
loadFiles()

// Upload files
async function upload() {
    const file = document.getElementById('fileInput').files[0];

    if(!file)
        return alert('Choose file')

    const formData = new FormData();
    formData.append('file', file);

    await fetch ('/upload', {
        method: 'POST',
        body: formData
    });

    loadFiles();
}