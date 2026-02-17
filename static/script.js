
// Display Files in uploads/ folder
async function loadFiles() {
    const res = await fetch("/files");
    const files = await res.json();

    const fileList = document.getElementById("fileList");
    fileList.innerHTML = '';

    files.forEach(f => {
        const fileDiv = document.createElement('div');
        const fileDetailsDiv = document.createElement('div');
        const fileName = document.createElement('p');
        const fileSize = document.createElement('p');
        const downloadButton = document.createElement('button');

        fileDiv.classList.add("fileDiv");
        fileDetailsDiv.classList.add("fileDetailsDiv");
        fileName.classList.add("fileName");
        fileSize.classList.add("fileSize");
        downloadButton.classList.add("downloadButton");

        fileName.textContent = f.name;
        fileSize.textContent = `${f.size} B`

        downloadButton.textContent = "Download";

        downloadButton.onclick = ((e) => {
            downloadButton.textContent = "Downloading...";
            downloadButton.disabled = true;

            const link = document.createElement("a");
            link.href = `/download/${f.name}`;
            link.download = f.name;
            document.body.appendChild(link);
            link.click();
            document.body.removeChild(link);

            setTimeout(() => {
                downloadButton.textContent = "Download";
                downloadButton.disabled = false;
            }, 1000);
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
    const uploadButton = document.getElementById("uploadButton");
    uploadButton.textContent = "Uploading";
    uploadButton.disabled = true;

    const files = document.getElementById('fileInput').files;

    if (!files.length) {
        uploadButton.textContent = "Upload";
        uploadButton.disabled = false;
        return alert('Choose file');
    }

    const formData = new FormData();
    
    for(let i=0; i<files.length; i++){
        formData.append('files', files[i]);
    }

    await fetch('/upload', {
        method: 'POST',
        body: formData
    });

    document.getElementById("fileInput").value = "";
    uploadButton.textContent = "Upload";
    uploadButton.disabled = false;

    loadFiles();
}