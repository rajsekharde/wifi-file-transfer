# WiFi File Transfer

A simple offline file transfer web app built with FastAPI that allows file transfer between a laptop and a phone over WiFi without using Bluetooth, USB cables, or internet connectivity.

The app runs a local web server on the laptop, and any device connected to the same WiFi network or hotspot can upload/download files through a browser.

## Installation

Clone Git repository

Create and activate a python virtual environment

Install python dependencies using the **requirements.txt** file

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

A QR code gets generated and stored as a png file under /static/

Scan the QR on your phone and open the link in a browser

### Manually

Run the backend server:
```bash
cd wifi-file-transfer
uvicorn main:app --host 0.0.0.0 --port 8000
```

Get the laptop's IP Address

Open **http://<laptop IP>:8000** in a browser on your phone. Both the laptop and phone should be connected over a local network 

## How to use

### Uploading files from phone to laptop

Select a file from your phone and click on 'Upload'

After successful upload, the file gets stored under **uploads** folder under root project folder on the laptop and the **Files** list gets updated with the new contents of **uploads** folder

### Uploading files from laptop to phone

Copy the file to **uploads** folder under root project folder

Refresh the page on the phone. The file list gets updated with the new contents 

Click on the name of the file and it gets downloaded on the phone