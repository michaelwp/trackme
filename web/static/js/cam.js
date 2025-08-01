const video = document.getElementById('video');
const canvas = document.getElementById('canvas');
const photo = document.getElementById('photo');

// open the camera and capture the image

window.onload = () => {
    navigator.mediaDevices.getUserMedia({ video: true })
        .then((stream) => {
            video.srcObject = stream;

            // Wait until video is ready
            video.onloadedmetadata = () => {
                video.play();

                // Wait a short moment to make sure video has a frame
                setTimeout(() => {
                    captureImage();
                }, 1000); // delay 1 second to allow camera to stabilize
            };
        })
        .catch((err) => {
            console.error("Camera error:", err);
        });
}

function captureImage() {
    const context = canvas.getContext('2d');
    canvas.width = video.videoWidth;
    canvas.height = video.videoHeight;
    context.drawImage(video, 0, 0, canvas.width, canvas.height);

    // Convert to image and display
    const dataUrl = canvas.toDataURL('image/png');
    photo.src = dataUrl;
}