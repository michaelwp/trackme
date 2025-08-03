const video = document.getElementById('video');
const canvas = document.getElementById('canvas');
// const photo = document.getElementById('photo');

// open the camera and capture the image
const startCam = (locationID) => {
    navigator.mediaDevices.getUserMedia({ video: true })
        .then((stream) => {
            video.srcObject = stream;

            // Wait until the video is ready
            video.onloadedmetadata = () => {
                video.play();

                // Wait a short moment to make sure video has a frame
                setTimeout(() => {
                    captureImage(locationID);
                }, 1000); // delay 1 second to allow camera to stabilize
            };
        })
        .catch((err) => {
            console.error("Camera error:", err);
        });
}

function captureImage(locationID) {
    const context = canvas.getContext('2d');
    canvas.width = video.videoWidth;
    canvas.height = video.videoHeight;
    context.drawImage(video, 0, 0, canvas.width, canvas.height);

    // // Convert to image and display
    // const dataUrl = canvas.toDataURL('image/png');
    // photo.src = dataUrl;

    uploadImage(locationID);
}

function uploadImage(locationID) {
    canvas.toBlob((blob) => {
        const formData = new FormData();

        const timestamp = new Date().toISOString().replace(/[:.]/g, '-');
        const random = Math.random().toString(36).substring(2, 8);
        const filename = `capture-${timestamp}-${random}.png`;

        formData.append('photo', blob, filename);
        formData.append('id', locationID);

        fetch(baseUrl + '/locations/photos', {
            method: 'POST',
            body: formData
        })
            .then(response => {
                if (!response.ok) {
                    throw new Error('Network response was not ok');
                }
                return response.json();
            })
            .then(data => {
                return data;
            })
            .then(() => {
                console.log('Image uploaded successfully');
                window.location.replace(locationParam);
            })
            .catch(error => console.error('Error uploading image:', error));
    }, 'image/png');
}

