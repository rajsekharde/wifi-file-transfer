async function uploadFile() {
    const fileInput = document.getElementById("fileInput");
    const file = fileInput.files[0];

    if (!file) {
        alert("Select a file first");
        return;
    }

    const formData = new FormData();
    formData.append("file", file);

    await fetch("/upload", {
        method: "POST",
        body: formData
    });

    alert("Upload complete");
    loadFiles();
}

async function loadFiles() {
    const res = await fetch("/files");
    const files = await res.json();

    const list = document.getElementById("fileList");
    list.innerHTML = "";

    files.forEach(file => {
        const li = document.createElement("li");

        const link = document.createElement("a");
        link.href = `/download/${file}`;
        link.textContent = file;

        li.appendChild(link);
        list.appendChild(li);
    });
}

loadFiles();
