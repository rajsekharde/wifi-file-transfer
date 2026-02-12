# WiFi File Transfer

A simple offline file transfer web app built with FastAPI that allows file transfer between a laptop and a phone over WiFi without using Bluetooth, USB cables, or internet connectivity.

The app runs a local web server on the laptop, and any device connected to the same WiFi network or hotspot can upload/download files through a browser.

## Installation

Clone Git repository:
```bash
git clone https://github.com/rajsekharde/wifi-file-transfer.git
cd wifi-file-transfer
```

Create and activate a python virtual environment:
```bash
python3 -m venv venv
source venv/bin/activate
```

Install dependencies:
```bash
pip3 install -r requirements.txt
```

## Running the app

### Using script

Modify the run.sh script
```bash
cd ~/projects/wifi-file-transfer
# Change this to current project path
```

Make the script executable:
```bash
chmod +x run.sh
```

Run the script:
```bash
./run.sh
```

### Manually

Run the backend server:
```bash
cd wifi-file-transfer
uvicorn main:app --host 0.0.0.0 --port 8000
```

Get the laptop's IP Address:
```bash
hostname -I
# Displays- 90.92.94.96
```

Open the app on a phone connected on the same local network:
```bash
Open http://90.92.94.96:8000 on a browser
```

## How to use

### Uploading files from phone to laptop

Select a file from the device and click on 'Upload'

After successful upload, the file gets stored inside /uploads folder in the project folder on laptop and the 'Files:' list gets updates with the new contents of /uploads folder

### Uploading files from laptop to phone

Copy the file to /uploads folder under root project folder

Refresh the page on the phone. The files list should be updated with new contents

Click on the name of the file and it gets downloaded to the phone